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
	"github.com/bupd/digital-wellbeing/pkg/retry"
	"github.com/bupd/digital-wellbeing/utils"
	"github.com/rs/cors"
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

func NewServer(PORT, DBName string) *http.Server {
	// chanHook := hook.Start()
	// defer hook.End()
	//
	// for ev := range chanHook {
	// 	fmt.Printf("hook: %v\n", ev)
	// }

	port, err := strconv.Atoi(PORT)
	if err != nil {
		log.Fatalf("PORT is not valid: %v", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable to get user home directory: %v", err)
	}

	db, err := sql.Open("sqlite", home+"/.digital-wellbeing/"+DBName+".db")
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

	handler := cors.Default().Handler(NewServer.RegisterRoutes())

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Set the busy timeout (in milliseconds)
	db.Exec("PRAGMA busy_timeout = 5000;") // Wait for 5 seconds if the database is locked

	log.Println("Started hook listener listening foor the keys")
	go events.StartHookListener(dbQueries)
	go captureWindowData(dbQueries)

	return server
}

// captureWindowData runs in a goroutine and captures window data every second
func captureWindowData(dbQueries *database.Queries) {
	ticker := time.NewTicker(12 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			var ActiveDuration, Duration int64
			// Capture current window data
			allWindows := events.GetAllWindows()
			currentWindow := events.GetCurrentWindow()
			windows := events.MergeWindows(allWindows, currentWindow)
			fmt.Println(windows)

			for _, window := range windows {
				// GET equivalent window from DB
				windw, err := dbQueries.GetWinByWmName(context.TODO(), window.WmName)
				if err == nil {
					previousUpdate := windw.UpdatedAt
					duration := int64(time.Since(previousUpdate).Seconds())
					// check if duration is greater than 30 secs if it is skip it
					// because we check the windows every 12 seconds so 30 or more could mean the window is dead
					// so we should skip this entry
					if duration > int64(30) {
						duration = 0
					}
					Duration = duration + windw.Duration

					if window.IsActive {
						ActiveDuration = duration + windw.ActiveDuration
					} else {
						ActiveDuration = windw.ActiveDuration
					}
				} else {
					log.Printf("failed to get WmName: %s, err: %v", window.WmName, err)
				}

				query := database.AddWmClassParams{
					WmClass:        window.WmClass,
					WmName:         window.WmName,
					EndTime:        time.Time{},
					Duration:       Duration,
					ActiveDuration: ActiveDuration,
					IsActive:       utils.BoolToInt(window.IsActive),
					UpdatedAt:      time.Now(),
				}
				err = dbQueries.AddWmClass(context.Background(), query)
				if err != nil {
					log.Printf("failed to add WmClass: %v", err)
				}

				// enclose the above one in retry mechanism
				err = retry.Retry(3, 2*time.Second, func() error {
					// Replace this with the actual operation you want to retry.
					fmt.Println("Trying to perform the operation...")
					err := dbQueries.AddWmClass(context.Background(), query)
					if err != nil {
						log.Printf("failed to add WmClass: %v", err)
						return fmt.Errorf("failed to add WmClass to DB: %v", err)
					}
					return nil
				})

				if err != nil {
					fmt.Println("Add to DB failed after retries:", err)
				} else {
					fmt.Println("Curretn Window added to DB")
				}

				// fmt.Println("added window to db: ", res)
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
