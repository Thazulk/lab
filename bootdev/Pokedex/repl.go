package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", words[0])
	}
}

func cleanInput(text string) []string {
	// var words []string
	// for _, word := range strings.Fields(text) {
	// 	words = append(words, strings.TrimSpace(word))
	// }

	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}
