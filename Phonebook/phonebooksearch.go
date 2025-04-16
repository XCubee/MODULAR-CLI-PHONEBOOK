package phonebook

import (
	"bufio"
	"fmt"
	"strings" 
)


func SearchContact(phoneBook map[string]string, scanner *bufio.Scanner) {
	fmt.Println("Enter the name you want to search:")
	scanner.Scan()
	search := strings.TrimSpace(scanner.Text()) // Trimming extra spaces

	
	if number, exists := phoneBook[search]; exists {
		fmt.Printf("Found! %s: %s\n", search, number)
	} else {
		fmt.Println("The name is not present in the directory.")
	}
}
