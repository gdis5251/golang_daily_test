package main

import "strings"

func reverse(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	for index := range words {
		words[index] = reverse(words[index])
	}

	return strings.Join(words, " ")
}
