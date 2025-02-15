package keymap

import "fmt"

// logic to map Keychar to Keyname
func GetKeyName(keychar rune) string {
	switch keychar {
	case 32:
		return "Space"
	case 48:
		return "0"
	case 49:
		return "1"
	case 50:
		return "2"
	case 51:
		return "3"
	case 52:
		return "4"
	case 53:
		return "5"
	case 54:
		return "6"
	case 55:
		return "7"
	case 56:
		return "8"
	case 57:
		return "9"
	case 65:
		return "A"
	case 66:
		return "B"
	case 67:
		return "C"
	case 68:
		return "D"
	case 69:
		return "E"
	case 70:
		return "F"
	case 71:
		return "G"
	case 72:
		return "H"
	case 73:
		return "I"
	case 74:
		return "J"
	case 75:
		return "K"
	case 76:
		return "L"
	case 77:
		return "M"
	case 78:
		return "N"
	case 79:
		return "O"
	case 80:
		return "P"
	case 81:
		return "Q"
	case 82:
		return "R"
	case 83:
		return "S"
	case 84:
		return "T"
	case 85:
		return "U"
	case 86:
		return "V"
	case 87:
		return "W"
	case 88:
		return "X"
	case 89:
		return "Y"
	case 90:
		return "Z"
	case 65 + 32:
		return "a"
	case 66 + 32:
		return "b"
	case 67 + 32:
		return "c"
	case 68 + 32:
		return "d"
	case 69 + 32:
		return "e"
	case 70 + 32:
		return "f"
	case 71 + 32:
		return "g"
	case 72 + 32:
		return "h"
	case 73 + 32:
		return "i"
	case 74 + 32:
		return "j"
	case 75 + 32:
		return "k"
	case 76 + 32:
		return "l"
	case 77 + 32:
		return "m"
	case 78 + 32:
		return "n"
	case 79 + 32:
		return "o"
	case 80 + 32:
		return "p"
	case 81 + 32:
		return "q"
	case 82 + 32:
		return "r"
	case 83 + 32:
		return "s"
	case 84 + 32:
		return "t"
	case 85 + 32:
		return "u"
	case 86 + 32:
		return "v"
	case 87 + 32:
		return "w"
	case 88 + 32:
		return "x"
	case 89 + 32:
		return "y"
	case 90 + 32:
		return "z"
	case 123:
		return "F12"
	case 27:
		return "Escape"
	case 13:
		return "Enter"
	case 8:
		return "Backspace"
	case 9:
		return "Tab"
	case 16:
		return "Shift"
	case 17:
		return "Ctrl"
	case 18:
		return "Alt"
	case 20:
		return "CapsLock"
	case 19:
		return "Pause"
	case 36:
		return "Home"
	case 37:
		return "LeftArrow"
	case 38:
		return "UpArrow"
	case 39:
		return "RightArrow"
	case 40:
		return "DownArrow"
	case 45:
		return "Insert"
	case 46:
		return "Delete"
	case 33:
		return "PageUp"
	case 34:
		return "PageDown"
	case 192:
		return "`" // Grave accent
	case 189:
		return "-" // Minus key
	case 187:
		return "=" // Equals key
	case 219:
		return "[" // Left bracket
	case 221:
		return "]" // Right bracket
	case 220:
		return "\\" // Backslash
	case 186:
		return ";" // Semicolon
	case 222:
		return "'" // Quote
	case 188:
		return "," // Comma
	case 190:
		return "." // Period
	case 191:
		return "/" // Slash
	default:
		return "Unknown: " + fmt.Sprintf("%v", keychar)
	}
}

// logic to map Function Keys Rawcode to Keyname
func GetFKeyName(rawcode uint16) string {
	// Rawcode: 65470 F1
	// Rawcode: 65481 F12
	switch rawcode {
	case 65470:
		return "F1"
	case 65471:
		return "F2"
	case 65472:
		return "F3"
	case 65473:
		return "F4"
	case 65474:
		return "F5"
	case 65475:
		return "F6"
	case 65476:
		return "F7"
	case 65477:
		return "F8"
	case 65478:
		return "F9"
	case 65479:
		return "F10"
	case 65480:
		return "F11"
	case 65481:
		return "F12"
	default:
		return "Unknown: " + fmt.Sprintf("%v", rawcode)
	}
}
