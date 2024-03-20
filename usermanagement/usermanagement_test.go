package usermanagement

import (
	"testing"
)

func TestActiveAndContains_ActiveUserWithEmailContainingCriteria(t *testing.T) {
	// Arrange 
	user := User{Id: 1, FirstName: "John", LastName: "Doe", Email: "johndoe@email.com", IsActive: true}
	criteria := "doe"

	// Act 
	filteredUser, isActive := user.ActiveAndContains(criteria)

	// Assert
	if !isActive {
		t.Errorf("Expected user to be active, but got inactive")
	}
	if filteredUser.Email != user.Email {
		t.Errorf("Expected filtered user to be the same user")
	}
}

func TestActiveAndContains_InactiveUserWithEmailContainingCriteria(t *testing.T) {
	// Arrange 
	user := User{Id: 1, FirstName: "John", LastName: "Doe", Email: "johndoe@email.com", IsActive: false}
	criteria := "doe"

	// Act
	_, isActive := user.ActiveAndContains(criteria)

	// Assert
	if isActive {
		t.Errorf("Expected user to be inactive, but got active")
	}
}

func TestActiveAndContains_NoCriteriaMatch(t *testing.T) {
	// Arrange 
	user := User{Id: 1, FirstName: "John", LastName: "Doe", Email: "johndoe@email.com", IsActive: true}
	criteria := "bar"

	// Act 
	filteredUser, isActive := user.ActiveAndContains(criteria)

	// Assert 
	if isActive {
		t.Errorf("Expected user to be inactive due to no criteria match")
	}
	if filteredUser != nil {
		t.Errorf("Expected no user to be returned for no criteria match")
	}
}
