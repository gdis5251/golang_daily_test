package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type QpsLimiter struct {
	BucketMap *sync.Map
	qps       int
	tickerMap *sync.Map
}

func NewQpsLimiter(qps int) *QpsLimiter {
	if qps <= 0 {
		qps = 1
	}

	bucket := new(sync.Map)
	ticker := new(sync.Map)
	ql := &QpsLimiter{
		BucketMap: bucket,
		qps:       qps,
		tickerMap: ticker,
	}

	return ql
}

func (ql *QpsLimiter) Add(api string) {
	if _, ok := ql.BucketMap.Load(api); !ok {
		bucket := make(chan int, 1)
		ql.BucketMap.Store(api, bucket)
		ticker := time.NewTicker(time.Second / time.Duration(ql.qps))
		ql.tickerMap.Store(api, ticker)
	}

	go ql.StartTicker(api)
}

func (ql *QpsLimiter) StartTicker(api string) {
	if bucket, ok := ql.BucketMap.Load(api); ok {
		if ticker, ok := ql.tickerMap.Load(api); ok {
			t := ticker.(*time.Ticker).C
			go func() {
				for range t {
					bucket.(chan int) <- 1
				}
			}()
		}
	}
}

func (ql *QpsLimiter) GetToken(api string) bool {
	if bucket, ok := ql.BucketMap.Load(api); ok {
		for range bucket.(chan int) {
			return true
		}
	}
	return false
}

func (ql *QpsLimiter) Stop(api string) {
	if ticker, ok := ql.tickerMap.Load(api); ok {
		t := ticker.(*time.Ticker)
		t.Stop()
	}
}

func main() {
	limiter := NewQpsLimiter(3)

	limiter.Add("123")
	go func() {
		for i := 0; i < 20; i++ {
			go func(limiter *QpsLimiter) {
				if limiter.GetToken("123") {
					fmt.Println("123----------" + strconv.FormatInt(time.Now().Unix(), 10))
				}
			}(limiter)
		}
	}()

	limiter.Add("456")
	go func() {
		for i := 0; i < 20; i++ {
			go func(limiter *QpsLimiter) {
				if limiter.GetToken("456") {
					fmt.Println("456----------" + strconv.FormatInt(time.Now().Unix(), 10))
				}
			}(limiter)
		}
	}()

	limiter.Add("789")
	go func() {
		for i := 0; i < 20; i++ {
			go func(limiter *QpsLimiter) {
				if limiter.GetToken("789") {
					fmt.Println("789----------" + strconv.FormatInt(time.Now().Unix(), 10))
				}
			}(limiter)
		}
	}()
	time.Sleep(time.Second * 10000)

	limiter.Stop("123")
	limiter.Stop("456")
	limiter.Stop("789")



	time.Sleep(time.Second * 10000)
}
