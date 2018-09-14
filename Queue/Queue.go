package queue

import (
	"fmt"
	"sync"
)

// 队列元素
type Element interface{}

// 定义队列
type Queue struct {
	element []Element
	lock    sync.RWMutex
}

// 初始化队列
func (q *Queue) Init() *Queue {
	q.element = []Element{}
	return q
}

// 创建队列
func New() *Queue {
	return new(Queue).Init()
}

// 入队
func (q *Queue) Enqueue(e Element) {
	q.lock.Lock()
	q.element = append(q.element, e)
	q.lock.Unlock()
}

// 出队
func (q *Queue) Dequeue() Element {
	q.lock.Lock()
	if q.IsEmpty() {
		fmt.Println("queue is empty!")
		return nil
	}
	e := q.element[0]
	q.element = q.element[1:len(q.element)]
	q.lock.Unlock()
	return e
}

// 取队列第一个元素
func (q *Queue) Front() Element {
	q.lock.Lock()
	e := q.element[0]
	q.lock.Unlock()
	return e
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.element) == 0
}

// 获取队列的长度
func (q *Queue) Size() int {
	return len(q.element)
}

// 清空队列
func (q *Queue) Clear() bool {
	if q.IsEmpty() {
		fmt.Println("queue is empty!")
		return false
	}
	q.lock.Lock()
	for i := 0; i < q.Size(); i++ {
		q.element[i] = nil
	}
	q.element = nil
	q.lock.Unlock()
	return true
}
