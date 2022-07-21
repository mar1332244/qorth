package stack

import (
	"errors"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
    size int
    top  *Node
}

func (s Stack) Size() int {
    return s.size
}

func (s *Stack) Push(value int) {
    s.size++
    s.top = &Node{data: value, next: s.top}
}

func (s *Stack) Pop() error {
    if s.size == 0 {
        return errors.New("stack is empty")
    }
    s.size--
    s.top = s.top.next
	return nil
}

func (s *Stack) Peek() (int, error) {
    if s.size == 0 {
        return -1, errors.New("stack is empty")
    }
    return s.top.data, nil
}
