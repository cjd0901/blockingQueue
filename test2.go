package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	q := NewQueue2()

	c3 := Consumer2{}
	c4 := Consumer2{}
	go func() {
		c3.Pop(q, handle3)
	}()
	go func() {
		c4.Pop(q, handle4)
	}()

	for i := 0; i < 10; i++ {
		//if i == 5 {
		//	q.Close()
		//}
		if i == 1 {
			q.Add(i)
		}
		q.Add(i)
		// time.Sleep(time.Second)
	}

	select {
	case <- time.After(time.Minute):
	}
}

func handle3(i int) error {
	if i % 2 != 0 {
		return errors.New("非偶数")
	}
	fmt.Println("handle3 处理了", i)
	return nil
}

func handle4(i int) error {
	if i % 2 == 0 {
		return errors.New("非奇数")
	}
	fmt.Println("handle4 处理了", i)
	return nil
}