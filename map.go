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

type userGroup struct {
	users map[int]*user
}

func main() {
	ug := userGroup{user{name: "Max", email: "1@ex.com"}}
}
