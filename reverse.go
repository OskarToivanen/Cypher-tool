package main

import "strings"

func reverse(s string) string {
	// Variable to store the result
	var result strings.Builder

	// Iterate over each character in the string
	for _, char := range s {
		// Check if it's a lowercase letter
		if char >= 'a' && char <= 'z' {
			// Calculate the reverse alphabet character for lowercase
			reversed := 'z' - (char - 'a')
			result.WriteRune(reversed)
		} else if char >= 'A' && char <= 'Z' {
			// Calculate the reverse alphabet character for uppercase
			reversed := 'Z' - (char - 'A')
			result.WriteRune(reversed)
		} else {
			// If it's not a letter, keep it unchanged
			result.WriteRune(char)
		}
	}

	// Return the final encrypted string
	return result.String()
}
