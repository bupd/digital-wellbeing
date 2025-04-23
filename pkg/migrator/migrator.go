// pkg/migrator/migrator.go
package migrator

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

// RunMigrations is exported to be used in main.go
func RunMigrations(DBname string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error(fmt.Sprintf("Cannot find user home dir: %v", err))
		return err
	}
	migrationDir := fmt.Sprintf("%s/.digital-wellbeing/migrations", home)
	// Build Goose command
	gooseCmd := fmt.Sprintf("%s/.digital-wellbeing/%s.db", home, DBname)
	cmd := exec.Command("goose", "-dir", migrationDir, "sqlite", gooseCmd, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running migrations...")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("running goose migrations: %w", err)
	}

	fmt.Println("Migrations completed successfully.")
	return nil
}
