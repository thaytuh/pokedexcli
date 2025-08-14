package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		rawText := scanner.Text()
		cleanText := cleanInput(rawText)
		fmt.Printf("Your command was: %v\n", cleanText[0])
	}
}

func cleanInput(text string) []string {
	cleanText := strings.Fields(text)
	for i := range cleanText {
		cleanText[i] = strings.ToLower(cleanText[i])
	
	}
	return cleanText
}