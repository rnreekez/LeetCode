package kit

import (
	"fmt"
)

// ListNode :=
type ListNode struct {
	Val  int
	Next *ListNode
}

// Print :=
func (l *ListNode) Print() {
	if l == nil {
		fmt.Println("List is empty")
		return
	}

	fmt.Printf("%d", l.Val)
	for l.Next != nil {
		fmt.Printf("->")
		l = l.Next
		fmt.Printf("%d", l.Val)
	}

	fmt.Printf("->NULL")

	fmt.Println("")
	return
}
