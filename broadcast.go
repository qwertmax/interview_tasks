package main

import (
	"fmt"
	"time"
)

func main() {
	value := make(chan interface{})
	quit := make(chan int)

	go func(value chan interface{}) {
		for {
			select {
			case v := <-value:
				fmt.Printf("case v: %v\n", v)
			case <-quit:
				fmt.Println("Gorutine exited")
				return
			}
		}
	}(value)

	fmt.Println("started")
	time.Sleep(2 * time.Second)

	go func() {
		for i := 0; i < 10; i++ {
			value <- interface{}(i + 1)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(400 * time.Millisecond)

	quit <- 0

	time.Sleep(300 * time.Millisecond)
}
