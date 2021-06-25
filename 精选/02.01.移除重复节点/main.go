package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    nodeSet := make(map[int]interface{})
    vHead := new(ListNode)
    vHead.Next = head

    prev := vHead

    for cur := head; cur != nil; {
        if _, ok := nodeSet[cur.Val]; ok {
            prev.Next = cur.Next
            cur = cur.Next
        } else {
            nodeSet[cur.Val] = 1
            prev = prev.Next
            cur = cur.Next
        }
    }

    return vHead.Next
}

func main() {
    l1 := ListNode{
        Val:  1,
        Next: nil,
    }
    l2 := ListNode{
        Val:  2,
        Next: nil,
    }
    l3 := ListNode{
        Val:  3,
        Next: nil,
    }
    l4 := ListNode{
        Val:  4,
        Next: nil,
    }
    l5 := ListNode{
        Val:  5,
        Next: nil,
    }
    l6 := ListNode{
        Val:  6,
        Next: nil,
    }


    l1.Next = &l2
    l2.Next = &l3
    l3.Next = &l4
    l4.Next = &l5
    l5.Next = &l6

    removeDuplicateNodes(&l1)
}