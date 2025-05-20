package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/lib/pq"
)

// Modèle pour les produits
type Product struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey;" json:"id"`
	Name           string         `gorm:"type:varchar(100);not null" json:"name"`
	Description    string         `gorm:"type:text" json:"description"`
	Category       string         `gorm:"type:varchar(100);" json:"category"`
	Price          float64        `gorm:"type:numeric(10,2);not null" json:"price"`
	Taille         string         `gorm:"type:varchar(50);" json:"taille"`
	Marque         string         `gorm:"type:varchar(100);" json:"marque"`
	Couleur        string         `gorm:"type:varchar(50);" json:"couleur"`
	Caracteristique string        `gorm:"type:text;" json:"caracteristique"`
	Images         pq.StringArray `gorm:"type:text[]" json:"images"` 
	Stock          int            `gorm:"type:int;not null" json:"stock"`
	Rating         float64 			  `gorm:"type:int;not null" json:"rating"`
	Favorit        bool           `gorm:"default:false" json:"favorit"` 
}

// Génération automatique de l'UUID avant la création
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
