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
