/*
Problem 3: Count Vowels in a String
Concept: Using map for lookups, rune iteration, and case conversion.

Problem Statement:
Write a function to count the number of vowels (a, e, i, o, u) in a string, case-insensitive.

Sample Input: "Hello World!"
Sample Output: 3 (e, o, o)
*/
package main

import (
	"fmt"
	"unicode"
)

func countVowels(s string) int {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	count := 0
	for _, c := range s {
		// Convert character to lowercase and check if it's a vowel
		if vowels[unicode.ToLower(c)] {
			count++
		}
	}
	return count
}

func main() {
	input := "Hello World!"
	fmt.Println("Vowel count:", countVowels(input)) // Output: 3
}

/*
Explanation:

Map for Lookups: A map[rune]bool stores vowels for O(1) membership checks.

Case Insensitivity: unicode.ToLower(c) converts characters to lowercase to handle both A and a as vowels.

Rune Iteration: Iterating over s using range gives each Unicode character as a rune.
*/
