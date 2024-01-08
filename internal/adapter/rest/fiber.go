package rest

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tommynurwantoro/kafkid/config"
)

type Fiber struct {
	*fiber.App
	Conf *config.Configuration `inject:"config"`
}

func (f *Fiber) Startup() error {
	//starting http
	f.App = fiber.New(fiber.Config{
		ReadTimeout:  time.Duration(f.Conf.Http.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(f.Conf.Http.WriteTimeout) * time.Second,
	})

	return nil
}

func (f *Fiber) Shutdown() error { return f.App.Shutdown() }
