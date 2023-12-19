package main

import (
	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var llave = "RGAPI-d4a3e543-bd9c-4911-836c-36e648f1834e"

type Champ struct {
	Nombre string
	Tipos  []string
}

func inicioHandler(c *fiber.Ctx) error {
	return c.SendFile("public/inicio.html")
}

func randlolHandler(c *fiber.Ctx) error {
	return c.SendFile("public/randlol.html")
}

func API2Handler(c *fiber.Ctx) error {
	return c.SendFile("public/API2.html")
}

func API3Handler(c *fiber.Ctx) error {
	return c.SendFile("public/API3.html")
}

func listachampsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listachamps, _ := client.DataDragon.GetChampions()
	return c.JSON(listachamps)
}

func listaitemsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listaitems, _ := client.DataDragon.GetItems()
	return c.JSON(listaitems)
}

func main() {
	web := fiber.New()
	web.Get("/", inicioHandler)
	web.Static("/css", "/public/inicio.css")
	web.Static("/logo", "public/imagenes-inicio/logo-lolpedia.png")

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
	web.Static("/randlol/js", "/public/randlol.js")

	//Aqui la seccion perteneciente al manejo de recusos de API2
	web.Post("/randitems", API2Handler)
	web.Static("/randitems/css", "/public/API2.css")
	web.Get("/randitems/listaitems", listaitemsHandler)
	web.Static("/randitems/js", "public/API2.js")

	//Aqui la seccion perteneciente al manejo de recusos de API3
	web.Post("/randline", API3Handler)
	web.Static("/randline/css", "/public/API3.css")

	web.Listen(":403")
}
