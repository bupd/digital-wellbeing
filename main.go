package main

import (
	"log"
	"os/exec"
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
