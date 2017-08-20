package main

import (
	"fmt"
)

func reverse(x int) int {
	new_int := 0
	sign := false
	if x < 0 {
		sign = true
		x = -x
	}

	for x > 0 {
		remainder := x % 10
		new_int *= 10
		new_int += remainder
		x /= 10
	}

	if sign {
		new_int = -new_int
	}

	if new_int > 2147483647 || new_int < -2147483647 {
		return 0
	}

	return new_int
}

func main() {
	fmt.Printf("%d\n", reverse(123))
	fmt.Printf("%d\n", reverse(-321))
	fmt.Printf("%d\n", reverse(1534236469))
	fmt.Printf("%d\n", reverse(900000))
	fmt.Printf("%d\n", reverse(-2147483648))
	fmt.Printf("%d\n", reverse(-2147483412))
}
