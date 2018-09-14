package Stack

import "sync"

type Element struct {
	data interface{}
	next *Element
}

// 定义栈结构
type Stack struct {
	top  *Element
	lock sync.RWMutex
}

// 初始化栈
func Init() *Stack {
	s := new(Stack)
	s.top = nil
	return s
}

// 入栈
func (s *Stack) Push(i interface{}) bool {
	data := &Element{data: i}
	s.lock.Lock()
	if s.top == nil {
		s.top = data
		s.lock.Unlock()
		return true
	}
	data.next = s.top
	s.top = data
	s.lock.Unlock()
	return true
}

// 出栈
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	if s.top == nil {
		return -1
	}
	var data interface{}
	data = s.top.data
	if s.top.next != nil {
		s.top = s.top.next
	} else {
		s.top = nil
	}
	s.lock.Unlock()
	return data
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	if s.top == nil {
		return true
	}
	return false
}

// 获取栈顶元素
func (s *Stack) GetTop() interface{} {
	if s.top != nil {
		return s.top.data
	}
	return -1
}

// 返回栈的大小
func (s *Stack) Size() int {
	if s.top == nil {
		return 0
	}
	count := 0
	current := s.top
	for current != nil {
		count++
		current = current.next
	}
	return count
}
