package bootstrap

import (
	"github.com/tommynurwantoro/kafkid/config"
	"github.com/tommynurwantoro/kafkid/internal/adapter/kafka"
	"github.com/tommynurwantoro/kafkid/internal/adapter/rest"
)

func (b *Bootstrap) RegisterRest(conf *config.Configuration) {
	b.AppContainer.RegisterService("fiber", new(rest.Fiber))
	b.AppContainer.RegisterService("publisher", new(kafka.Publisher))
}

func (b *Bootstrap) RegisterConsumer() {
	b.AppContainer.RegisterService("subscriber", new(kafka.Subscriber))
}
