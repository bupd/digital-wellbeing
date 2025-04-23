// pkg/migrator/migrator.go
package migrator

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// waitForPostgres waits until PostgreSQL is ready
func waitForPostgres(dbHost, dbPort string) error {
	for {
		fmt.Printf("Waiting for PostgreSQL at %s:%s to be ready...\n", dbHost, dbPort)
		err := exec.Command("pg_isready", "-h", dbHost, "-p", dbPort).Run()
		if err == nil {
			fmt.Println("PostgreSQL is ready.")
			return nil
		}
		time.Sleep(10 * time.Second)
	}
}

// RunMigrations is exported to be used in main.go
func RunMigrations() error {
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USERNAME")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbPassword == "" || dbUser == "" || dbName == "" || dbHost == "" || dbPort == "" {
		return fmt.Errorf("missing required environment variables")
	}

	if err := waitForPostgres(dbHost, dbPort); err != nil {
		return fmt.Errorf("waiting for postgres: %w", err)
	}

	// Build Goose command
	gooseCmd := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	cmd := exec.Command("goose", "postgres", gooseCmd, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running migrations...")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("running goose migrations: %w", err)
	}

	fmt.Println("Migrations completed successfully.")
	return nil
}
