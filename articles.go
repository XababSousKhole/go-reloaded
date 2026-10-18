package main

import "regexp"

var articleRe = regexp.MustCompile(`\b([aA])(\s+)([aeiouhAEIOUH])`)

func fixArticles(text string) string {
	return articleRe.ReplaceAllString(text, "${1}n${2}${3}")
}
