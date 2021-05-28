package main

import "math"

func findMin(nums []int) int {
	res := math.MinInt32

	for _, num := range nums {
		if res > num {
			res = num
		}
	}

	return res
}
