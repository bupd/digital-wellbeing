package server

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "modernc.org/sqlite"

	"github.com/bupd/digital-wellbeing/internal/database"
	hook "github.com/robotn/gohook"
)

type Server struct {
	port      int
	db        *sql.DB
	dbQueries *database.Queries
}

var (
	dbName   = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	PORT     = os.Getenv("DB_PORT")
	HOST     = os.Getenv("DB_HOST")
)

func NewServer() *http.Server {
	chanHook := hook.Start()
	defer hook.End()

	for ev := range chanHook {
		fmt.Printf("hook: %v\n", ev)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("PORT is not valid: %v", err)
	}

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		slog.Error("New server creation failed in connection to sqlite in memory DB")
		slog.Error(fmt.Sprintf("Failed to open sqlite: %v", err))
		return nil
	}

	dbQueries := database.New(db)

	NewServer := &Server{
		port:      port,
		db:        db,
		dbQueries: dbQueries,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

// func InputsHook() {
// 	chanHook := hook.Start()
// 	defer hook.End()
//
// 	for ev := range chanHook {
// 		fmt.Printf("hook: %v\n", ev)
// 	}
// }
