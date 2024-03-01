package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Assalomu alaykum hammaga! Nihfoyat! Uxalayman",
	}

	return c.JSON(resp)
}
