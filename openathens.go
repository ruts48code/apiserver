package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func OpenAthensLogin(ctx *fiber.Ctx) error {
	log.Printf("%v\n", conf.OpenAthens.ConnectionURI)
	return ctx.Render("openathens", fiber.Map{
		"connectionID":  conf.OpenAthens.ConnectionID,
		"connectionURI": conf.OpenAthens.ConnectionURI,
		"returnURL":     conf.OpenAthens.ReturnURL,
		"apiKey":        conf.OpenAthens.APIKey,
	})
}
