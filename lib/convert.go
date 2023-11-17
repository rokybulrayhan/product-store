package lib

import "strconv"

func ConvertToInt(s string, fallback int) int {

	if s == "" {
		return fallback
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return fallback
	}

	return i
}
