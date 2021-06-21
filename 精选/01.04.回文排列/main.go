package main

func canPermutePalindrome(s string) bool {
	byteMap := make(map[int32]int)

	for _, c := range s {
		if _, ok := byteMap[c]; ok {
			byteMap[c]++
		} else {
			byteMap[c] = 1
		}
	}

	singleChar := 0
	for _, v := range byteMap {
		if v % 2 != 0 {
			singleChar++
		}
	}

	if singleChar > 1 {
		return false
	}

	return true
}
