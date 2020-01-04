package main

import (
	"github.com/rnreekez/LeetCode/kit"
)

// ListNode used for LC compile
type ListNode = kit.ListNode

// iterative
func reverseListV1(head *ListNode) *ListNode {
	var previous *ListNode = nil
	var next *ListNode = nil
	current := head

	for current != nil {
		next = current.Next
		current.Next = previous
		previous, current = current, next
	}

	return previous
}

// recursive
func reverseList(head *ListNode) *ListNode {
	return reverseListHelper(head, nil)
}

func reverseListHelper(head *ListNode, prev *ListNode) *ListNode {
	if head == nil {
		return prev
	}

	next := head.Next
	head.Next = prev

	return reverseListHelper(next, head)
}

func main() {

	input := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	reverseList(&input).Print()
}
