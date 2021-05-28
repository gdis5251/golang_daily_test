package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}



func detectCycle(head *ListNode) *ListNode {
    node, ok := hasCycle(head)
    if !ok {
        return nil
    }

    left := head
    right := node.Next
    node.Next = nil

    leftLen := getListLength(left)
    rightLen := getListLength(right)

    if leftLen > rightLen {
        stepNum := leftLen - rightLen
        for step := 0; step < stepNum; step++ {
            left = left.Next
        }
    } else {
        stepNum := rightLen - leftLen
        for step := 0; step < stepNum; step++ {
            right = right.Next
        }
    }

    for left != right {
        left = left.Next
        right = right.Next
    }

    return left
}

func getListLength(head *ListNode) int {
    length := 0
    for head != nil {
        length++
        head = head.Next
    }

    return length
}

func hasCycle(head *ListNode) (*ListNode, bool) {
    if head == nil {
        return nil, false
    }

    slow := head
    fast := head

    for  {
        slow = (*slow).Next
        if slow == nil {
            return nil, false
        }
        fast = (*fast).Next
        if fast == nil {
            return nil, false
        }
        fast = (*fast).Next
        if fast == nil {
            return nil, false
        }

        if slow == fast {
            return slow, true
        }

    }
}