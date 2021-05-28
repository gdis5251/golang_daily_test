package main

import "math"

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	res := 0

	for _, num := range nums {
		if num != 0 {
			count ++
		} else {
			res = int(math.Max(float64(res), float64(count)))
			count = 0
		}
	}

	res = int(math.Max(float64(res), float64(count)))

	return res
}