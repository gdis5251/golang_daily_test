package main

import (
	"fmt"
	"strings"
)

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1 + s1, s2)
}

func isFlipedString2(s1 string, s2 string) bool {
	if len(s1) != len(s2) || (len(s1) == 1 && len(s2) == 1 && s1 != s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	for i := 0; i < len(s1); i++ {
		last := s2[len(s2) - 1]
		s2 = string(last) + s2[:len(s2) - 1]

		if s1 == s2 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(isFlipedString("waterbottle", "erbottlewat"))
}
