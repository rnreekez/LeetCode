package main

import (
	"github.com/rnreekez/LeetCode/kit"
)

// ListNode used for LC compile
type ListNode = kit.ListNode

func deleteNodeV1(node *ListNode) {
	current := node

	for current != nil {
		next := current.Next
		if next.Next == nil {
			current.Next = nil
		}

		current.Val = next.Val
		current = current.Next
	}
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}
	node1 := ListNode{Val: 1}
	node9 := ListNode{Val: 9}

	node4.Next = &node5
	node5.Next = &node1
	node1.Next = &node9

	deleteNode(&node5)
	node4.Print()
}
