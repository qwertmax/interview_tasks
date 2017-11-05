package main

import (
	"fmt"
)

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

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

func main() {
	a := [][]int{[]int{9, 9, 4}, []int{6, 6, 8}, []int{2, 1, 1}}
	b := [][]int{[]int{3, 4, 5}, []int{3, 2, 6}, []int{2, 2, 1}}
	c := [][]int{[]int{1, 2}}

	fmt.Printf("%#v\n", longestIncreasingPath(a))
	fmt.Printf("%#v\n", longestIncreasingPath(b))
	fmt.Printf("%#v\n", longestIncreasingPath(c))
}
