package main

import (
	"log"
	"product-service/db"
	"product-service/models"
	"product-service/routes"
	"github.com/gofiber/fiber/v2"

	
)

func main() {
	
	// Initialiser la connexion à la base de données
	db.Connect()

	// Synchroniser le modèle avec la DB
	db.DB.AutoMigrate(&models.Product{})

	// Configurer l'application Fiber
	app := fiber.New()

	// Activer CORS pour autoriser les requêtes depuis le frontend (localhost:3003)
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*", // Autoriser toutes les origines
	// 	AllowMethods:     "GET,POST,PUT,DELETE",
	// 	AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	// 	AllowCredentials: false, // Désactiver les cookies / sessions
	// }))

	app.Static("/images", "./images")

	// Ajouter les routes
	routes.ProductRoutes(app)

	// Démarrer le serveur
	log.Fatal(app.Listen(":3001"))
}
