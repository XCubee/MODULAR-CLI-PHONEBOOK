// phonebook/logic.go
package phonebook

import (
	"fmt"
)

func ViewContacts(phoneBook map[string]string) {
	fmt.Println("The Contacts Include ")
	fmt.Println(phoneBook)
}