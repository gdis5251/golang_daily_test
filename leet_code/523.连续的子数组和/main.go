package main

import (
	"fmt"
	"time"
)

// 近似与暴力法，会超时
func checkSubarraySum(nums []int, k int) bool {
	if k == 0 {
		return true
	} else if len(nums) == 0 {
		return false
	}

	dp := make([]int, len(nums) + 1)
	for index, num := range nums {
		dp[index + 1] += num + dp[index]
	}

	for i := 0; i < len(dp); i++ {
		for j := i + 2; j < len(dp); j++ {
			lastSum := dp[j] - dp[i]

			if (k == 0 && lastSum == 0) || (k != 0 && lastSum % k == 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	now := time.Now()
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
	fmt.Println(time.Now().Sub(now))
}