package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return ConnectLeftAndRightTreeNode(0, len(nums) - 1, nums)
}

func ConnectLeftAndRightTreeNode(start, end int, nums []int) *TreeNode {
	if start > end {
		return nil
	} else if start == end {
		return &TreeNode{
			Val: nums[start],
		}
	} else if start + 1 == end {
		return &TreeNode{
			Val:  nums[end],
			Left: &TreeNode{
				Val: nums[start],
			},
		}
	}

	mid := (start + end) / 2

	return &TreeNode{
		Val:   nums[mid],
		Left:  ConnectLeftAndRightTreeNode(start, mid - 1, nums),
		Right: ConnectLeftAndRightTreeNode(mid + 1, end, nums),
	}
}
