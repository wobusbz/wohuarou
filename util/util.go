package util

import "os"

func Exists(d string) bool {
	_, err := os.Stat(d)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
