package main

import (
	"encoding/json"
	"math/rand"
	"os"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Variable global que almacena la clave de la API de League of Legends.
var key, _ = os.ReadFile("key.txt")
var llave = string(key)

// Definición de la estructura Champ para representar información de campeones.
type Champ struct {
	Nombre string   `json:"id"`
	Tipos  []string `json:"tags"`
	Lore   string   `json:"blurb"`
}

type Champstat struct {
	Nombre string `json:"id"`
	Titulo string `json:"title"`
}

// Handler para la ruta raíz que devuelve un archivo HTML.
func inicioHandler(c *fiber.Ctx) error {
	return c.SendFile("public/inicio.html")
}

// Handlers para diferentes rutas que devuelven archivos HTML.
func randlolHandler(c *fiber.Ctx) error {
	return c.SendFile("public/randlol.html")
}

func API2Handler(c *fiber.Ctx) error {
	return c.SendFile("public/API2.html")
}

func API3Handler(c *fiber.Ctx) error {
	return c.SendFile("public/API3.html")
}

// Handler que devuelve una lista de campeones obtenida desde la API de League of Legends.
func listachampsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listachamps, _ := client.DataDragon.GetChampions()
	return c.JSON(listachamps)
}

// Handler que devuelve un campeón aleatorio de la lista de campeones.
func randallchampsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listaallchamps, _ := client.DataDragon.GetChampions()
	jsonlallchamps, _ := json.Marshal(listaallchamps)
	var champs []Champ
	err := json.Unmarshal(jsonlallchamps, &champs)
	if err != nil {
		return err
	}
	randchamp := rand.Intn(len(champs))
	return c.JSON(champs[randchamp])
}

// Handler que devuelve un asesino aleatorio de la lista de campeones.
func randasesinochampsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listaallchamps, _ := client.DataDragon.GetChampions()
	jsonlallchamps, _ := json.Marshal(listaallchamps)
	var champs []Champ
	err := json.Unmarshal(jsonlallchamps, &champs)
	if err != nil {
		return err
	}
	var asesinos []Champ
	for i := range champs {
		for j := range champs[i].Tipos {
			if champs[i].Tipos[j] == "Assassin" {
				asesinos = append(asesinos, champs[i])
			}
		}
	}
	randasesino := rand.Intn(len(asesinos))
	return c.JSON(asesinos[randasesino])
}

func randpeleadorchampsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listaallchamps, _ := client.DataDragon.GetChampions()
	jsonlallchamps, _ := json.Marshal(listaallchamps)
	var champs []Champ
	err := json.Unmarshal(jsonlallchamps, &champs)
	if err != nil {
		return err
	}
	var peleadores []Champ
	for i := range champs {
		for j := range champs[i].Tipos {
			if champs[i].Tipos[j] == "Fighter" {
				peleadores = append(peleadores, champs[i])
			}
		}
	}
	randpeleador := rand.Intn(len(peleadores))
	return c.JSON(peleadores[randpeleador])
}

func randmagochampsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listaallchamps, _ := client.DataDragon.GetChampions()
	jsonlallchamps, _ := json.Marshal(listaallchamps)
	var champs []Champ
	err := json.Unmarshal(jsonlallchamps, &champs)
	if err != nil {
		return err
	}
	var mago []Champ
	for i := range champs {
		for j := range champs[i].Tipos {
			if champs[i].Tipos[j] == "Mage" {
				mago = append(mago, champs[i])
			}
		}
	}
	randmago := rand.Intn(len(mago))
	return c.JSON(mago[randmago])
}

func randtiradorchampsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listaallchamps, _ := client.DataDragon.GetChampions()
	jsonlallchamps, _ := json.Marshal(listaallchamps)
	var champs []Champ
	err := json.Unmarshal(jsonlallchamps, &champs)
	if err != nil {
		return err
	}
	var tirador []Champ
	for i := range champs {
		for j := range champs[i].Tipos {
			if champs[i].Tipos[j] == "Marksman" {
				tirador = append(tirador, champs[i])
			}
		}
	}
	randtirador := rand.Intn(len(tirador))
	return c.JSON(tirador[randtirador])
}

func randsoportechampsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listaallchamps, _ := client.DataDragon.GetChampions()
	jsonlallchamps, _ := json.Marshal(listaallchamps)
	var champs []Champ
	err := json.Unmarshal(jsonlallchamps, &champs)
	if err != nil {
		return err
	}
	var soporte []Champ
	for i := range champs {
		for j := range champs[i].Tipos {
			if champs[i].Tipos[j] == "Support" {
				soporte = append(soporte, champs[i])
			}
		}
	}
	randsoporte := rand.Intn(len(soporte))
	return c.JSON(soporte[randsoporte])
}

func randtanquechampsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listaallchamps, _ := client.DataDragon.GetChampions()
	jsonlallchamps, _ := json.Marshal(listaallchamps)
	var champs []Champ
	err := json.Unmarshal(jsonlallchamps, &champs)
	if err != nil {
		return err
	}
	var tanque []Champ
	for i := range champs {
		for j := range champs[i].Tipos {
			if champs[i].Tipos[j] == "Tank" {
				tanque = append(tanque, champs[i])
			}
		}
	}
	randtanque := rand.Intn(len(tanque))
	return c.JSON(tanque[randtanque])
}

// Handler que devuelve una lista de items obtenida desde la API de League of Legends.
func listaitemsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listaitems, _ := client.DataDragon.GetItems()
	return c.JSON(listaitems)
}

func champstatsHandler(c *fiber.Ctx) error {
	return c.SendFile("public/champsstats.html")
}

func randstatschampsHandler(c *fiber.Ctx) error {
	client := golio.NewClient(llave,
		golio.WithRegion(api.RegionLatinAmericaNorth),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	listaallchamps, _ := client.DataDragon.GetChampions()
	jsonlallchamps, _ := json.Marshal(listaallchamps)
	var champs []Champ
	err := json.Unmarshal(jsonlallchamps, &champs)
	if err != nil {
		return err
	}
	randchamp := rand.Intn(len(champs))
	var idchamp string
	jsonchamps, _ := json.Marshal(champs[randchamp])
	json.Unmarshal(jsonchamps, &idchamp)
	randstatschamp, _ := client.DataDragon.GetChampion(idchamp)
	jsonrandstatchamp, _ := json.Marshal(randstatschamp)
	var statchamp Champstat
	err2 := json.Unmarshal(jsonrandstatchamp, &statchamp)
	if err2 != nil {
		return err2
	}
	return c.JSON(randstatschamp)
}

// Función principal donde se configuran las rutas y se inicia el servidor web.
func main() {
	web := fiber.New()
	web.Get("/", inicioHandler)
	web.Static("/css", "/public/inicio.css")
	web.Static("/logo", "public/imagenes-inicio/logo-lolpedia.png")

	// Sección para el manejo de recursos de RANDLOL
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
	web.Get("/randlol/all", randallchampsHandler)
	web.Get("/randlol/asesinos", randasesinochampsHandler)
	web.Get("/randlol/peleadores", randpeleadorchampsHandler)
	web.Get("/randlol/magos", randmagochampsHandler)
	web.Get("/randlol/tiradores", randtiradorchampsHandler)
	web.Get("/randlol/soportes", randsoportechampsHandler)
	web.Get("/randlol/tanques", randtanquechampsHandler)

	// Sección para el manejo de recursos de API2
	web.Post("/randitems", API2Handler)
	web.Static("/randitems/css", "/public/API2.css")
	web.Get("/randitems/listaitems", listaitemsHandler)

	// Sección para el manejo de recursos de API3
	web.Post("/randline", API3Handler)
	web.Static("/randline/css", "/public/API3.css")

	//Sección para el manejo de recursos de Champstats
	web.Post("/champstats", champstatsHandler)
	web.Static("/champstats/css", "public/champsstats.css")
	web.Get("/champstats/stat", randstatschampsHandler)

	web.Listen(":403")
}
