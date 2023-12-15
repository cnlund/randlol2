package main

import (
	"github.com/gofiber/fiber/v2"
)

func inicioHandler(c *fiber.Ctx) error {
	return c.SendFile("public/inicio.html")
}

func randlolHandler(c *fiber.Ctx) error {
	return c.SendFile("public/randlol.html")
}

func main() {
	web := fiber.New()
	web.Get("/", inicioHandler)
	web.Static("/css", "/public/inicio.css")
	web.Post("/randlol", randlolHandler)
	web.Static("/randlol/css", "/public/randlol.css")
	web.Static("/randlol/todas", "/public/imagenes-randlol/emblem-challenger.png")
	web.Static("/randlol/top", "/public/imagenes-randlol/Position_Gold-Top.png")
	web.Static("/randlol/jg", "/public/imagenes-randlol/Position_Gold-Jungle.png")
	web.Static("/randlol/mid", "/public/imagenes-randlol/Position_Gold-Mid.png")
	web.Static("/randlol/bot", "/public/imagenes-randlol/Position_Gold-Bot.png")
	web.Static("/randlol/supp", "/public/imagenes-randlol/Position_Gold-Support.png")
	web.Static("/listachamps", "/champsdata.json")
	web.Listen(":403")
}
