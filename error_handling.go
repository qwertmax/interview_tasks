package main

import (
	"fmt"
	"os"
)

type SyntaxError struct {
	msg    string // description of error
	Offset int64  // error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string { return e.msg }

func main() {

	v1("README.md")
	v2("README.md_ERROR")
	// v3("README.md_ERROR")
}

func v1(fileName string) {
	fmt.Println("v1")

	_, err := os.Open(fileName)
	if err != nil {
		if serr, ok := err.(*os.PathError); ok {
			fmt.Printf("%#v\n", serr)
		} else {
			fmt.Printf("default err\n%#v\n", err)
		}
	}

	fmt.Println("END v1\n")
}

func v2(fileName string) {
	fmt.Println("v2")

	_, err := os.Open(fileName)
	if err != nil {
		if serr, ok := err.(*os.SyscallError); ok {
			fmt.Printf("%#v\n", serr)
		} else {
			fmt.Printf("default err\n%#v\n", err)
		}
	}

	fmt.Println("END v2\n")
}
