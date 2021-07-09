package main

import "fmt"

func main() {
	//tmp:= "111"
	//tmp += "&12312378172839172893719827389172389172389172893182734&req_id=__REQUEST_TD__"
	//
	//if tmp[(len(tmp)-len("&req_id=__REQUEST_TD__")):] == "&req_id=__REQUEST_TD__" {
	//	fmt.Println("yeah")
	//}
	//
	//test, _ := json.Marshal(tmp[(len(tmp)-len("&req_id=__REQUEST_TD__")):])
	//
	//var test2 string
	//
	//json.Unmarshal(test, &test2)
	//
	//fmt.Println(string(test))
	//fmt.Println(string(test2))

	//testMap := make(map[string]int)
	//testMap["a"] = 1
	//fn(testMap)
	//
	//fmt.Println(testMap)

	fmt.Println((true && (true|| false)))
	fmt.Println(!(true && (true|| false)))
	fmt.Println(true || (true && false))
	fmt.Println(true || (true && false))
}

func fn(testMap map[string]int) {
	testMap["b"] = 2
}
