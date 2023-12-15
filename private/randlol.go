package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/p-ob/go-riot/lol"
)

func inicioHandler(c *fiber.Ctx) error {
	return c.SendFile("public/inicio.html")
}

func randlolHandler(c *fiber.Ctx) error {
	return c.SendFile("public/randlol.html")
}

func main() {
	apiKey := "RGAPI-5312ef19-d176-4eec-9961-4140554b5d8c"
	region := lol.Lan
	client := lol.NewClient(apiKey, region, nil)
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
	web.Listen(":403")
}
