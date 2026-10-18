package main

import "regexp"

var quoteRe = regexp.MustCompile(`'\s+(.*?)\s+'`)

func fixQuotes(text string) string {
	return quoteRe.ReplaceAllString(text, "'$1'")
}
