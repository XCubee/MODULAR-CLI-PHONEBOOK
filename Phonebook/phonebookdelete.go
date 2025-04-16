package phonebook

import (
	"bufio"
	"fmt"
	"strings"
)
func DeleteContact(phoneBook map[string]string, scanner *bufio.Scanner) {
	fmt.Println("Enter the name you want to delete ")
	scanner.Scan()
	deleted := strings.TrimSpace(scanner.Text())
	if _, exists := phoneBook[deleted]; exists {
		delete(phoneBook, deleted)
		fmt.Println("Deleted Successfully ")
	} else {
		fmt.Println("Contact not found.")
	}
}