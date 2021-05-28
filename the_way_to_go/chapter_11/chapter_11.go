package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Triangle struct {
	length float32
	high   float32
}

func (t Triangle) Area() float32 {
	return (t.length * t.high) / 2
}

func main() {
	var s Shaper = Triangle{
		length: 3.1,
		high:   2.3,
	}

	fmt.Println(s.Area())
}
