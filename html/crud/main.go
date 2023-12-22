package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Block struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Email string `json:"email"`
	Text  string `json:"text"`
}

func main() {
	app := fiber.New()
	blocks := []Block{}

	app.Get("/healthchck", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	app.Post("/block", func(c *fiber.Ctx) error {
		block := &Block{}
		if err := c.BodyParser(block); err != nil {
			return err
		}
		block.Id = len(blocks) + 1

		blocks = append(blocks, *block)
		return c.JSON(blocks)
	})

	app.Get("block/see", func(c *fiber.Ctx) error {
		return c.JSON(blocks)
	})
	log.Fatal(app.Listen(":4000"))

}
