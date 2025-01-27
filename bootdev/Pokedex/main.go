package main

import (
	"fmt"
	"strings"
)

func CleanInput(text string) []string {
	var words []string
	for _, word := range strings.Fields(text) {
		words = append(words, strings.TrimSpace(word))
	}
	return words
}

func main() {
	fmt.Println("Hello world")
}
