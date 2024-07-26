package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	idleThreshold     = 10 * time.Minute  // 5 minutes
	notifyThreshold   = 2 * time.Minute  // 2 minutes
	checkInterval     = 1 * time.Second // Check every 10 seconds
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

func notifyUser() error {
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
	for {
		idleTime, err := getIdleTime()
		if err != nil {
			fmt.Println("Error getting idle time:", err)
			time.Sleep(checkInterval)
			continue
		}
    log.Println("idle time: ", idleTime)

		if idleTime >= idleThreshold {
			if err := notifyUser(); err != nil {
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
