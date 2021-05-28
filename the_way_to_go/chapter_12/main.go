package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileHandler, err := os.Open("fileName.txt")
	if err != nil {
		fmt.Println("open file err: %v", err)
		return
	}

	defer fileHandler.Close()

	reader := bufio.NewReader(fileHandler)
	ioutil.ReadAll(reader)
}
