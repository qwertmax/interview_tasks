package main

import (
  "fmt"
)

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

func main() {
  a := []int{1, 2, 1, 3, 2, 5}
  fmt.Printf("[1,2,1,3,2,5] -> %#v\n", singleNumber(a))
}
