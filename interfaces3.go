package main

import "fmt"

type Welcome interface {
	Hi()
}
type Farewell interface {
	Bye()
}

type Polite interface {
	Welcome
	Farewell
}
type Max struct {
	Name string
}
type Alexa struct {
	Name string
}

func (m Max) Hi() {
	fmt.Println("Hi", m.Name)
}
func (a Alexa) Hi() {
	fmt.Println("Hi", a.Name)
}
func (m Max) Bye() {
	fmt.Println("Bye", m.Name)
}

func (a Alexa) Bye() {
	fmt.Println("Bye", a.Name)
}

func main() {
	var max Polite = Max{"Max"}
	var alexa Polite = Alexa{"Alexa"}

	max.Hi()
	alexa.Hi()

	max.Bye()
	alexa.Bye()
}
