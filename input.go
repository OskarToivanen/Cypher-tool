// TODO:
// Get the input data required for the operation

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getInput() string {
	for {
		reader := bufio.NewReader(os.Stdin)
		secretmessage, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return err.Error()
		}

		secretmessage = strings.TrimSpace(secretmessage)
		secretmessage = strings.ReplaceAll(secretmessage, " ", "")
		secretmessage = strings.ReplaceAll(secretmessage, "\t", "")
		secretmessage = strings.ReplaceAll(secretmessage, "\n", "")

		hasNonPrintable := false
		for i, c := range secretmessage {
			if !unicode.IsPrint(c) && c != '\n' {
				fmt.Printf("Found non-printable character: %q at index %d\n", c, i)
				hasNonPrintable = true
			}
		}
		if hasNonPrintable {
			fmt.Print("Please Enter a valid message.\n>")
		}

		emptyInput := false
		if secretmessage == "" {
			emptyInput = true
			fmt.Print("You entered an empty string.\n")
			fmt.Print("Please Enter a valid message.\n>")
		}

		if !emptyInput && !hasNonPrintable {
			return secretmessage
		}

	}
}

func getUserMessage() string {
	fmt.Print("Enter the message:\n> ")
	return getInput()
}

func getValidOperationInput() int {
	for {
		fmt.Print("Select operation (1/2):\n1. Encrypt.\n2. Decrypt.\n> ")
		input := getInput()

		// Try to convert input to integer and validate
		operation, err := strconv.Atoi(input)
		if err == nil && (operation == 1 || operation == 2) {
			return operation
		}
		fmt.Println("Invalid input, please enter 1 for Encrypt or 2 for Decrypt.")
	}
}

// Function to get a valid cypher input (ROT13/Reverse)
func getValidCypherInput() int {
	for {
		fmt.Print("Choose cypher:\n1. ROT13.\n2. Reverse.\n3. Shift6\n> ")
		input := getInput()

		// Try to convert input to integer and validate
		cypher, err := strconv.Atoi(input)
		if err == nil && (cypher == 1 || cypher == 2 || cypher == 3) {
			return cypher
		}
		fmt.Println("Invalid input, please enter 1 for ROT13 or 2 for Reverse or 3 for Shift6.")
	}
}
