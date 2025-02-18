package utils

func BoolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
