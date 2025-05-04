/*
Problem 1: Reverse a String
Concept: Strings are immutable in Go. Use slices of rune to handle Unicode characters.

Problem Statement:
Write a function to reverse a given string.

Sample Input: "hello"
Sample Output: "olleh"
*/

package main

func ReverseString(s string) string {
	// Convert string to a slice of runes (to handle Unicode characters)
	runes := []rune(s)

	// Swap characters from start and end, moving toward the center
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert back to string
	return string(runes)
}

// func main() {
// 	input := "hello"
// 	fmt.Println("Reversed:", ReverseString(input)) // Output: "olleh"
// }

/*
Explanation:

Runes: Go uses rune to represent Unicode code points (e.g., emojis or non-ASCII characters). Converting a string to []rune ensures proper handling of multi-byte characters.

Immutability: Strings cannot be modified directly, so we convert them to a slice of runes for manipulation.

Loop Logic: Use two pointers (i starting from the beginning, j from the end) to swap characters until they meet in the middle.
*/
