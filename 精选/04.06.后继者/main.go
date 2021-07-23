package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var resArr = make([]*TreeNode, 0)

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	midOrder(root)

	for index := range resArr {
		if resArr[index] == p && index < len(resArr) - 1 {
			return resArr[index + 1]
		}
	}
	return nil
}

func midOrder(root *TreeNode) {
	if root == nil {
		return
	}

	midOrder(root.Left)
	resArr = append(resArr, root)
	midOrder(root.Right)
}
