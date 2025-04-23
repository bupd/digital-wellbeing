// pkg/migrator/migrator.go
package migrator

import (
	"fmt"
	"os"
	"os/exec"
)

// RunMigrations is exported to be used in main.go
func RunMigrations(DBname string) error {
	// Build Goose command
	gooseCmd := fmt.Sprintf("~/.digital-wellbeing/%s.db", DBname)
	cmd := exec.Command("goose", "sqlite", gooseCmd, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running migrations...")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("running goose migrations: %w", err)
	}

	fmt.Println("Migrations completed successfully.")
	return nil
}
