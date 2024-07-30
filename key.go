package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/bupd/digital-wellbeing/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

func ensureDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating directory: %v", err)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home directory: %v", err)
	}

	dbDir := filepath.Join(homeDir, "digital-wellbeing")
	ensureDirectory(dbDir)

	dbPath := filepath.Join(dbDir, "kumar.db")
	log.Printf("Database path: %s", dbPath)

	sqldb, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer sqldb.Close()

	db := database.New(sqldb)
	db.CreateUser(ctx, database.CreateUserParams{
		Name:      "kumar",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	// statement, err := database.Prepare(
	// 	"CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)",
	// )
	// if err != nil {
	// 	log.Fatalf("Error preparing statement: %v", err)
	// }
	// defer func() {
	// 	if err := statement.Close(); err != nil {
	// 		log.Fatalf("Error closing statement: %v", err)
	// 	}
	// }()
	//
	// _, err = statement.Exec()
	// if err != nil {
	// 	log.Fatalf("Error executing statement: %v", err)
	// }
	log.Println("completed sqlite")
}
