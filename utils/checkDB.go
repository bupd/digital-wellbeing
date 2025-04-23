package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Check if db exists, if not create one
func CheckIfDBExists(DBname, home string) {
	dbPath := filepath.Join(home, ".digital-wellbeing/"+DBname+".db")

	// Check if the database file exists
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		// If the file does not exist, run goose migration
		cmd := exec.Command("goose", "sqlite", dbPath, "up")
		if err := cmd.Run(); err != nil {
			log.Fatalf("failed to run goose migration: %v", err)
		}
		fmt.Println("Database created and migrations applied.")
	}
}
