package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// Type pour JSON "images"
type StringArray []string

func (sa *StringArray) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}
	return json.Unmarshal(b, sa)
}

func (sa StringArray) Value() (driver.Value, error) {
	return json.Marshal(sa)
}

// Modèle Product avec UUID comme ID
type Product struct {
	ID              uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	Category        string      `json:"category"`
	Price           float64     `json:"price"`
	Taille          string      `json:"taille"`
	Marque          string      `json:"marque"`
	Couleur         string      `json:"couleur"`
	Caracteristique string      `json:"caracteristique"`
	Images          StringArray `gorm:"type:jsonb" json:"images"`
	Stock           int         `json:"stock"`
	Rating          int         `json:"rating"`
	Favorit         bool        `json:"favorit"`
}

// Avant de créer un produit, génère l'UUID
func (p *Product) BeforeCreate(tx any) (err error) {
	p.ID = uuid.New()
	return
}
