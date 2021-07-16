package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
层序遍历，组成链表
 */
func listOfDepth(tree *TreeNode) []*ListNode {
	queue := make([]*TreeNode, 0)
	res := make([]*ListNode, 0)
	if tree == nil {
		return res
	}

	queue = append(queue, tree)

	for len(queue) != 0 {
		nums := len(queue)
		subList := new(ListNode)
		cur := subList
		for i := 0; i < nums; i++ {
			cur.Next = &ListNode{Val: queue[0].Val}
			cur = cur.Next

			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		res = append(res, subList.Next)
	}

	return res
}
