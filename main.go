package main

import (
	"fmt"
	"reflect"
)

func main() {
	chan1 := make(chan int, 1)
	chan2 := make(chan string, 2)
	chan3 := make(chan float64, 1)

	// 异步监听
	go func() {
		for {
			// select 会等待某个管道就绪，若同时多个就绪，会随机选一个处理
			select {
			case params := <- chan1:
				fmt.Println(params)
			case <- chan2:
			case <- chan3:
			default:
				// do nothing
			}
		}
	}()

	reflect.SliceHeader{}
}
