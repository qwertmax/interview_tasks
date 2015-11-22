# Interview Tasks

##Table of Contents

- [Palindrom](#palindromgo)
- [Funny task for understanding methods and pointers](#go_fun)

## palindrom.go 

test number if it is palindrom or not 

### example:

```go
func main() {
	numbers := []int64{123, 123321, 987654321123456789}

	for _, val := range numbers {
		if val == reverce(val) {
			fmt.Printf("%d is Palindrom\n", val)
		} else {
			fmt.Printf("%d is NOT Palindrom\n", val)
		}
	}
}

func reverce(num int64) int64 {
	r := int64(0)
	for num != 0 {
		r = r*10 + num%10
		num = num / 10
	}

	return r
}
```

### Output:

```shell
123 is NOT Palindrom
123321 is Palindrom
987654321123456789 is Palindrom
```


## go fun

very funny implementation of go lang (good thing for interview question)

### example

```go
type T int

func (t T) Bar() { t++; println(t) }

func main() {
	var foo T = 1
	var fooPtr *T = &foo

	foo.Bar()
	(*fooPtr).Bar()
	fooPtr.Bar()
}
```

### Output

```shell
2
2
2
```


## go_fun_v2.go

very funny implementation of go lang with pointers (good thing for interview question)

### example

```go
package main

type T int

func (t *T) Bar() { *t++; println(*t) }

func main() {
	var foo T = 1
	var fooPtr *T = &foo

	foo.Bar()
	(*fooPtr).Bar()
	fooPtr.Bar()
}
```

### Output

```shell
2
3
4
```

## Lambda / Closure example in Go

### Example 

```go
package main

import "fmt"

func Adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	add5 := Adder(5)
	fmt.Println(add5(1))

	add8 := Adder(8)
	fmt.Println(add8(9))
}
```

### Output

```shell
6
17
```

## Data Race by Dave Chenny

[Data Race by Dave Chenny](http://dave.cheney.net/2015/11/18/wednesday-pop-quiz-spot-the-race)

### Example
```go
package main

import (
	"fmt"
	"time"
)

type RPC struct {
	result int
	done   chan struct{}
}

func (rpc *RPC) compute() {
	time.Sleep(time.Second) // strenuous computation intensifies
	rpc.result = 42
	close(rpc.done)
}

func (RPC) version() int {
	return 1 // never going to need to change this
}

func main() {
	rpc := &RPC{done: make(chan struct{})}

	go rpc.compute()         // kick off computation in the background
	version := rpc.version() // grab some other information while we're waiting
	<-rpc.done               // wait for computation to finish
	result := rpc.result

	fmt.Printf("RPC computation complete, result: %d, version: %d\n", result, version)
}
```


### Run command

```shell
go run --race data_race.go

```

### Output

```shell
go run --race data_race.go
==================
WARNING: DATA RACE
Write by goroutine 6:
  main.(*RPC).compute()
      github.com/qwertmax/interview_tasks/data_race.go:15 +0x3b

Previous read by main goroutine:
  main.main()
      github.com/qwertmax/interview_tasks/data_race.go:27 +0xfd

Goroutine 6 (running) created at:
  main.main()
      github.com/qwertmax/interview_tasks/data_race.go:26 +0xe6
==================
RPC computation complete, result: 42, version: 1
Found 1 data race(s)
exit status 66
```

