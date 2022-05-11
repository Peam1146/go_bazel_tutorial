package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peammy1146/go_bazel_tutorial/src/greet"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString(greet.Greet())
	})

	log.Fatal(app.Listen(":3000"))
}
