package bootstrap

import (
	"github.com/tommynurwantoro/kafkid/internal/application/api"
	"github.com/tommynurwantoro/kafkid/internal/application/service"
	"github.com/tommynurwantoro/kafkid/internal/pkg/validator"
)

func (b *Bootstrap) RegisterService() {
	b.AppContainer.RegisterService("publisherService", new(service.PublisherService))
}

func (b *Bootstrap) RegisterApi() {
	// Initialize struct validator
	b.AppContainer.RegisterService("validator", validator.NewGoValidator())
	b.AppContainer.RegisterService("pingHandler", new(api.PingHandler))
	b.AppContainer.RegisterService("publisherHandler", new(api.PublisherHandler))
}
