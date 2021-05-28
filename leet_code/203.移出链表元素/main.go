package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	prevHead := &ListNode{
		Val:  0,
		Next: head,
	}
	prev := prevHead

	for cur := head; cur != nil;{
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = prev.Next
		}

		cur = cur.Next
	}

	return prevHead.Next
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	fmt.Println(removeElements(head, 3))
}