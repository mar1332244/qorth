package queue

import (
    "fmt"
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
        return fmt.Errorf("queue is empty")
    }
    q.size--
    q.front = q.front.next
    return nil
}

func (q Queue) Peek() (int, error) {
    if q.Size() == 0 {
        return 0, fmt.Errorf("queue is empty")
    }
    return q.front.data, nil
}

func (q Queue) String() string {
    return "Queue<Token>"
}
