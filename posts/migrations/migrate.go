package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"blog-services/common"
)

//go:embed *.sql
var migrations embed.FS

func main() {
	dbUser := common.Env("DB_USER", "postgres")
	dbName := common.Env("DB_NAME", "blog_posts_db")
	dbPassword := common.Env("DB_PASSWORD", "postgres")
	dbHost := common.Env("DB_HOST", "localhost")
	dbPort := common.Env("DB_PORT", "5432")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	migrationFiles, err := migrations.ReadDir(".")
	if err != nil {
		log.Fatal("Failed to read migrations:", err)
	}

	for _, file := range migrationFiles {
		if file.Name() == "migrate.go" {
			continue
		}

		content, err := migrations.ReadFile(file.Name())
		if err != nil {
			log.Fatalf("Failed to read migration %s: %v", file.Name(), err)
		}

		log.Printf("Running migration: %s", file.Name())
		
		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatalf("Failed to execute migration %s: %v", file.Name(), err)
		}
	}

	log.Println("Migrations completed successfully!")
} 