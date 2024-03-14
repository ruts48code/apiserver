package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type (
	Counter struct {
		Counter int64 `json:"counter"`
	}
)

func counterup(c *fiber.Ctx) error {
	cc := fiber.Get("http://counter_api/counter/" + c.Params("name"))
	cc.InsecureSkipVerify()
	_, data, err := cc.String()
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "server",
		})
	}
	CounterX := Counter{}
	err2 := json.Unmarshal([]byte(data), &CounterX)
	if err2 != nil {
		return c.JSON(fiber.Map{
			"status": "json",
		})
	}
	return c.JSON(fiber.Map{
		"status":  "ok",
		"counter": CounterX.Counter,
	})
}

func counterget(c *fiber.Ctx) error {
	cc := fiber.Get("http://counter_api/counter/" + c.Params("name") + "/get")
	cc.InsecureSkipVerify()
	_, data, err := cc.String()
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "server",
		})
	}
	CounterX := Counter{}
	err2 := json.Unmarshal([]byte(data), &CounterX)
	if err2 != nil {
		return c.JSON(fiber.Map{
			"status": "json",
		})
	}
	return c.JSON(fiber.Map{
		"status":  "ok",
		"counter": CounterX.Counter,
	})
}
