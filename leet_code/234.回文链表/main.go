package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	stack := make([]int, 0)
	length := getListLen(head)
	if length == 1 {
		return true
	}

	cur := head

	for i := 0; i < length / 2; i++ {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}

	if length % 2 != 0 {
		cur = cur.Next
	}

	for cur != nil {
		if stack[len(stack) - 1] != cur.Val {
			return false
		}

		stack = stack[:len(stack) - 1]
		cur = cur.Next
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

func getListLen(head *ListNode) int {
	length := 0

	for head != nil {
		length++
		head = head.Next
	}

	return length
}
