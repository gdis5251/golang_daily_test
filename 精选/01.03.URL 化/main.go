package main

import "strings"

func replaceSpaces(S string, length int) string {
	res := S[:length]

	return strings.ReplaceAll(res, " ", "%20")
}
