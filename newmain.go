package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"

	"github.com/bupd/digital-wellbeing/internal/database"
)

var loggerLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func run() error {
	ctx := context.Background()
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error(fmt.Sprintf("Cannot find user home dir: %v", err))
		return err
	}

	db, err := sql.Open("sqlite3", filepath.Join(home, ".digital-wellbeing/data.db"))
	if err != nil {
		slog.Error(fmt.Sprintf("Cannot open Database: %v", err))
		return err
	}
	defer db.Close()

	queries := database.New(db)
	tx, err := db.BeginTx(ctx, nil)
	users, err := queries.WithTx(tx).ListUsers(ctx)
	if err != nil {
		return err
	}
	log.Println(users)
	return nil
}

type opts struct {
	logger slog.Level
}

func main() {
	var logLevel string

	flag.StringVar(&logLevel, "log-level", "debug", "sets logger level")
	flag.Parse()

	loglevel, ok := loggerLevels[logLevel]
	if !ok {
		log.Fatalf("invalid log level you dumb shit: %v", logLevel)
	}

	opts := &slog.HandlerOptions{
		Level: loglevel,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
