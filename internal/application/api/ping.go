package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommynurwantoro/kafkid/internal/adapter/rest"
)

type PingHandler struct {
	App *rest.Fiber `inject:"fiber"`
}

func (h *PingHandler) Startup() error {
	v1 := h.App.Group("/v1")
	v1.Get("/ping", h.Ping)

	return nil
}

func (h *PingHandler) Shutdown() error { return nil }

func (h *PingHandler) Ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "pong",
	})
}
