
package db

import (
    "fmt"
    "log"
    "os"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	
)

var DB *gorm.DB

// getEnv récupère une variable d'env, avec une valeur par défaut si absente.
func getEnv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}

func Connect() {
    // Récupération des variables d'environnement
    host := os.Getenv("PG_HOST")
    port := os.Getenv("PG_PORT")
    user := os.Getenv("PG_USER")
    password := os.Getenv("PG_PASSWORD")
    dbname := os.Getenv("PG_DB_NAME")
    sslmode := os.Getenv("PG_SSLMODE")
      if host == "" || port == "" || user == "" || dbname == "" {
        log.Fatal("❌ Environment variables for PostgreSQL are not set")
    }

	// Construction de la chaîne de connexion
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        host, user, password, dbname, port, sslmode,
    )

    var err error

    const maxAttempts = 5
    const delay = 5 * time.Second

    for i := 1; i <= maxAttempts; i++ {
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            log.Println("✅ Database connection established")
            return
        }

        log.Printf("❌ Attempt %d/%d: failed to connect to database: %v", i, maxAttempts, err)

        if i < maxAttempts {
            log.Printf("⏳ Retrying in %s...\n", delay)
            time.Sleep(delay)
        }
    }

    // Si toutes les tentatives échouent, on log.fatal
    log.Fatalf("Failed to connect to database after %d attempts: %v", maxAttempts, err)
}
