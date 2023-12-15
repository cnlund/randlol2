package main

import "github.com/gofiber/fiber/v2"

func inicioHandler(c *fiber.Ctx) error {
	return c.SendFile("public/inicio.html")
}

func main() {
	web := fiber.New()
	web.Get("/", inicioHandler)
	web.Listen(":403")
}
