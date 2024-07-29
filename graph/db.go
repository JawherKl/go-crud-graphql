package graph

import (
   "log"
   "os"
   
   "github.com/go-pg/pg"
)
func Connect() *pg.DB {
   	//connStr := os.Getenv("DB_URL")
   // Lire les informations de connexion depuis les variables d'environnement
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	addr := os.Getenv("DB_ADDR")
	database := os.Getenv("DB_NAME")

	opts := &pg.Options{
		User:     user,
		Password: password,
		Addr:     addr,
		Database: database,
	}

	// Connectez-vous à la base de données
	db := pg.Connect(opts)

	// Vérifiez la connexion
	if _, err := db.Exec("SELECT 1"); err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	log.Println("Connection to PostgreSQL successful.")
	return db
}