package main

import "fmt"

type number int

func print(n number)   { fmt.Printf("print\t %h\t %d\n", &n, n) }
func pprint(n *number) { fmt.Printf("pprint\t %h\t% d\n", n, *n) }

func main() {
	var n number
	defer print(n)
	defer pprint(&n)
	defer func() { print(n) }()
	defer func() { pprint(&n) }()

	n = 3
}
