package main

func moveZeroes(nums []int)  {
	slowIndex := 0
	fastIndex := 0

	for fastIndex < len(nums) {
		if nums[fastIndex] != 0 {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
		fastIndex++
	}

	for slowIndex < len(nums) {
		nums[slowIndex] = 0
		slowIndex++
	}
}
