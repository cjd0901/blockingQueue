package main

import (
	"errors"
	"fmt"
)

func main() {
	q := NewQueue()
	c1 := &Consumer{Id: 1}
	c2 := &Consumer{Id: 2}
	go func() {
		for i := 1; i < 100; i ++ {
			e := &Event{
				T: EVENT_ADD,
				K: fmt.Sprintf("%d", i),
				V: fmt.Sprintf("%d", i),
			}
			q.Add(e)
		}
	}()
	go func() {
		for i := 1; i < 500; i ++ {
			e := &Event{
				T: EVENT_UPDATE,
				K: fmt.Sprintf("%d", i),
				V: fmt.Sprintf("%d", i),
			}
			q.Add(e)
		}
	}()
	go c2.Pop(q, updateHandle)
	go c1.Pop(q, addHandle)
	select {

	}
}

func addHandle(e *Event) error {
	if e.T != EVENT_ADD {
		return errors.New("不支持新增操作")
	}
	fmt.Printf("执行了新增操作key:%s value:%s\n", e.K, e.V)
	return nil
}

func updateHandle(e *Event) error {
	if e.T != EVENT_UPDATE {
		return errors.New("不支持更新操作")
	}
	fmt.Printf("执行了更新操作key:%s value:%s\n", e.K, e.V)
	return nil
}
