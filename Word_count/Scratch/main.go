package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a word: ")
	var word string
	fmt.Scanln(&word)

	count := 0
	for _, char := range word {
		if char >= 'a' && char <= 'z' {
			count++
		}
	}

	fmt.Printf("The word %s has %d letters.\n", word, count)
}
