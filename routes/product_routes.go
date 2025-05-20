package routes

import (
	"github.com/gofiber/fiber/v2"
	"product-service/handlers"
	
)

func ProductRoutes(app *fiber.App) {
	// Groupe de routes avec middleware d'authentification
	api := app.Group("/products")
	// Routes protégées
	api.Post("/AddProduct", handlers.AddProduct)
	api.Get("/GetAllProducts", handlers.GetAllProducts)
	api.Get("/GetProductByID/:id", handlers.GetProductByID)
	api.Put("/UpdateProduct/:id", handlers.UpdateProduct)
	api.Delete("/DeleteProduct/:id", handlers.DeleteProduct)


}
