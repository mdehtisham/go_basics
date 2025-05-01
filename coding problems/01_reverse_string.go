/*
Problem 1: Reverse a String
Concept: Strings are immutable in Go. Use slices of rune to handle Unicode characters.

Problem Statement:
Write a function to reverse a given string.

Sample Input: "hello"
Sample Output: "olleh"
*/

package main

import "fmt"

func reverseString(s string) string {
	// Convert string to a slice of runes (to handle Unicode characters)
	runes := []rune(s)

	// Swap characters from start and end, moving toward the center
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert back to string
	return string(runes)
}

func main() {
	input := "hello"
	fmt.Println("Reversed:", reverseString(input)) // Output: "olleh"
}
