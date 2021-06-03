package main

func findMaxLength(nums []int) int {
	mp := make(map[int]int)
	maxLen := 0
	counter := 0
	mp[counter] = -1

	// bce 优化
	_ = nums[len(nums) - 1]

	for index, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}

		if preIndex, ok := mp[counter]; ok {
			if length := index - preIndex; length > maxLen {
				maxLen = length
			}
		} else {
			mp[counter] = index
		}
	}

	return maxLen
}
