package service

import "strings"

func solveImageExtension(extension string) bool {
	e := strings.ToLower(extension)
	switch e {
	case ".jpg":
		return true
	case ".jpeg":
		return true
	default:
		return false
	}
}
