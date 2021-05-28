package main

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")

	res := ""

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			res += strings.TrimSpace(words[i]) + " "
		}
	}

	if len(res) > 1 {
		res = res[:len(res) - 1]
	}

	return res
}

func main() {
	reverseWords("a good   example")
}