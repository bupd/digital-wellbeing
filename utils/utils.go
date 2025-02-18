package utils

func BoolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func IntToBool(i int64) bool {
	if i == 0 {
		return false
	}
	return true
}
