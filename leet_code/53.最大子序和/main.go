package main

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))

	// bce 优化
	_ = nums[len(nums) - 1]
	for index, num := range nums {
		dp[index] = num
	}

	res := dp[0]

	_ = dp[len(dp) - 1]
	for i := 1; i < len(dp); i++ {
		dp[i] += max(dp[i - 1], 0)

		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}

func max(rhs, lhs int) int {
	if rhs > lhs {
		return rhs
	}
	return lhs
}