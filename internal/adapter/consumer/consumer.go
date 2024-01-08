package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/tommynurwantoro/kafkid/config"
	"github.com/tommynurwantoro/kafkid/internal/adapter/kafka"
	"github.com/tommynurwantoro/kafkid/internal/pkg/logger"
)

type App struct {
	Subscriber *kafka.Subscriber     `inject:"subscriber"`
	Config     *config.Configuration `inject:"config"`
}

func (c *App) Startup() error { return nil }

func (c *App) Shutdown() error { return nil }

func (c *App) Consume(topic string, messages <-chan *message.Message) {
	for msg := range messages {
		var messageContent interface{}

		if err := json.Unmarshal(msg.Payload, &messageContent); err != nil {
			logger.Error("Error while unmarshalling message", err)
		} else {
			logger.Info(fmt.Sprintf("Message received from topic %s", topic))
			logger.Info(fmt.Sprintf("%v\n", messageContent))
		}

		msg.Ack()
	}
}
