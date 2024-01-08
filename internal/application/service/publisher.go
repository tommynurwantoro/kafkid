package service

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type PublisherService struct {
	Publisher message.Publisher `inject:"publisher"`
}

func (s *PublisherService) Startup() error { return nil }

func (s *PublisherService) Shutdown() error { return nil }

func (s *PublisherService) Publish(topic string, m interface{}) (*message.Message, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	msg := message.NewMessage(watermill.NewUUID(), b)

	if err := s.Publisher.Publish(topic, msg); err != nil {
		return msg, err
	}

	return msg, nil
}
