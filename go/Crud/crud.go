package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

func main() {
	app := fiber.New()
	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})
	app.Post("/api/toodos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		todo.Id = len(todos) + 1
		todos = append(todos, *todo)
		return c.JSON(todos)
	})
	log.Fatal(app.Listen(":4000"))

}
