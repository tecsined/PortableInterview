package usermanagement

import (
	"fmt"
	"strings"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (u User) ActiveAndContains(c string) (*User, bool) {
	email := strings.ToLower(u.Email)
	c = strings.ToLower(c)
	if strings.Contains(email, c) && u.IsActive {
		return &u, true
	}
	return nil, false
}

type Users []User

func (us Users) PrintFilteredActiveUsers(filterCriteria string) {
	printedHeader := false

	for _, user := range us {
		filteredUser, hasMatchingEmail := user.ActiveAndContains(filterCriteria)

		if hasMatchingEmail {
			if !printedHeader {
				fmt.Printf("Active users with %s in their email:\n", filterCriteria)
				printedHeader = true
			}
			fmt.Printf("ID: %d, Name: %s %s, Email: %s\n", filteredUser.Id, filteredUser.FirstName, filteredUser.LastName, filteredUser.Email)
		}
	}
}
