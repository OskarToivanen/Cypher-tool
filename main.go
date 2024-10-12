// Main logic, envoking other functions
/*TODO:
1. Trim input
2. Input validitation
3. Encrypt in reverse works
*/
package main

import (
	"fmt"
	"strings"
)

// MAIN program //
func main() {
	fmt.Printf("Welcome to the Cypher Tool!\n\n")

	// Loop the program until bool = true
	var exit bool = false
	for !exit {
		operation := getValidOperationInput()
		cypher := getValidCypherInput()
		message := getUserMessage()          // Get the user message func.
		message = strings.TrimSpace(message) // Trim message to remove extra space
		var result string

		if operation == 1 {
			switch cypher {
			case 1:
				result = rot13(message)
				fmt.Printf("Encrypted message using ROT13: %s\n\n", result)
			case 2:
				result = reverse(message)
				fmt.Printf("Encrypted message using reverse: %s\n\n", result)
			case 3:
				result = shift6(message)
				fmt.Printf("Encrypted message using shift6: %s\n\n", result)
			default:
				fmt.Printf("Incorrect value\n\n")
				continue
			}
		} else if operation == 2 {
			switch cypher {
			case 1:
				result = rot13(message)
				fmt.Printf("Decrypted message using ROT13: %s\n\n", result)
			case 2:
				result = reverse(message)
				fmt.Printf("Decrypted message using reverse: %s\n\n", result)
			case 3: 
				result = shift6decrypt(message)
				fmt.Printf("Decrypted message using shift6: %s\n\n", result)
			default:
				fmt.Printf("Incorrect value\n\n")
				continue
			}
			

		} else {
			fmt.Println("Invalid operation! Please select 1 or 2")
			continue
		}

		// Prompt the user to exit the program.
		var exitInput int // Variable for exit boolean
		fmt.Printf("Do you want to close the program? (1. Yes / 2. No): ")
		fmt.Scan(&exitInput)
		if exitInput == 1 {
			exit = true
		}
	}

}
