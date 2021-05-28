package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))
	for index := range dp {
		dp[index] = make([]bool, len(s))
	}
	dp[0][0] = true

	ans := s[:1]
	maxLen := 1

	for right := 1; right < len(s); right++ {
		for left := 0; left < right; left++ {
			// 如果左右相等 --> 如果两个字符之间相隔 0 或 1 个字符，那么就是回文
			// 如果两个字符中间夹着的字符串是回文，那么他们也是回文
			if (s[left] == s[right]) && (right - left <= 2 || dp[right - 1][left + 1]) {
				dp[right][left] = true

				palindromeStrLen := right - left + 1
				if palindromeStrLen > maxLen {
					maxLen = palindromeStrLen
					ans = s[left:right + 1]
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}