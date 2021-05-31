package main

import "fmt"

func max(lhs, rhs int) int {
	if lhs >= rhs {
		return lhs
	}

	return rhs
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for index, num := range nums {
		dp[index] = num
	}
	res := nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] += max(dp[i - 1], 0)

		res = max(dp[i], res)
	}
	return res
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}