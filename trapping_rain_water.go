package main

import (
	"fmt"
)

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
