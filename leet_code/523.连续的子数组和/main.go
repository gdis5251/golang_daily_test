package main

import (
	"fmt"
	"time"
)

func checkSubarraySum(nums []int, k int) bool {
	if k == 0 {
		return true
	} else if len(nums) == 0 {
		return false
	}

	hashMap := make(map[int]int)

	hashMap[0] = -1
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sum %= k

		if _, ok := hashMap[sum]; ok {
			if i - hashMap[sum] > 1 {
				return true
			}
		} else {
			hashMap[sum] = i
		}
	}

	return false
}

func main() {
	now := time.Now()
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
	fmt.Println(time.Now().Sub(now))
}