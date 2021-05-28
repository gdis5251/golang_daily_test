package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    slow := head
    fast := head

    for  {
        slow = (*slow).Next
        if slow == nil {
            return false
        }
        fast = (*fast).Next
        if fast == nil {
            return false
        }
        fast = (*fast).Next
        if fast == nil {
            return false
        }

        if slow == fast {
            return true
        }

    }
}
