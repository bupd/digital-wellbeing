package events

import (
	"context"
	"fmt"
	"log"

	"github.com/bupd/digital-wellbeing/internal/database"
	hook "github.com/robotn/gohook"
)

func mouseDown(ev hook.Event, db *database.Queries) string {
	var keyname string

	switch ev.Button {
	case 1:
		keyname = "LeftClick"
	case 2:
		keyname = "MiddleClick"
	case 3:
		keyname = "RightClick"
	case 4:
		keyname = "Back"
	case 5:
		keyname = "Forward"
	default:
		keyname = "Click"
	}

	param := database.AddMouseDownParams{
		EventType: "Click",
		Button:    keyname,
	}

	// Call function to add mouse Click to database
	row, err := db.AddMouseDown(context.Background(), param)
	if err != nil {
		log.Println("Error adding key to database:", err)
	}

	fmt.Println("Printing row from DB: ", row)
	return ""
}

func mouseScroll(ev hook.Event, db *database.Queries) string {
	var keyname string

	switch ev.Rotation {
	case -1:
		keyname = "ScrollUp"
	case 1:
		keyname = "ScrollDown"
	}

	param := database.AddMouseDownParams{
		EventType: "Scroll",
		Button:    keyname,
	}

	// Call function to add mouse Click to database
	row, err := db.AddMouseDown(context.Background(), param)
	if err != nil {
		log.Println("Error adding key to database:", err)
	}

	fmt.Println("Printing row from DB: ", row)
	return ""
}
