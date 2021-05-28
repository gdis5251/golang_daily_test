package main

import (
	"fmt"
	"sync"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func producer(c chan interface{}, arr []int, wg *sync.WaitGroup) {
	defer (*wg).Done()

	for _, num := range arr {
		c <- num
	}

	c <- "done"
}

func consumer(c chan interface{}, wg *sync.WaitGroup) {
	defer (*wg).Done()
forloop:
	for {
		select {
		case data := <- c:
			switch data.(type) {
			case string:
				break forloop
			case int:
				fmt.Println(data)
			}
		}
	}
}

func main() {
	//out := make(chan int)
	//// 由于创建的 chan 是无缓冲的，所以会一直阻塞在这一步。
	//out <- 2
	//go f1(out)

	c := make(chan interface{})
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go producer(c, []int{10, 20, 30, 40, 50, 60, 70, 80, 90}, wg)
	go consumer(c, wg)

	wg.Wait()
	fmt.Println("done")
}