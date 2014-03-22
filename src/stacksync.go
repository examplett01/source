package lab4

import "sync"

type SafeStack struct {
	mu   sync.Mutex
	top  *Element
	size int
}

func (ss *SafeStack) Len() int {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	return ss.size
}

func (ss *SafeStack) Push(value interface{}) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	ss.top = &Element{value, ss.top}
	ss.size++
}

func (ss *SafeStack) Pop() (value interface{}) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	if ss.size > 0 {
		value, ss.top = ss.top.value, ss.top.next
		ss.size--
		return
	}

	return nil
}
