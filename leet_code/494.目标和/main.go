package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	res := 0

	//dfs(curSum, target, index, &res, nums, '+')
	//dfs(curSum, target, index, &res, nums, '-')

	dfsBetter(nums, target, 0, 0, &res)

	return res
}

func dfsBetter(nums []int, target, sum, index int, res *int) {
	if index == len(nums) {
		if sum == target {
			*res++
		}
	} else {
		dfsBetter(nums, target, sum + nums[index], index + 1, res)
		dfsBetter(nums, target, sum - nums[index], index + 1, res)
	}
}

func dfs(curSum, target, index int, res *int, nums []int, op byte) {
	if op == '+' {
		curSum += nums[index]
	} else {
		curSum -= nums[index]
	}
	index++

	if curSum == target && index == len(nums){
		*res += 1
		return
	} else if index == len(nums) {
		return
	} else  {
		dfs(curSum, target, index, res, nums, '+')
		dfs(curSum, target, index, res, nums, '-')
	}
}



func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
