package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	sub := dfs(root.Left) - dfs(root.Right)
	return sub <= 1 && sub >= -1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}
