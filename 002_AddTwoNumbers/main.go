package main

import "fmt"

/*
 * Definition for singly-linked list.
 */
// ListNode :=
type ListNode struct {
	Val  int
	Next *ListNode
}

// Print :=
func (l *ListNode) Print() {
	fmt.Println("")
	if l == nil {
		fmt.Println("List is empty")
		return
	}

	fmt.Printf("%d ", l.Val)
	for l.Next != nil {
		l = l.Next
		fmt.Printf("%d ", l.Val)
	}
	fmt.Println("")

	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := ListNode{}
	current := &head
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10
		current.Next = &ListNode{
			Val: sum % 10,
		}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{
			Val: carry,
		}
	}

	return head.Next
}

func main() {
	/*
		Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
		Output: 7 -> 0 -> 8
		Explanation: 342 + 465 = 807.
	*/
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	/*
		Input: (5) + (5)
		Output: 0 -> 1
		Explanation: 342 + 465 = 807.
	*/
	// l1 := ListNode{
	// 	Val: 5,
	// }

	// l2 := ListNode{
	// 	Val: 5,
	// }

	addTwoNumbers(&l1, &l2).Print()
}
