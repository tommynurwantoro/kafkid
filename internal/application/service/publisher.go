package service

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/tommynurwantoro/kafkid/internal/adapter/kafka"
)

type PublisherService struct {
	Publisher *kafka.Publisher `inject:"publisher"`
}

func (s *PublisherService) Startup() error { return nil }

func (s *PublisherService) Shutdown() error { return nil }

func (s *PublisherService) Publish(topic string, m interface{}) (*message.Message, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	msg := message.NewMessage(watermill.NewUUID(), b)

	s.Publisher.SetKey("test-key")
	if err := s.Publisher.Publish(topic, msg); err != nil {
		return msg, err
	}

	return msg, nil
}
