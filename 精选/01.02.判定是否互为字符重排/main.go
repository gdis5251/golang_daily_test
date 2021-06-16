package main

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1Counter := generateCharCounter(s1)
	s2Counter := generateCharCounter(s2)

	return equalCounter(s1Counter, s2Counter)
}

func generateCharCounter(s string) map[byte]uint32 {
	charCounter := make(map[byte]uint32, 0)

	for _, c := range s {
		if _, ok := charCounter[byte(c)]; ok {
			charCounter[byte(c)]++
		} else {
			charCounter[byte(c)] = 1
		}
	}

	return charCounter
}

func equalCounter(s1Counter, s2Counter map[byte]uint32) bool {
	for c, count := range s1Counter {
		if count != s2Counter[c] {
			return false
		}
	}

	return true
}