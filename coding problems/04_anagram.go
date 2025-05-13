// Problem 4: Check if Two Strings Are Anagrams
// Concept: Using slices and sorting with sort.Slice.

// Problem Statement:
// Write a function to check if two strings are anagrams (contain the same characters in any order).

// Sample Input 1: "listen", "silent" → Output: true
// Sample Input 2: "hello", "world" → Output: false

package main

import (
	"fmt"
	"sort"
)

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// Convert strings to rune slices for sorting
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	// Sort both slices
	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	// Compare sorted slices
	return string(runes1) == string(runes2)
}

func main() {
	fmt.Println(isAnagram("listen", "silent")) // true
	fmt.Println(isAnagram("hello", "world"))   // false
}
