package kafka

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/tommynurwantoro/kafkid/config"
)

type Publisher struct {
	*kafka.Publisher
	Config *config.Configuration `inject:"config"`
}

func (p *Publisher) Startup() error {
	kafkaConfig := kafka.PublisherConfig{
		Brokers:   p.Config.Kafka.Addresses,
		Marshaler: kafka.DefaultMarshaler{},
	}

	pub, err := kafka.NewPublisher(kafkaConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		return err
	}

	p.Publisher = pub
	return nil
}

func (p *Publisher) Shutdown() error { return nil }
