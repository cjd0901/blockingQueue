package main

import (
	"sync"
)

const (
	EVENT_ADD = "add"
	EVENT_UPDATE = "update"
)

// 消费者
type Consumer struct {
	Id     int32 // id
}

// 事件队列
type Queue struct {
	sync.Mutex
	*sync.Cond
	Es     []*Event // 事件列表
	Closed bool // 是否关闭
	M      map[*Event]int // 用于事件去重的map
}

// 事件结构体
type Event struct {
	T      string // 事件的类型
	K      string // 事件操作的键
	V      string // 事件操作的值
}

func NewQueue() *Queue {
	q := &Queue{
		M: map[*Event]int{},
	}
	q.Cond = sync.NewCond(&q.Mutex)
	return q
}

// 向队列中添加事件
func (q *Queue) Add(e *Event) {
	// 判断队列是否已经被关闭
	if q.Closed {
		panic("the queue has been closed")
	}
	q.Cond.L.Lock()
	// 事件去重
	if _, ok := q.M[e]; ok {
		q.Cond.L.Unlock()
		return
	}
	// 插入事件
	q.Es = append(q.Es, e)
	q.M[e] = 1
	// 插入完成后执行通知操作
	q.Cond.Broadcast()
	q.Cond.L.Unlock()
}

// 关闭事件队列
func (q *Queue) Close() {
	q.Closed = true
	q.M= nil
	q.Es = nil
}

// 消费者获取事件并消费
func (c *Consumer) Pop(q *Queue, handle func(e *Event) error) {
	for {
		// 判断队列是否已经被关闭
		if q.Closed {
			panic("the queue has been closed")
		}

		q.Cond.L.Lock()
		// 如果队列为空，执行等待通知操作
		if len(q.Es) == 0 {
			q.Cond.Wait()
		}
		if len(q.Es) ==0 {
			continue
		}
		// 从队列中取出事件
		e := q.Es[0]
		q.Es = q.Es[1:]
		delete(q.M, e)
		q.Cond.L.Unlock()

		// 处理事件，如果报错就重新插入该事件
		err := handle(e)
		if err != nil {
			q.Add(e)
		}
	}
}

