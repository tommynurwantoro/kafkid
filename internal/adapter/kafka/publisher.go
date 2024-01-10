package kafka

import (
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/tommynurwantoro/kafkid/config"
)

type Publisher struct {
	*kafka.Publisher
	PartitionKey kafka.GeneratePartitionKey
	Key          string
	Config       *config.Configuration `inject:"config"`
}

func (p *Publisher) Startup() error {
	defaultConfig := kafka.DefaultSaramaSyncPublisherConfig()
	defaultConfig.ClientID = "kafkid-producer"

	p.PartitionKey = func(topic string, msg *message.Message) (string, error) { return p.Key, nil }
	kafkaConfig := kafka.PublisherConfig{
		Brokers:               p.Config.Kafka.Addresses,
		Marshaler:             kafka.NewWithPartitioningMarshaler(p.PartitionKey),
		OverwriteSaramaConfig: defaultConfig,
	}

	pub, err := kafka.NewPublisher(kafkaConfig, nil)
	if err != nil {
		return err
	}

	p.Publisher = pub
	p.PartitionKey = func(topic string, msg *message.Message) (string, error) { return p.Key, nil }
	return nil
}

func (p *Publisher) Shutdown() error { return nil }

func (p *Publisher) SetKey(key string) {
	p.Key = key
}
