package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := getListLength(headA)
	lenB := getListLength(headB)

	curA := headA
	curB := headB

	// 先让两个链表距离相等
	if lenA >= lenB {
		for i := 0; i < lenA - lenB; i++ {
			curA = curA.Next
		}
	} else {
		for i := 0; i < lenB - lenA; i++ {
			curB = curB.Next
		}
	}

	// 一起走
	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}

		curA = curA.Next
		curB = curB.Next
	}

	return nil
}

func getListLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}

	return length
}
