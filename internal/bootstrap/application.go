package bootstrap

import (
	"github.com/tommynurwantoro/kafkid/internal/application/api"
	"github.com/tommynurwantoro/kafkid/internal/application/service"
)

func RegisterService() {
	appContainer.RegisterService("publisherService", new(service.PublisherService))
}

func RegisterApi() {
	appContainer.RegisterService("pingHandler", new(api.PingHandler))
	appContainer.RegisterService("publisherHandler", new(api.PublisherHandler))
}
