package server

import (
	"context"
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
	"github.com/bupd/digital-wellbeing/pkg/events"
	"github.com/bupd/digital-wellbeing/utils"
	// hook "github.com/robotn/gohook"
)

// to-do: every 24 hrs put this data to another db named cumulative and aggregated so this will have the results of the past day
// to-do: aggregated holds data for frontend and websites
// to-do: cumulative holds all data more like a scrap of things

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
	// chanHook := hook.Start()
	// defer hook.End()
	//
	// for ev := range chanHook {
	// 	fmt.Printf("hook: %v\n", ev)
	// }

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("PORT is not valid: %v", err)
	}
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable to get user home directory: %v", err)
	}

	db, err := sql.Open("sqlite", home+"/.digital-wellbeing/data.db")
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

	log.Println("Started hook listener listening foor the keys")
	go events.StartHookListener(dbQueries)
	go captureWindowData(dbQueries)

	return server
}

// captureWindowData runs in a goroutine and captures window data every second
func captureWindowData(dbQueries *database.Queries) {
	ticker := time.NewTicker(8 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			var TotalCount, Duration int64
			// Capture current window data
			allWindows := events.GetAllWindows()
			currentWindow := events.GetCurrentWindow()
			windows := events.MergeWindows(allWindows, currentWindow)
			fmt.Println(windows)

			for _, window := range windows {
				query := database.AddWmClassParams{
					WmClass:    window.WmClass,
					WmName:     window.WmName,
					StartTime:  time.Now(),
					EndTime:    time.Time{},
					Duration:   Duration,
					TotalCount: TotalCount,
					IsActive:   utils.BoolToInt(window.IsActive),
					UpdatedAt:  time.Now(),
				}
				res, err := dbQueries.AddWmClass(context.Background(), query)
				if err != nil {
					log.Fatalf("failed to add WmClass: %v", err)
				}

				fmt.Println("added window to db: ", res)
			}

			// Insert data into the DB
			// err := InsertWindowData(db, currentWindow)
			// if err != nil {
			// 	log.Println("Error inserting window data:", err)
			// } else {
			// 	log.Printf("Inserted window data: %+v", currentWindow)
			// }
		}
	}
}

// func InputsHook() {
// 	chanHook := hook.Start()
// 	defer hook.End()
//
// 	for ev := range chanHook {
// 		fmt.Printf("hook: %v\n", ev)
// 	}
// }
