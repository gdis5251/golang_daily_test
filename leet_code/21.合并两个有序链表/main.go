package main

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }

    res := new(ListNode)

    cur1 := l1
    cur2 := l2
    cur3 := res


    for cur1 != nil && cur2 != nil {
        if cur1.Val < cur2.Val {
            tmp := ListNode{
                Val:  cur1.Val,
                Next: nil,
            }
            cur3.Next = &tmp
            cur1 = cur1.Next
        } else {
            tmp := ListNode{
                Val:  cur2.Val,
                Next: nil,
            }
            cur3.Next = &tmp
            cur2 = cur2.Next
        }

        cur3 = cur3.Next
    }

    if cur1 == nil {
        cur3.Next = cur2
    } else if cur2 == nil {
        cur3.Next = cur1
    }

    return res.Next
}
