package main

import (
	"strings"
	"unicode"
)

// changeCase applies up, low or cap to a single word.
func changeCase(w, tag string) string {
	switch tag {
	case "up":
		return strings.ToUpper(w)
	case "low":
		return strings.ToLower(w)
	case "cap":
		r := []rune(w)
		if len(r) == 0 {
			return w
		}
		return string(unicode.ToUpper(r[0])) + strings.ToLower(string(r[1:]))
	}
	return w
}
