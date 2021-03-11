package main

import (
	"time"
)

type Queue2 struct {
	ch      chan int
}

type Consumer2 struct {}

func NewQueue2() Queue2 {
	return Queue2{ch: make(chan int, 10)}
}

func (q Queue2) Add(i int) {
	q.ch <- i
}

func (q Queue2) Close() {
	close(q.ch)
}

func (c Consumer2) Pop(q Queue2, handle func (i int) error) {
	for {
		n := <- q.ch
		err := handle(n)
		if err != nil {
			q.Add(n)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
