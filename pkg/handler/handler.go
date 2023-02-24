package handler

import (
	"fmt"
	"goapi/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func GetRoutes(app *fiber.App) {
	app.Get("/movie", GetMovie)
	app.Post("/movie", PostMovie)
	app.Get("/movie/:id", GetByIdMovie)
	app.Put("/movie/:id", PutByIdMovie)
	app.Delete("/movie/:id", DeleteByIdMovie)
}

// crud clasic
func GetMovie(c *fiber.Ctx) error {
	_, data := service.GetMovie()
	fmt.Println(data)
	return c.SendString("hola")

}
func PostMovie(c *fiber.Ctx) error {
	return c.SendString("Post Fiber Movie")
}
func GetByIdMovie(c *fiber.Ctx) error {
	return c.SendString("Get By Id " + c.Params("id") + " Fiber Movie")
}
func PutByIdMovie(c *fiber.Ctx) error {
	return c.SendString("Put By Id " + c.Params("id") + " Fiber Movie")
}
func DeleteByIdMovie(c *fiber.Ctx) error {
	return c.SendString("Delete By Id " + c.Params("id") + " Fiber Movie")
}
