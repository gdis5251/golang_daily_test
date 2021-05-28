package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := ""
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]

		isSame := true
		jumpOut := false
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				jumpOut = true
				break
			}
			if strs[j][i] != ch {
				isSame = false
			}
		}

		if jumpOut {
			break
		}

		if isSame {
			res += string(ch)
		} else {
			break
		}
	}

	return res
}
