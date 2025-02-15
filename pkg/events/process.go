package events

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/bupd/digital-wellbeing/internal/database"
	"github.com/bupd/digital-wellbeing/pkg/keymap"
	hook "github.com/robotn/gohook"
)

const (
	KeyDown    = 3
	MouseDown  = 8
	MouseDrag  = 10
	MouseWheel = 11
)

type rune = int32

type Event struct {
	Kind    string
	Rawcode int
	Keychar int
}

// takes down all keymaps for keyboards and mouse
func StartHookListener(db *database.Queries) {
	// hook.Start() initializes the hook listener
	chanHook := hook.Start()
	defer hook.End()

	// For each event in the hook channel, process key events and add them to the database
	for ev := range chanHook {
		if ev.Kind == KeyDown {
			KeyboardPresses(ev, db)
		} else {
			switch ev.Kind {
			case MouseDown:
				mouseDown(ev, db)
				continue
			case MouseWheel:
				mouseScroll(ev, db)
				continue
			}
		}
	}
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
