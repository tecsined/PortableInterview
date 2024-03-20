package main

import (
	"portable.int/usermanagement"
)

func main() {
	filterCriteria := "J"
	usermanagement.ListUsers.PrintFilteredActiveUsers(filterCriteria)
}
