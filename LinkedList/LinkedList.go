package LinkedList

import (
	"fmt"
	"sync"
)

type Node struct {
	data interface{}
	next *Node
}

// 定义链表结构
type LinkedList struct {
	head *Node
	lock sync.RWMutex
}

// 初始化链表
func Init() *LinkedList {
	l := new(LinkedList)
	l.head = nil
	return l
}

// 插入元素
func (l *LinkedList) InsertElem(i interface{}) {
	l.lock.Lock()
	data := &Node{data: i}
	if l.head == nil {
		l.head = data
		l.lock.Unlock()
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = data
	l.lock.Unlock()
}

// 获取指定索引的元素
func (l *LinkedList) GetElem(i int) interface{} {
	current := l.head
	for i > 0 && current.next != nil {
		i--
		current = current.next
	}
	return current.data
}

// 返回元素的索引
func (l *LinkedList) LocateElem(i interface{}) int {
	index := 0
	current := l.head
	for current.next != nil {
		if current.data == i {
			return index
		}
		index++
		current = current.next
	}
	return -1
}

// 判断链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

// 删除指定索引的元素
func (l *LinkedList) DeleteByIndex(i int) bool {
	if i < 0 {
		return false
	}
	if l.head == nil {
		return false
	}
	l.lock.Lock()
	if i == 0 {
		l.head = l.head.next
		l.lock.Unlock()
		return true
	}
	current := l.head
	for current.next != nil {
		if i == 0 {
			current.next = current.next.next
			l.lock.Unlock()
			return true
		}
		i--
		current = current.next
	}
	l.lock.Unlock()
	return false
}

// 删除指定值的元素
func (l *LinkedList) DeleteByValue(i interface{}) bool {
	if l.head == nil {
		return false
	}
	l.lock.Lock()
	if l.head.data == i {
		l.head = l.head.next
		l.lock.Unlock()
		return true
	}
	current := l.head
	for current.next != nil {
		if current.data == i {
			current.next = current.next.next
			l.lock.Unlock()
			return true
		}
		current = current.next
	}
	l.lock.Unlock()
	return false
}

// 清空链表
func (l *LinkedList) Clear() bool {
	l.lock.Lock()
	if l.head == nil {
		return false
	}
	l.head = nil
	l.lock.Unlock()
	return true
}

// 链表长度
func (l *LinkedList) Len() int {
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// 打印链表元素
func (l *LinkedList) Print() bool {
	if l.head == nil {
		fmt.Println("队列为空......")
		return false
	}
	current := l.head
	for current != nil {
		fmt.Print(current.data)
		current = current.next
	}
	return true
}
