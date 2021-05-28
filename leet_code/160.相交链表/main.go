package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	left := headA
	right := headB

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
		if left == nil || right == nil {
			return nil
		}
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