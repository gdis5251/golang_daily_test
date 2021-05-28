package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
        return nil
    }

    listLen := getListLength(head)

    newHead := &ListNode{
        Val:  0,
        Next: head,
    }
    cur := newHead

    for step := 0; step < listLen - n; step++ {
        cur = cur.Next
    }

    cur.Next = cur.Next.Next

    return newHead.Next
}

func getListLength(head *ListNode) int {
    length := 0
    for head != nil {
        length++
        head = head.Next
    }

    return length
}