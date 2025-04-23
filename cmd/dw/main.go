package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bupd/digital-wellbeing/pkg/config"
	"github.com/bupd/digital-wellbeing/pkg/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conf := config.GetConfig()
	port := conf.PORT
	dbName := conf.DBNAME

	// update new db daily
	if dbName == "daily" || len(dbName) == 0 {
		// Format: daily-YYYY-MM-DD
		today := time.Now().Format("2006-01-28") // user-friendly format
		dbName = today
		fmt.Println("Updated dbName to:", dbName)
	} else if dbName == "weekly" {
		// Get current year and ISO week number
		year, week := time.Now().ISOWeek()
		// Format: weekly-YYYY-week-WW
		dbName = fmt.Sprintf("%d-week-%02d", year, week)
		fmt.Println("Updated dbName to:", dbName)
	}

	server := server.NewServer(port, dbName)

	fmt.Printf("\nDigital Wellbeing running on server: %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
