package base

import (
	"github.com/gofiber/fiber/v2"
)

type BaseHandler struct {
	repository *GormRepository
}

func (handler *BaseHandler) GetAll(c *fiber.Ctx) error {
	data, err := handler.repository.get("", "")

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(data)
}
