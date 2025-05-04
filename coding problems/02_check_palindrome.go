/*
Problem 2: Check if a String is a Palindrome
Concept: Compare the original string with its reversed version.

Problem Statement:
Write a function to check if a string reads the same backward as forward (e.g., "madam" is a palindrome).

Sample Input 1: "madam" → Output: true
Sample Input 2: "hello" → Output: false
*/

package main

import "fmt"

func isPalindrome(s string) bool {
	reversed := ReverseString(s) // Reuse the reverse function from Problem 1
	return s == reversed
}

func main() {
	fmt.Println(isPalindrome("madam")) // true
	fmt.Println(isPalindrome("hello")) // false
}
