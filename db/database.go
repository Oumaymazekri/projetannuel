package db

import (
	"log"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	 dsn := "host=product-db user=postgres password=0000 dbname=products port=5432 sslmode=disable"
	//  dsn := "user=postgres password=0000 dbname=products port=5432 sslmode=disable"
	var err error

	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Database connection established")
			return
		}

		log.Printf("Failed to connect to database: %v. Retrying in 5 seconds...\n", err)
		time.Sleep(5 * time.Second)
	}

	log.Fatal("Failed to connect to database after 5 attempts:", err)
}