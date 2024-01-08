package publisher

import "github.com/ThreeDotsLabs/watermill/message"

type Service interface {
	Publish(topic string, m interface{}) (*message.Message, error)
}
