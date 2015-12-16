package main

import "fmt"

var battle = make(chan string)
var i int = 0

func warrior(name string, done chan struct{}) {
	i++
	// fmt.Println("i:", i, "name: ", name)
	select {
	case opponent := <-battle:
		// fmt.Println("i:", i, "opponent \t\tname:", name)
		fmt.Printf("%s beat %s\n", name, opponent)
	case battle <- name:
		// fmt.Println("i:", i, "name \t\tname:", name)
		// I lost :-(
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	langs := []string{"Go", "C", "C++", "Java", "Perl", "Python"}
	for _, l := range langs {
		go warrior(l, done)
	}

	for _ = range langs {
		<-done
	}
}
