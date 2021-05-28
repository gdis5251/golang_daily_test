package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	listLen := getListLen(head)
	k = k % listLen


	for i := 0; i < k; i++ {
		head = lastNodeToFirst(head)
	}

	return head
}

func lastNodeToFirst(head *ListNode) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head

	preLastNode := preHead
	lastNode := head

	for lastNode.Next != nil {
		preLastNode = lastNode
		lastNode = lastNode.Next
	}

	preLastNode.Next = nil
	lastNode.Next = head
	head = lastNode

	return head
}

func getListLen(head *ListNode) int {
	length := 0

	for head != nil {
		length++
		head = head.Next
	}

	return length
}
