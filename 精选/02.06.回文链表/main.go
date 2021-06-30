package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    length := getListLength(head)

    stack := make([]int, length / 2)
    cur := head

    for i := 0; i < length / 2; i++ {
        stack[i] = cur.Val
        cur = cur.Next
    }

    if length % 2 != 0 {
        cur = cur.Next
    }

    for cur != nil {
        if cur.Val != stack[len(stack) - 1] {
            return false
        } else {
            stack = stack[:len(stack) - 1]
            cur = cur.Next
        }
    }

    return true
}

func getListLength(head *ListNode) int {
    length := 0
    for head != nil {
        length++
        head = head.Next
    }

    return length
}