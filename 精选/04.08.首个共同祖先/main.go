package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	lNode := lowestCommonAncestor(root.Left, p, q)
	rNode := lowestCommonAncestor(root.Right, p, q)

	if lNode != nil && rNode != nil {
		return root
	}

	if lNode == nil {
		return rNode
	}

	return lNode
}
