package main

import (
	"bufio"
	"fmt"
	"modular-cli-phonebook/phonebook"

	"os"
)

func main() {
	phoneBook := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Phone Book CLI")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contact")
		fmt.Println("3. Search Contact")
		fmt.Println("4. Exit")
		fmt.Println("5. Delete Contact")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			phonebook.AddContact(phoneBook, scanner)
		case "2":
			phonebook.ViewContacts(phoneBook)
		case "3":
			phonebook.SearchContact(phoneBook, scanner)
		case "4":
			fmt.Println("Exiting Phone Book CLI. Goodbye!")
			return
		case "5":
			phonebook.DeleteContact(phoneBook, scanner)
		default:
			fmt.Println("Invalid choice! Please enter a valid option.")
		}
	}
}
