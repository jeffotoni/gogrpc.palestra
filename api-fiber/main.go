package main

import (
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/embed"

	"github.com/gofiber/fiber"
	mw "github.com/gofiber/fiber/middleware"

	_ "github.com/jeffotoni/gogrpc.palestra/api-fiber/statik"
	"github.com/rakyll/statik/fs"
)

var (
	port = 8080
)

//
func main() {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	////
	app.Use(cors.New())
	app.Use(mw.Compress(mw.CompressLevelBestSpeed))
	app.Use(mw.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n"))

	///
	app.Get("/ping", Ping)

	// Static file server
	app.Use("/", embed.New(embed.Config{
		Root: statikFS,
	}))

	app.Listen(port)

}

func Ping(c *fiber.Ctx) {
	c.Status(200).Send("pong")
}
