package router

import (
	"github.com/gofiber/fiber"

	h "../handlers"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api")

	SetupProductRouter(api)

}

func SetupProductRouter(router fiber.Router) {
	route := router.Group("/product")

	route.Get("/", h.GetProduct)
	route.Post("/", h.CreateProduct)
	route.Put("/", h.UpdateProduct)
	route.Delete("/", h.DeleteProduct)
}
