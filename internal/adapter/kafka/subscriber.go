package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/tommynurwantoro/kafkid/config"
	"github.com/tommynurwantoro/kafkid/internal/pkg/logger"
)

type Subscriber struct {
	*kafka.Subscriber
	Conf *config.Configuration `inject:"config"`
}

func (s *Subscriber) Startup() error {
	logger.Info("Starting up subscriber")
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()

	saramaSubscriberConfig.ClientID = "kafkid"
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	if s.Conf.Consumer.OffsetFromNewest {
		saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	}

	subs, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               s.Conf.Kafka.Addresses,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         s.Conf.Consumer.GroupID,
		},
		nil,
	)
	if err != nil {
		logger.Error("Failed to create subscriber", err)
		return err
	}

	s.Subscriber = subs

	return nil
}

func (s *Subscriber) Shutdown() error {
	return s.Subscriber.Close()
}

func (c *Subscriber) Consume(topic string, messages <-chan *message.Message) {
	for msg := range messages {
		var messageContent interface{}
		var partition int32
		var offset int64

		if err := json.Unmarshal(msg.Payload, &messageContent); err != nil {
			logger.Info("Message is not in json format")
			messageContent = string(msg.Payload)
		}

		if part, ok := kafka.MessagePartitionFromCtx(msg.Context()); ok {
			partition = part
		}

		if off, ok := kafka.MessagePartitionOffsetFromCtx(msg.Context()); ok {
			offset = off
		}

		logger.Info(fmt.Sprintf("===== Message Received =====\nTopic: %s\nPartition: %d\nOffset: %d\nRaw: %s\nUnmarshal: %v\n============================", topic, partition, offset, msg.Payload, messageContent))

		msg.Ack()
	}
}
