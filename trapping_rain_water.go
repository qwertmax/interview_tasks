package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
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
	// fmt.Printf("%#v\n", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	// load("data.txt")
	data := load("data.txt")
	// fmt.Printf("%#v\n", data)
	start := time.Now()
	trap(data)
	elapsed := time.Since(start)
	fmt.Printf("took %s", elapsed)

}

func load(file string) []int {

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	var data_int []int
	data := strings.Split(strings.Replace(string(dat), "\n", "", 1), ",")
	for _, i := range data {
		num, _ := strconv.Atoi(i)
		data_int = append(data_int, num)
	}
	// fmt.Printf("%#v\n", data)
	return data_int
}
