package bootstrap

import (
	"github.com/tommynurwantoro/kafkid/config"
	"github.com/tommynurwantoro/kafkid/internal/adapter/kafka"
	"github.com/tommynurwantoro/kafkid/internal/adapter/rest"
)

func RegisterRest(conf *config.Configuration) {
	appContainer.RegisterService("fiber", new(rest.Fiber))
	appContainer.RegisterService("publisher", new(kafka.Publisher))

	if conf.Consumer.Enable {
		appContainer.RegisterService("subscriber", new(kafka.Subscriber))
	}
}
