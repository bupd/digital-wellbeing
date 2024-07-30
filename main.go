package main

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	idleThreshold     = 10 * time.Minute
	notifyThreshold   = 5 * time.Minute
	checkInterval     = 10 * time.Second // Check every 10 seconds
	notifyCommand     = "notify-send"
	shutdownCommand   = "systemctl powerkill poweroff"
	xprintidleCommand = "xprintidle"
)

func getIdleTime() (time.Duration, error) {
	out, err := exec.Command(xprintidleCommand).Output()
	if err != nil {
		return 0, err
	}
	idleMilliseconds, err := strconv.ParseInt(strings.TrimSpace(string(out)), 10, 64)
	if err != nil {
		return 0, err
	}
	return time.Duration(idleMilliseconds) * time.Millisecond, nil
}

func notifyUser(msg string) error {
	if msg != "" {
		return exec.Command(notifyCommand, msg).
			Run()
	}
	return exec.Command(notifyCommand, "AFK Alert", "You have been inactive for 5 minutes. The PC will shut down in 2 minutes if no activity is detected.").
		Run()
}

func shutdownPC() error {
	if err := exec.Command("powerkill").Run(); err != nil {
		log.Fatalf("Failed while executing powerkill: %v", err)
		return err
	}
	return exec.Command("systemctl", "poweroff").Run()
}

func main() {
  runSqlite()
	notifyUser(fmt.Sprintf("Minutes AFK:%s , %s", idleThreshold, checkInterval))
	for {
		idleTime, err := getIdleTime()
		if err != nil {
			fmt.Println("Error getting idle time:", err)
			time.Sleep(checkInterval)
			continue
		}
		log.Println("idle time: ", idleTime)

		if idleTime >= idleThreshold {
			if err := notifyUser(""); err != nil {
				fmt.Println("Error notifying user:", err)
				time.Sleep(checkInterval)
				continue
			}
			time.Sleep(notifyThreshold)

			idleTime, err = getIdleTime()
			if err != nil {
				fmt.Println("Error getting idle time:", err)
				time.Sleep(checkInterval)
				continue
			}

			if idleTime >= idleThreshold+notifyThreshold {
				if err := shutdownPC(); err != nil {
					fmt.Println("Error shutting down PC:", err)
				}
			}
		}

		time.Sleep(checkInterval)
	}
}

func runSqlite() error {
	log.Println("running sqlite")
	database, err := sql.Open("sqlite3", fmt.Sprintf("./kumar.db"))
	if err != nil {
		log.Printf("err getting db")
		return err
	}
	defer database.Close()
	statement, err := database.Prepare(
		"CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)",
	)
	if err != nil {
		log.Printf("err getting db")
		return err
	}
	statement.Exec()
	log.Println("completed sqlite")
	return nil
}
