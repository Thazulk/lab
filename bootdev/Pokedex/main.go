package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a single REPL
	for {
		userInput := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		userInput.Scan()
		input := userInput.Text()
		output := cleanInput(input)
		fmt.Printf("Your command was: %s\n", output[0])

	}
}
