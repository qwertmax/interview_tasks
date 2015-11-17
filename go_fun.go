package main

type T int

func (t T) Bar() { t++; println(t) }

func main() {
	var foo T = 1
	var fooPtr *T = &foo

	foo.Bar()
	(*fooPtr).Bar()
	fooPtr.Bar()
}
