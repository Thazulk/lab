package main

import "strings"

func cleanInput(text string) []string {
	// var words []string
	// for _, word := range strings.Fields(text) {
	// 	words = append(words, strings.TrimSpace(word))
	// }

	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}
