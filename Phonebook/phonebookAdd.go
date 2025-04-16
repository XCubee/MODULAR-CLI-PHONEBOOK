package phonebook

import (
	"bufio"
	"fmt"
)

func AddContact(phoneBook map[string]string, scanner *bufio.Scanner) {
	fmt.Println("Enter the name ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Enter the phone number ")
	scanner.Scan()
	PhoneNmber := scanner.Text()
	phoneBook[name] = PhoneNmber

	fmt.Println("Contact Added Successfully ")
}
