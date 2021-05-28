package main

import (
	"encoding/json"
	"fmt"
)

type utest struct {
	A float64 `json:"a"`
}
func main() {
	test := "{\"a\": 11.11}"
	a := new(utest)

	json.Unmarshal([]byte(test), a)

	fmt.Printf("%+v", a)

}
