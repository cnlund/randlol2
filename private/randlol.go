package main

import (
	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var llave = "RGAPI-e6584dc2-20f1-43aa-af99-f9e8036a543e"

func inicioHandler(c *fiber.Ctx) error {
	return c.SendFile("public/inicio.html")
}

func randlolHandler(c *fiber.Ctx) error {
	return c.SendFile("public/randlol.html")
}

func API2Handler(c *fiber.Ctx) error {
	return c.SendFile("public/API2.html")
}

func listachampsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listachamps, _ := client.DataDragon.GetChampions()
	return c.JSON(listachamps)
}

func main() {
	web := fiber.New()
	web.Get("/", inicioHandler)
	web.Static("/css", "/public/inicio.css")

	//Aqui la seccion perteneciente al manejo de recusos de RANDLOL
	web.Post("/randlol", randlolHandler)
	web.Static("/randlol/css", "/public/randlol.css")
	web.Static("/randlol/todas", "/public/imagenes-randlol/emblem-challenger.png")
	web.Static("/randlol/asesino", "/public/imagenes-randlol/assassin.png")
	web.Static("/randlol/peleador", "/public/imagenes-randlol/fighter.png")
	web.Static("/randlol/mago", "/public/imagenes-randlol/mage.png")
	web.Static("/randlol/tirador", "/public/imagenes-randlol/marksman.png")
	web.Static("/randlol/soporte", "/public/imagenes-randlol/support.png")
	web.Static("/randlol/tanque", "/public/imagenes-randlol/tank.png")
	web.Get("/listachamps", listachampsHandler)

	//Aqui la seccion perteneciente al manejo de recusos de API2
	web.Post("/API2", API2Handler)
	web.Static("/API2/css", "/public/API2.css")

	web.Listen(":403")
}
