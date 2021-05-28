package main

func removeDuplicates(nums []int) int {
	slowIndex := 0
	fastIndex := 0

	for fastIndex < len(nums) {
		if nums[slowIndex] == nums[fastIndex] {
			fastIndex++
		} else {
			slowIndex++
			nums[slowIndex] = nums[fastIndex]
		}
	}

	return slowIndex + 1
}
