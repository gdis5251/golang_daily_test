package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minLenth := len(nums) + 1
	sum := 0


	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum += nums[j]

			if sum >= target {
				minLenth = int(math.Min(float64(j - i + 1), float64(minLenth)))
				sum = 0
				break
			}
		}
		sum = 0
	}

	if minLenth == len(nums) + 1 {
		return 0
	}
	return minLenth
}

func main() {
	fmt.Println(minSubArrayLen(11, []int{1,1,1,1,1,1,1,1}))
}