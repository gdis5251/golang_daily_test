package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	oddHead := &ListNode{}
	oddCur := oddHead
	evenHead := &ListNode{}
	evenCur := evenHead

	index := 1

	for cur := head; cur != nil; {
		if index % 2 != 0 {
			oddCur.Next = cur
			oddCur = oddCur.Next
			cur = cur.Next
			oddCur.Next = nil
		} else {
			evenCur.Next = cur
			evenCur = evenCur.Next
			cur = cur.Next
			evenCur.Next = nil
		}
		index++
	}

	oddCur.Next = evenHead.Next

	return oddHead.Next
}
