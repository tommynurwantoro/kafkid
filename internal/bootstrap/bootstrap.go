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
)

type Bootstrap struct {
	AppContainer container.Container
}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		AppContainer: container.New(),
	}
}

func (b *Bootstrap) RunPubisher(conf *config.Configuration) {
	b.AppContainer.RegisterService("config", conf)

	logger.Info("Serving...")

	// Register adapter
	b.RegisterRest(conf)

	// Register application
	b.RegisterService()
	b.RegisterApi()

	// Startup the container
	if err := b.AppContainer.Ready(); err != nil {
		logger.Panic("Failed to populate service", err)
	}

	// Start server
	fiberApp := b.AppContainer.GetServiceOrNil("fiber").(*rest.Fiber)
	errs := make(chan error, 2)
	go func() {
		fmt.Printf("Listening on port :%d", conf.Http.Port)
		errs <- fiberApp.Listen(fmt.Sprintf(":%d", conf.Http.Port))
	}()

	logger.Info("Publisher started")

	b.gracefulShutdown()
}

func (b *Bootstrap) RunConsumer(conf *config.Configuration) {
	b.AppContainer.RegisterService("config", conf)

	bootstrapContext := context.Background()
	logger.Info("Serving...")

	// Register adapter
	b.RegisterConsumer()

	// Startup the container
	if err := b.AppContainer.Ready(); err != nil {
		logger.Panic("Failed to populate service", err)
	}

	// Start consumer
	subscriberApp := b.AppContainer.GetServiceOrNil("subscriber").(*kafka.Subscriber)

	for _, topic := range conf.Consumer.Topics {
		message, err := subscriberApp.Subscribe(bootstrapContext, topic)
		if err != nil {
			logger.Panic("Failed to subscribe", err)
		}

		go subscriberApp.Consume(topic, message)

		logger.Info(fmt.Sprintf("Subscribed to topic %s", topic))
	}

	logger.Info("Consumer started")

	b.gracefulShutdown()
}

func (b *Bootstrap) gracefulShutdown() {
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	delay := 5 * time.Second

	logger.Info(fmt.Sprintf("Signal termination received. Waiting %v to shutdown.", delay))

	time.Sleep(delay)

	logger.Info("Cleaning up resources...")

	// This will shuting down all the resources
	b.AppContainer.Shutdown()

	logger.Info("Bye")
}
