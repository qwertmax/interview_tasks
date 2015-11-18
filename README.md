# Interview Tasks

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


## go_fun.go

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
	val := add5(1)

	fmt.Println(val)

	add8 := Adder(8)
	val = add8(9)
	fmt.Println(val)
}
```

### Output

```shell
6
17
```