package queue

import (
    "errors"
)

type Node struct {
    data int
    prev *Node
    next *Node
}

type Queue struct {
    size  int
    front *Node
	back  *Node
}

func (q Queue) Size() int {
	return q.size
}

func (q *Queue) Clear() {
    for q.size > 0 {
        q.Pop()
    }
}

func (q *Queue) Push(n int) {
    if q.front == nil {
		q.front = &Node{data: n}
        q.back = q.front
    } else {
        q.back.next = &Node{data: n}
        q.back = q.back.next
    }
    q.size++
}

func (q *Queue) Pop() error {
    if q.size == 0 {
        return errors.New("queue is empty")
    }
    q.size--
    q.front = q.front.next
    return nil
}

func (q Queue) Peek() (int, error) {
    if q.Size() == 0 {
        return 0, errors.New("queue is empty")
    }
    return q.front.data, nil
}
