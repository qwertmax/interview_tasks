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
	_, err := os.Open("README.md1")
	if err != nil {
		if serr, ok := err.(*os.PathError); ok {
			fmt.Printf("%#v\n", serr)
		} else {
			fmt.Printf("default err\n%#v\n", err)
		}
	}

}
