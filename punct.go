package main

import "regexp"

var spaceBefore = regexp.MustCompile(`\s+([.,!?:;]+)`)
var spaceAfter = regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`)

func fixPunctuation(text string) string {
	text = spaceBefore.ReplaceAllString(text, "$1")
	text = spaceAfter.ReplaceAllString(text, "$1 $2")
	return text
}
