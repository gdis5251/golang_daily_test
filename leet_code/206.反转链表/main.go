package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    resHead := &ListNode{
		Val:  0,
		Next: nil,
	}

    for head != nil {
        cur := &ListNode{
			Val:  head.Val,
			Next: resHead.Next,
		}
		resHead.Next = cur

        head = head.Next
    }

    return resHead
}
