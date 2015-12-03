package main

import (
	"fmt"
)

// Person is a strcut for person object
type Person struct {
	Fname string
	Lname string
	Address
}

type Address struct {
	addr  string
	city  string
	state string
}

func (a *Address) Foo() {
	a.addr = "q"
}

func (p *Person) Foo() {
	p.Address.addr = "w"
}

func main() {
	p := Person{}
	p.Fname = "Maxim"
	p.Lname = "Tishchenko"
	p.Address.addr = "Marksa ave"

	fmt.Println(p.Address.addr)
	p.Foo()
	fmt.Println(p.Address.addr)
}
