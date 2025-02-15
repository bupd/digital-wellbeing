package events

import (
	"context"
	"fmt"
	"log"

	"github.com/bupd/digital-wellbeing/internal/database"
	hook "github.com/robotn/gohook"
)

const KeyDown = 3

type rune = int32

type Event struct {
	Kind    string
	Rawcode int
	Keychar int
}

func getKeyName(keychar rune) string {
	// Implement your logic to map Keychar to Keyname
	// For example, a basic lookup might be:
	switch keychar {
	case 106:
		return "j"
	case 107:
		return "k"
	case 13:
		return "Enter"
	// Add more mappings as needed
	default:
		return "Unknown"
	}
}

func StartHookListener(db *database.Queries) {
	// hook.Start() initializes the hook listener
	chanHook := hook.Start()
	defer hook.End()

	// For each event in the hook channel, process key events and add them to the database
	for ev := range chanHook {
		if ev.Kind == KeyDown {
			// keycode := ev.Rawcode
			keychar := ev.Keychar

			keyname := getKeyName(keychar)
			param := database.AddKeyParams{
				Keyname: keyname,
				Keycode: int64(keychar),
			}

			// Call function to add key to database
			row, err := db.AddKey(context.Background(), param)
			if err != nil {
				log.Println("Error adding key to database:", err)
			}

			fmt.Println("Printing row from DB: ", row)
		}
	}
}
