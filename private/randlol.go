package main

import "github.com/gofiber/fiber/v2"

func main() {
	web := fiber.New()
	web.Listen(":403")
}
