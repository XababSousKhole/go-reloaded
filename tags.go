package main

import (
	"regexp"
	"strconv"
	"strings"
)

var tagRe = regexp.MustCompile(`\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)`)
var tokenRe = regexp.MustCompile(`\((?:hex|bin|up|low|cap)(?:,\s*\d+)?\)|\S+`)

func applyTags(text string) string {
	var words []string
	for _, tok := range tokenRe.FindAllString(text, -1) {
		m := tagRe.FindStringSubmatch(tok)
		if m == nil {
			words = append(words, tok)
			continue
		}
		n := 1
		if m[2] != "" {
			n, _ = strconv.Atoi(m[2])
		}
		applyTag(words, m[1], n)
	}
	return strings.Join(words, " ")
}

func applyTag(words []string, tag string, n int) {
	if len(words) == 0 {
		return
	}
	if n > len(words) {
		n = len(words)
	}
	switch tag {
	case "hex":
		convertBase(words, 16)
	case "bin":
		convertBase(words, 2)
	default:
		for i := len(words) - n; i < len(words); i++ {
			words[i] = changeCase(words[i], tag)
		}
	}
}

func convertBase(words []string, base int) {
	last := len(words) - 1
	v, err := strconv.ParseInt(words[last], base, 64)
	if err == nil {
		words[last] = strconv.FormatInt(v, 10)
	}
}
