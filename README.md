# Interview Tasks

## palindrom.go 

test number if it is palindrom or not 

example 

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

```shell
123 is NOT Palindrom
123321 is Palindrom
987654321123456789 is Palindrom
```