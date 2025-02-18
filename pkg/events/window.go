package events

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

type Window struct {
	WmName   string `json:"wm_name"`
	WmClass  string `json:"wm_class"`
	IsActive bool   `json:"is_active"`
}

// get current Active Window
func GetCurrentWindow() Window {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// xdotool getactivewindow getwindowname getwindowclassname
	cmd := exec.Command("xdotool", "getactivewindow", "getwindowname", "getwindowclassname")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to execute xdotool: %s", err)
	}
	// fmt.Println(string(out))

	// Convert the byte slice to string and split by newline to get the name and class
	windowInfo := string(out)
	lines := splitByNewline(windowInfo)

	// Make sure we have the correct output lines
	if len(lines) < 2 {
		log.Printf("Expected both window name and class but got fewer lines")
		return Window{}
	}

	data := Window{
		WmName:   lines[0],
		WmClass:  lines[1],
		IsActive: true,
	}
	return data
}

func GetAllWindows() []Window {
	// wmctrl -lx
	// Execute wmctrl -lx
	cmd := exec.Command("wmctrl", "-lx")
	out, err := cmd.Output()
	if err != nil {
		log.Printf("failed to execute wmctrl: %v", err)
		return nil
	}

	// Split the output by lines
	lines := strings.Split(string(out), "\n")

	var data []Window

	// Loop through each line in the output
	for _, line := range lines {
		// Skip empty lines
		if line == "" {
			continue
		}

		// Split the line by spaces
		parts := strings.Fields(line)

		// Check if we have enough parts to parse
		if len(parts) < 5 {
			continue // Skip lines that don't have the expected number of columns
		}

		// Extract data from the line
		windowClass := parts[2]
		windowName := strings.Join(parts[4:], " ")

		classes := strings.Split(windowClass, ".")
		windowClass = classes[1]

		// fmt.Printf("class: %v, name: %v \n", windowClass, windowName)

		// Create a Window struct and append it to the data slice
		window := Window{
			WmClass: windowClass,
			WmName:  windowName,
		}
		data = append(data, window)
	}
	// fmt.Println(data)

	return data
}

// MergeWindows combines the current window with the list of all windows, avoiding duplicates
func MergeWindows(allWindows []Window, currentWindow Window) []Window {
	// Check if the current window is already in the allWindows slice
	for i, window := range allWindows {
		// fmt.Printf("class: %v, name: %v \n", window.WmClass, window.IsActive)
		if window.WmClass == currentWindow.WmClass && window.WmName == currentWindow.WmName {
			allWindows[i].IsActive = true
			return allWindows
		}
	}

	// If not found, add the current window to the list
	return append(allWindows, currentWindow)
}

// Helper function to split the window info by newline
func splitByNewline(s string) []string {
	return []string(strings.Split(s, "\n"))
}
