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
