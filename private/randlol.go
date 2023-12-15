package main

import (
	"context"
	"fmt"

	"github.com/Kyagara/equinox"
	"github.com/gofiber/fiber/v2"
)

func inicioHandler(c *fiber.Ctx) error {
	return c.SendFile("public/inicio.html")
}

func randlolHandler(c *fiber.Ctx) error {
	return c.SendFile("public/randlol.html")
}

func pruebaapiHandler(c *fiber.Ctx) error {
	client, err := equinox.NewClient("RGAPI-5a4bccc9-8e19-4592-9d6a-989b18103e8d")
	if err != nil {
		fmt.Println("Error al crear el cliente de riot: ", err)
	}
	ctx := context.Background()
	prueba, _ := client.DDragon.Champion.ByName(ctx, "13.24.1", "es_MX", "Aatrox")
	return c.JSON(prueba)
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
	web.Get("/pruebaapi", pruebaapiHandler)
	web.Listen(":403")
}
