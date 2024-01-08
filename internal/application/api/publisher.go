package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommynurwantoro/kafkid/internal/adapter/rest"
	"github.com/tommynurwantoro/kafkid/internal/domain/publisher"
	"github.com/tommynurwantoro/kafkid/internal/pkg/validator"
)

type PublisherHandler struct {
	App              *rest.Fiber         `inject:"fiber"`
	Validator        validator.Validator `inject:"validator"`
	PublisherService publisher.Service   `inject:"publisherService"`
}

func (h *PublisherHandler) Startup() error {
	v1 := h.App.Group("/v1")
	v1.Post("/publish", h.Publish)

	return nil
}

func (h *PublisherHandler) Shutdown() error { return nil }

func (h *PublisherHandler) Publish(c *fiber.Ctx) error {
	var payload *publisher.PublishRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing request body",
			"error":   err.Error(),
		})
	}

	if err := h.Validator.Validate(c.Context(), payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while validating request body",
			"error":   err.Error(),
		})
	}

	msg, err := h.PublisherService.Publish(payload.Topic, payload.Message)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while publishing message",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      msg.UUID,
		"message": "Message published successfully",
	})
}
