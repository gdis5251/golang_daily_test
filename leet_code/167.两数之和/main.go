package main

func twoSum(numbers []int, target int) []int {
	valueSlot := make(map[int]int, len(numbers))

	for index, num := range numbers {
		valueSlot[num] = index
	}

	for leftIndex, num := range numbers {
		if index, ok := valueSlot[target - num]; ok {
			return []int{leftIndex, index}
		}
	}

	return []int{}
}
