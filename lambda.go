package main

import "fmt"

func Adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	add5 := Adder(5)
	val := add5(1)

	fmt.Println(val)

	add8 := Adder(8)
	val = add8(9)
	fmt.Println(val)
}
