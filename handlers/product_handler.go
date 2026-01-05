package handlers

import (
	// "fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"product-service/db"
	"product-service/models"
	"strconv"

	// "strings"
	// "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// }

func AddProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	// Ajouter l'ID du produit
	product.ID = uuid.New()
	// Obtenir le répertoire de travail courant
	workingDir, err := os.Getwd()
	if err != nil {
		log.Println("Erreur lors de la récupération du répertoire courant :", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	// Vérification des fichiers d'image envoyés
	var imagePaths []string
	files, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse form"})
	}

	// Parcourir les fichiers envoyés
	for _, fileHeader := range files.File["images"] {
		// Enregistrer chaque image
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open file"})
		}
		defer file.Close()

		// Créer un nom de fichier unique
		fileName := uuid.New().String() + "_" + fileHeader.Filename

		// Créer le chemin complet
		imagesDir := filepath.Join(workingDir, "images", "products")
		filePath := filepath.Join(imagesDir, fileName)

		// Créer le dossier s'il n'existe pas
		if err := os.MkdirAll(imagesDir, os.ModePerm); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create image directory"})
		}

		// Créer le fichier destination
		dst, err := os.Create(filePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create image file"})
		}
		defer dst.Close()

		// Copier le contenu de l'image
		if _, err := io.Copy(dst, file); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save image file"})
		}

		// Log du chemin de l'image
		log.Println("Image enregistrée dans :", filePath)

		// Ajouter à la liste des chemins d'image
		// Ici on garde un chemin relatif pour le front : "/images/..."
		imagePaths = append(imagePaths, fileName)
	}

	// Affecter les images au produit
	product.Images = imagePaths

	// Enregistrer dans la base de données
	if result := db.DB.Create(&product); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Produit créé avec succès",
		"product": product,
	})
}

func GetAllProducts(c *fiber.Ctx) error {

	var products []models.Product
	if result := db.DB.Find(&products); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(products)
}

func GetProductByID(c *fiber.Ctx) error {

	id := c.Params("id")
	var product models.Product
	if result := db.DB.First(&product, "id = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if result := db.DB.First(&product, "id = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	// Lire les données texte du multipart/form
	name := c.FormValue("name")
	price := c.FormValue("price")
	stock := c.FormValue("stock")
	taille := c.FormValue("taille")
	marque := c.FormValue("marque")
	couleur := c.FormValue("couleur")
	category := c.FormValue("category")

	description := c.FormValue("description")
	caracteristique := c.FormValue("caracteristique")

	rating := c.FormValue("rating")

	// Mise à jour des champs
	product.Name = name
	product.Taille = taille
	product.Marque = marque
	product.Couleur = couleur
	product.Category = category
	product.Description = description
	product.Caracteristique = caracteristique

	// Conversion numérique si nécessaire
	if p, err := strconv.ParseFloat(price, 64); err == nil {
		product.Price = p
	}
	if s, err := strconv.Atoi(stock); err == nil {
		product.Stock = s
	}
	if r, err := strconv.Atoi(rating); err == nil {
		product.Rating = r
	}

	// Gestion des fichiers image
	form, err := c.MultipartForm()
	if err == nil && form.File["images"] != nil {
		workingDir, _ := os.Getwd()
		imagesDir := filepath.Join(workingDir, "images", "products")
		os.MkdirAll(imagesDir, os.ModePerm)

		var newImagePaths []string
		for _, fileHeader := range form.File["images"] {
			file, err := fileHeader.Open()
			if err != nil {
				continue
			}
			defer file.Close()

			fileName := uuid.New().String() + "_" + fileHeader.Filename
			filePath := filepath.Join(imagesDir, fileName)

			dst, err := os.Create(filePath)
			if err != nil {
				continue
			}
			defer dst.Close()

			_, err = io.Copy(dst, file)
			if err != nil {
				continue
			}

			newImagePaths = append(newImagePaths, fileName)
		}

		// Mise à jour des images
		if len(newImagePaths) > 0 {
			product.Images = newImagePaths
		}
	}

	// Sauvegarde
	if result := db.DB.Save(&product); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	// _, err := verifyJWT(c)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	// }

	id := c.Params("id")
	if result := db.DB.Delete(&models.Product{}, "id = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(fiber.Map{"message": "Product deleted"})
}
