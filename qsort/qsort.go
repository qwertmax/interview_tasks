package qsort

import (
	"math/rand"
	"sort"
)

type Interface interface {
	sort.Interface

	Partition(i int) (left Interface, right Interface)
}

type IntSlice []int

func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IntSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is IntSlice) Len() int {
	return len(is)
}

func (is IntSlice) Partition(i int) (left Interface, right Interface) {
	return IntSlice(is[:i]), IntSlice(is[i+1:])
}

func Qsort(a Interface, prng *rand.Rand) Interface {
	if a.Len() < 2 {
		return a
	}

	left, right := 0, a.Len()-1

	pivotIndex := prng.Int() % a.Len()
	a.Swap(pivotIndex, right)

	for i := 0; i < a.Len(); i++ {
		if a.Less(i, right) {

			a.Swap(i, left)
			left++
		}
	}

	a.Swap(left, right)

	leftSide, rightSide := a.Partition(left)
	Qsort(leftSide, prng)
	Qsort(rightSide, prng)

	return a
}

func QuickSort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}
