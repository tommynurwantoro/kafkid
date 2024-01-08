package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tommynurwantoro/kafkid/config"
	"github.com/tommynurwantoro/kafkid/internal/adapter/kafka"
	"github.com/tommynurwantoro/kafkid/internal/adapter/rest"
	"github.com/tommynurwantoro/kafkid/internal/pkg/container"
	"github.com/tommynurwantoro/kafkid/internal/pkg/logger"
	"github.com/tommynurwantoro/kafkid/internal/pkg/validator"
)

var appContainer = container.New()

func Run(conf *config.Configuration) {
	appContainer.RegisterService("config", conf)

	// Initialize struct validator
	appContainer.RegisterService("validator", validator.NewGoValidator())

	bootstrapContext := context.Background()
	logger.Info("Serving...")

	// Register adapter
	RegisterRest(conf)

	// Register application
	RegisterService()
	RegisterApi()

	// Startup the container
	if err := appContainer.Ready(); err != nil {
		logger.Panic("Failed to populate service", err)
	}

	// Start server
	fiberApp := appContainer.GetServiceOrNil("fiber").(*rest.Fiber)
	errs := make(chan error, 2)
	go func() {
		fmt.Printf("Listening on port :%d", conf.Http.Port)
		errs <- fiberApp.Listen(fmt.Sprintf(":%d", conf.Http.Port))
	}()

	// Start consumer
	subscriberApp := appContainer.GetServiceOrNil("subscriber").(*kafka.Subscriber)
	for _, topic := range conf.Consumer.Topics {
		message, err := subscriberApp.Subscribe(bootstrapContext, topic)
		if err != nil {
			logger.Panic("Failed to subscribe", err)
		}

		go subscriberApp.Consume(topic, message)
	}

	logger.Info("Your app started")

	gracefulShutdown()
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	delay := 5 * time.Second

	logger.Info(fmt.Sprintf("Signal termination received. Waiting %v to shutdown.", delay))

	time.Sleep(delay)

	logger.Info("Cleaning up resources...")

	// This will shuting down all the resources
	appContainer.Shutdown()

	logger.Info("Bye")
}
