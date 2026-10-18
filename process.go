package main

import "strings"

func Process(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = processLine(line)
	}
	return strings.Join(lines, "\n")
}

func processLine(text string) string {
	text = applyTags(text)
	return text
}
