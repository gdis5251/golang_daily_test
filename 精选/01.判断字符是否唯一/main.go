package main

func isUnique(astr string) bool {
	set := make(map[byte]uint8, len(astr))
	for _, c := range astr {
		if _, ok := set[byte(c)]; ok {
			return false
		}

		set[byte(c)] = 1
	}

	return true
}
