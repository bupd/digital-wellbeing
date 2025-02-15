package events

import (
	"context"
	"fmt"
	"log"

	"github.com/bupd/digital-wellbeing/internal/database"
	"github.com/bupd/digital-wellbeing/pkg/keymap"
	hook "github.com/robotn/gohook"
)

const KeyDown = 3

type rune = int32

type Event struct {
	Kind    string
	Rawcode int
	Keychar int
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

			keyname := keymap.GetKeyName(keychar)
			if keychar == 0 {
				if ev.Rawcode > 65469 && ev.Rawcode < 65482 {
					keyname = keymap.GetFKeyName(ev.Rawcode)
				}
				keyname = keymap.GetMiscKeyName(ev.Rawcode)
			}

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
