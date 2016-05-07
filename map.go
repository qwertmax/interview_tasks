package main

import (
	"fmt"
)

type user struct {
	name, email string
}

func (u *user) ChangeEmail(email string) {
	u.email = email
}

func (u user) String() string {
	return fmt.Sprintf("%s (%s)", u.name, u.email)
}

type userGroup struct {
	users map[int]*user
}

func (ug userGroup) String() string {
	output := "["
	for key, val := range ug.users {
		output += fmt.Sprintf("%d: {%s}; ", key, val)
	}

	output += "]"
	return output
}

func main() {
	ug := userGroup{
		map[int]*user{
			0: &user{
				name:  "Max",
				email: "1@ex.com"},
			1: &user{
				name:  "Nati",
				email: "2@ex.com"},
			2: &user{
				name:  "Alex",
				email: "3@ex.com"},
		},
	}

	fmt.Println(ug)
}
