package main

import (
	"fmt"
)

type fooer interface {
	foo() barer
}

type barer interface {
	bar() string
}

type fooImpl string

func (f fooImpl) foo() barer {
	return barImpl("ololo")
}

type barImpl string

func (b barImpl) bar() string {
	return string(b)
}

func main() {
	var f fooImpl
	lol(f)
}

func lol(f fooer) {
	fmt.Print(f.foo().bar())
}
