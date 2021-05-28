package main

func longestPalindrome(s string) int {
	charCount := make(map[byte]int)

	for _, ch := range s {
		charCount[byte(ch)]++
	}

	res := 0

	for _, v := range charCount {
		if res & 1 != 0 && v & 1 != 0 {
			res += v - 1
		} else {
			res += v
		}
	}

	return res
}