package main

import (
	"fmt"
)

func romanToInteger(s string) int {
	roman2arabic := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	lastDigit := 4000
	arabic := 0
	c := []byte(s)
	for _, v := range c {
		digit := roman2arabic[string(v)]
		if lastDigit < digit {
			arabic -= 2 * lastDigit
		}
		lastDigit = digit
		arabic += lastDigit
	}
	return arabic
}

func main() {
	fmt.Printf("%#v\n", romanToInteger("DCXXI"))
}
