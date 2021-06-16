package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// "aabcccccaaa"
func compressString(S string) string {
	if S == "" {
		return ""
	}

	byteCollecter := make(map[uint8]uint32)
	byteCollecter[S[0]] = 1

	res := new(bytes.Buffer)

	for i := 1; i < len(S); i++ {
		if _, ok := byteCollecter[S[i]]; ok {
			byteCollecter[S[i]]++
		} else {
			byteCollecter[S[i]] = 1

			res.WriteString(string(S[i-1]))
			res.WriteString(strconv.Itoa(int(byteCollecter[S[i-1]])))
			delete(byteCollecter, S[i-1])
		}
	}

	for k, v := range byteCollecter {
		res.WriteString(string(k))
		res.WriteString(strconv.Itoa(int(v)))
	}

	if len(res.String()) > len(S) {
		return S
	}

	return res.String()
}

func main() {
	fmt.Println(compressString("abbccd"))
}
