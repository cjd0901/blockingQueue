package main

import (
	"time"
)

type Queue2 struct {
	ch      chan int
	m       map[int]int
}

type Consumer2 struct {}

func NewQueue2() Queue2 {
	return Queue2{
		ch: make(chan int, 100),
		m: map[int]int{},
	}
}

func (q Queue2) Add(i int) {
	if _, ok := q.m[i]; ok {
		return
	}
	q.ch <- i
	q.m[i] = i
}

func (q Queue2) Close() {
	close(q.ch)
}

func (c Consumer2) Pop(q Queue2, handle func (i int) error) {
	for {
		n := <- q.ch
		delete(q.m, n)
		err := handle(n)
		if err != nil {
			q.Add(n)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
