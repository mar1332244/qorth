package main

import (
    "fmt"
)

type Node struct {
    data Token
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

func (q *Queue) Push(token Token) {
    if q.front == nil {
		q.front = &Node{data: token}
        q.back = q.front
    } else {
        q.back.next = &Node{data: token}
        q.back = q.back.next
    }
    q.Size++
}

func (q *Queue) Pop() error {
    if q.size == 0 {
        return fmt.Errorf("queue is empty")
    }
    q.size--
    q.front = q.front.next
    return nil
}

func (q Queue) Peek() (Token, error) {
    if q.Size() == 0 {
        return Token{}, fmt.Errorf("queue is empty")
    }
    return q.front.data, nil
}

func (q Queue) String() string {
    return "Queue<Token>"
}
