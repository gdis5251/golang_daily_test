package main

import (
	"fmt"
	"reflect"
	"time"
)

func f(arr []int) {
	go func() {
		for{
			fmt.Println(time.Now().Unix())
			time.Sleep(time.Second)
		}
	}()

	//time.Sleep(time.Second*100)
}

func main() {
	//t := time.NewTicker(time.Second / 3)
	//ch := t.C
	//
	//go func() {
	//	for range ch {
	//	fmt.Println("123----------" + strconv.FormatInt(time.Now().Unix(), 10))
	//}}()
	//
	//go func() {
	//	for range ch {
	//		fmt.Println("456----------" + strconv.FormatInt(time.Now().Unix(), 10))
	//	}
	//}()
	//
	//time.Sleep(time.Second * 10)
	//t.Stop()

	fmt.Println(4 % 5)
}
