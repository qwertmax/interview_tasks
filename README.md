# Go (Golang) interesting examples and tricky implementations

##Table of Contents

- [Palindrom](#palindrom)
- [Funny task for understanding methods and pointers](#go-fun)
- [Funny task for understanding methods and pointers version 2](#go-fun-v2)
- [Lambda/Closure example in go](#lambda--closure-example-in-go)
- [Data Race by Dave Chenny](#data-race-by-dave-chenny)
- [Channel Usage example](#channel-usage-example)
- [net/http explanation](#explanation-of-nethttp)
- [Error handling in Go](#error-handling-in-go)
- [Benchmarks](#benchmarks)
- [Defer](#defer)
- [Defer2](#defer2)
- [Sort vs quick sort](#sort-vs-quick-sort)
- [Map over users](#map-over-users)
- [Interfaces](#interfaces)

## Palindrom

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


## go fun v2

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

### How can we fix it? 

Pretty easty, we need to set version as pounter as well.

#### was

```go
func (RPC) version() int {
	return 1
}
```

#### Must be
```go
func (*RPC) version() int {
	return 1
}
```


## Channel Usage example

```go
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
				fmt.Println("case <-quit")
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
```

### Output

```shell
started
case v: 1
case v: 2
case v: 3
case v: 4
Gorutine exited
```


## Explanation of net/http

```go
package main

import (
	"net/http"
)

type myHandler struct {
	s string
}

func (myH *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my handler: " + myH.s))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Main)

	myH := &myHandler{
		s: "test my handler",
	}
	mux.Handle("/my-handler", myH)

	http.ListenAndServe(":3000", mux)
}

func Main(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("main max"))
}
```

## Error handling in Go


## Benchmarks

### Simple benchmark

you need just run command like this:

```shell
cd benchmark
go test -v -bench .
```

```go
func BenchmarkStrConcat(b *testing.B) {
	var sa = "hello"
	var sb = "world"
	var sc string
	for i := 0; i < b.N; i++ {
		sc += sa + sb
	}
	sa = sc
}

func BenchmarkFmtConcat(b *testing.B) {
	var sa = "hello"
	var sb = "world"
	var sc string
	for i := 0; i < b.N; i++ {
		sc = fmt.Sprintf("%s%s%s", sc, sa, sb)
	}
	sa = sc
}
```

```shell
BenchmarkStrConcat-8	  200000	    195032 ns/op
BenchmarkFmtConcat-8	  100000	    241140 ns/op
ok  	github.com/qwertmax/interview_tasks/benchmark	63.442s
```

##defer

```go
type number int

func print(n number)   { fmt.Printf("print\t %v\t %d\n", &n, n) }
func pprint(n *number) { fmt.Printf("pprint\t %v\t% d\n", n, *n) }

func main() {
	var n number
	defer print(n)
	defer pprint(&n)
	defer func() { print(n) }()
	defer func() { pprint(&n) }()

	n = 3
}
```

Console output

```shell
pprint	 0xc820060028	 3
print	 0xc820060060	 3
pprint	 0xc820060028	 3
print	 0xc820060078	 0
```

##defer2

```go
type T struct {
	ID int
}

func (t *T) PrintID() {
	fmt.Println(t.ID)
}

func F1() {
	ts := []T{{1}, {2}, {3}}
	for _, t := range ts {
		defer t.PrintID()
	}
}

func F2() {
	ts := []*T{&T{1}, &T{2}, &T{3}}
	for _, t := range ts {
		defer t.PrintID()
	}
}

func main() {
	fmt.Println("F1()")
	F1()
	fmt.Println()
	fmt.Println("F2()")
	F2()
}
```

Console output

```shell
F1()
3
3
3

F2()
3
2
1
```

##Sort vs quick sort

run benchmark tests

```shell
GOMAXPROCS=8 go test -bench=.
testing: warning: no tests to run
PASS
BenchmarkQsort-8      	       2	 630965101 ns/op
BenchmarkNativeQsort-8	       2	 726137348 ns/op
BenchmarkSort-8       	       5	 311290221 ns/op
ok  	github.com/qwertmax/interview_tasks/qsort	7.238s
```

looks like NativeQsort is the best solution for sort!
but it is not so simple as we could think. 

###Let's try to look at memory allocations!

```shell
GOMAXPROCS=8 go test -bench=. -benchmem
testing: warning: no tests to run
PASS
BenchmarkQsort-8      	       2	 626935514 ns/op	42648592 B/op	 1332781 allocs/op
BenchmarkNativeQsort-8	       2	 726719612 ns/op	852610040 B/op	 3239965 allocs/op
BenchmarkSort-8       	       5	 305692503 ns/op	      32 B/op	       1 allocs/op
ok  	github.com/qwertmax/interview_tasks/qsort	7.143s
```

###WOW. Did you see that? It looks like right now we know why it was so fast.


##Map over users

interesting implementation of map and useing it from any kind of function

```go
type user struct {
	name, email string
}

func (u *user) ChangeEmail(email string) {
	u.email = email
}

func (u user) String() string {
	return fmt.Sprintf("%s (%s)", u.name, u.email)
}

type userGroup struct {
	users map[int]*user
}

func (ug userGroup) String() string {
	output := "["
	for key, val := range ug.users {
		output += fmt.Sprintf("%d: {%s}; ", key, val)
	}

	output += "]"
	return output
}

// main magic goes here
func (ug *userGroup) mapOverUsers(fn func(u *user)) {
	for _, user := range ug.users {
		fn(user)
	}
}

func main() {
	ug := userGroup{
		map[int]*user{
			0: &user{
				name:  "Max",
				email: "1@ex.com"},
			1: &user{
				name:  "Nati",
				email: "2@ex.com"},
			2: &user{
				name:  "Alex",
				email: "3@ex.com"},
		},
	}
	fmt.Println(ug)

	ug.mapOverUsers(func(u *user) {
		u.ChangeEmail("new email")
	})

	fmt.Println(ug)
}
```

### Output

```shell
[0: {Max (1@ex.com)}; 1: {Nati (2@ex.com)}; 2: {Alex (3@ex.com)}; ]
[0: {Max (new email)}; 1: {Nati (new email)}; 2: {Alex (new email)}; ]
```


##Interfaces

```go
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

```

you can think that you can use 

```go
func (f fooImpl) foo() barImpl {
	return barImpl("ololo")
}

```

but this is wrong, becauce <mark>type fooer interface</mark> require <mark>foo() barer</mark> not a barImpl, that's why you have to use <mark>func (f fooImpl) foo() barer {</mark>
