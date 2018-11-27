# Go (Golang) interesting examples and tricky implementations

## Table of Contents

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
- [Two Sum](#two-sum)
- [Reverse](#reverse)
- [Add Two Numbers](#add-two-numbers)
- [Merge k sorted lists](#merge-k-sorted-lists)
- [Median of Two Sorted Arrays](#median-of-two-sorted-arrays)
- [Roman to Integer](#roman-to-integer)
- [Trapping Rain Water](#trapping-rain-water)
- [Maximum Gap](#maximum-gap)
- [Longest Increasing Path in a Matrix](#longest-increasing-path-in-a-matrix)
- [Longest Valid Parentheses](#longest-valid-parentheses)
- [Single Number 3](#single-number-3)

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

but this is wrong, becauce <strong>type fooer interface</strong> require <strong>foo() barer</strong> not a <strong>barImpl</strong>, that's why you have to use <strong>func (f fooImpl) foo() barer {</strong>


## Problems from leetcode.com

## Two Sum

```go
package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	indexs := make([]int, 2)
	hash := map[int]int{}

	for i := range nums {
		hash[target-nums[i]] = i
	}

	for i := range nums {
		index, ok := hash[nums[i]]
		if ok {
			if i == index {
				continue
			}
			indexs[0] = index
			indexs[1] = i
			sort.Ints(indexs)
			break
		}
		continue
	}

	return indexs
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	fmt.Printf("%#v\n", result)

}
```

## Reverse

```go
package main

import (
	"fmt"
)

func reverse(x int) int {
	new_int := 0
	sign := false
	if x < 0 {
		sign = true
		x = -x
	}

	for x > 0 {
		remainder := x % 10
		new_int *= 10
		new_int += remainder
		x /= 10
	}

	if sign {
		new_int = -new_int
	}

	if new_int > 2147483647 || new_int < -2147483647 {
		return 0
	}

	return new_int
}

func main() {
	fmt.Printf("%d\n", reverse(123))
	fmt.Printf("%d\n", reverse(-321))
	fmt.Printf("%d\n", reverse(1534236469))
	fmt.Printf("%d\n", reverse(900000))
	fmt.Printf("%d\n", reverse(-2147483648))
	fmt.Printf("%d\n", reverse(-2147483412))
}
```

## Add Two Numbers

```go
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	results := &ListNode{}
	node := results
	node1 := l1
	node2 := l2

	overten := false

	for node1 != nil || node2 != nil {

		tmp := 0

		if node1 != nil {
			tmp = tmp + node1.Val
			node1 = node1.Next
		}

		if node2 != nil {
			tmp = tmp + node2.Val
			node2 = node2.Next
		}
		if overten {
			tmp++
		}

		if tmp >= 10 {
			overten = true
			tmp -= 10
		} else {
			overten = false
		}

		node.Val = tmp

		if node1 != nil || node2 != nil {
			node.Next = &ListNode{}
			node = node.Next
		}
	}

	if overten {
		node.Next = &ListNode{}
		node = node.Next
		node.Val = 1
	}

	return results
}

func main() {
	res := addTwoNumbers(
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		})

	for res.Next != nil {
		fmt.Printf("%d -> ", res.Val)
		res = res.Next
	}
	fmt.Printf("%d", res.Val)
}
```

## Merge k sorted lists

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}
	return mergeTwoLists(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val < l2.Val {
		head = l1
		head.Next = mergeTwoLists(l1.Next, l2)
	} else {
		head = l2
		head.Next = mergeTwoLists(l2.Next, l1)
	}
	return head
}
```

## Median of Two Sorted Arrays

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	total := l1 + l2
	if total%2 == 1 {
		return float64(findKthNum(nums1, nums2, total/2+1))
	} else {
		return 0.5 * float64(findKthNum(nums1, nums2, total/2)+findKthNum(nums1, nums2, total/2+1))
	}
}

func findKthNum(nums1 []int, nums2 []int, k int) (val int) {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		return findKthNum(nums2, nums1, k)
	}
	if l1 == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}
	var pa, pb int
	if l1 < k/2 {
		pa = l1
	} else {
		pa = k / 2
	}
	pb = k - pa
	if nums1[pa-1] < nums2[pb-1] {
		return findKthNum(nums1[pa:], nums2, k-pa)
	} else if nums1[pa-1] > nums2[pb-1] {
		return findKthNum(nums1, nums2[pb:], k-pb)
	} else {
		return nums1[pa-1]
	}
}

func main() {
	res := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Printf("%#v\n", res)
	res = findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Printf("%#v\n", res)
}
```

## Roman to Integer

```go
func romanToInteger(s string) int {
	roman2arabic := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	lastDigit := 4000
	arabic := 0
	c := []byte(s)
	for _, v := range c {
		digit := roman2arabic[string(v)]
		if lastDigit < digit {
			arabic -= 2 * lastDigit
		}
		lastDigit = digit
		arabic += lastDigit
	}
	return arabic
}

func main() {
	fmt.Printf("%#v\n", romanToInteger("DCXXI"))
}
```

### Trapping Rain Water

```go
func trap(height []int) int {
	left, right, secHeight, area := 0, len(height)-1, 0, 0
	for left < right {
		if height[left] < height[right] {
			secHeight = max(secHeight, height[left])
			area += secHeight - height[left]
			left++
		} else {
			secHeight = max(secHeight, height[right])
			area += secHeight - height[right]
			right--
		}
	}
	return area
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Printf("%#v\n", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
```

### Maximum Gap

```go
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	max, min := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	bucketSize := (max-min)/len(nums) + 1
	bucket := make([][]int, len(nums))
	for i := range nums {
		idx := (nums[i] - min) / bucketSize
		if len(bucket[idx]) == 0 {
			bucket[idx] = []int{nums[i], nums[i]}
		} else {
			if nums[i] < bucket[idx][0] {
				bucket[idx][0] = nums[i]
			}
			if nums[i] > bucket[idx][1] {
				bucket[idx][1] = nums[i]
			}
		}
	}
	ret := bucket[0][1] - bucket[0][0]
	pre := 0
	for i := 1; i < len(nums); i++ {
		if len(bucket[i]) == 0 {
			continue
		}
		if bucket[i][0]-bucket[pre][1] > ret {
			ret = bucket[i][0] - bucket[pre][1]
		}
		pre = i
	}
	return ret

}
```


### Longest Increasing Path in a Matrix

*Examples:*

```
nums = [
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
```

Return `4`

The longest increasing path is `[1, 2, 6, 9]`.

*more examples:*

```
nums = [
  [3,4,5],
  [3,2,6],
  [2,2,1]
]
```

Return `4`

The longest increasing path is `[3, 4, 5, 6]`.



```go
func longestIncreasingPath(matrix [][]int) int {
	row := len(matrix)
	var col int
	if row > 0 {
		col = len(matrix[0])
	} else {
		col = 0
	}

	var path [][]int
	var visited [][]bool
	path = make([][]int, row)
	visited = make([][]bool, row)
	for i := 0; i < row; i++ {
		path[i] = make([]int, col)
		visited[i] = make([]bool, col)
	}

	ret := int(0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			ret = max(ret, helper(matrix, path, row, col, i, j, visited))
		}
	}

	return ret
}
```

and helper

```go
func helper(matrix [][]int, path [][]int, row, col, r, c int, visited [][]bool) int {
	if path[r][c] > 0 {
		return path[r][c]
	}
	visited[r][c] = true

	ret := int(0)

	if r > 0 && !visited[r-1][c] && matrix[r][c] < matrix[r-1][c] {
		ret = max(ret, helper(matrix, path, row, col, r-1, c, visited))
	}

	if r < row-1 && !visited[r+1][c] && matrix[r][c] < matrix[r+1][c] {
		ret = max(ret, helper(matrix, path, row, col, r+1, c, visited))
	}

	if c > 0 && !visited[r][c-1] && matrix[r][c] < matrix[r][c-1] {
		ret = max(ret, helper(matrix, path, row, col, r, c-1, visited))
	}

	if c < col-1 && !visited[r][c+1] && matrix[r][c] < matrix[r][c+1] {
		ret = max(ret, helper(matrix, path, row, col, r, c+1, visited))
	}

	visited[r][c] = false
	path[r][c] = ret + 1
	return path[r][c]
}
```


### Longest Valid Parentheses


```go
func longestValidParentheses(s string) int {
	bytes := []byte(s)
	if len(bytes) < 2 {
		return 0
	}
	lengthList := make([]int, len(bytes))
	var max int
	for i := 1; i < len(bytes); i++ {
		if bytes[i] == ')' {
			j := i - lengthList[i-1] - 1
			if j >= 0 && bytes[j] == '(' {
				lengthList[i] = lengthList[i-1] + 2
				if j-1 >= 0 {
					lengthList[i] += lengthList[j-1]
				}
			}
		}
		if lengthList[i] > max {
			max = lengthList[i]
		}
	}
	return max
}
```

### Single Number 3

```go
func singleNumber(numbers []int) []int {
  var allox int
  for i := range numbers {
    allox ^= numbers[i]
  }
  mask := 1
  for ; mask&allox == 0; mask <<= 1 {
  }
  var ret1, ret2 int
  for i := range numbers {
    if numbers[i]&mask != 0 {
      ret1 ^= numbers[i]
    } else {
      ret2 ^= numbers[i]
    }
  }
  return []int{ret1, ret2}
}
```
