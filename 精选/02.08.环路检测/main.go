package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	meetNode := hasCycle(head)
	// 如果不存在环，就返回空
	if meetNode == nil {
		return nil
	}

	// 以相遇节点为基准，把环拆开
	anotherHead := meetNode.Next
	meetNode.Next = nil

	return getIntersectionNode(head, anotherHead)
}

func hasCycle(head *ListNode) *ListNode {
	// 快慢指针先判断链表是否有环
	if head == nil {
		return nil
	}

	slowNode := head
	fastNode := head

	for  {
		fastNode = fastNode.Next
		if fastNode == nil {
			return nil
		}
		if fastNode == slowNode {
			return fastNode
		}
		fastNode = fastNode.Next
		if fastNode == nil {
			return nil
		}
		if fastNode == slowNode {
			return fastNode
		}

		slowNode = slowNode.Next
		if slowNode == nil {
			return nil
		}
		if fastNode == slowNode {
			return slowNode
		}
	}
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
