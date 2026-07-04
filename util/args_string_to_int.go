package util

import (
	"strconv"
)

func StringToInt(value string) (int, error) {
	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return id, nil
}
