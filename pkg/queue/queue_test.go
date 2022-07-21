package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	var queue Queue
	for i := 0; i < 10; i++ {
		queue.Push(i)
	}
	if size := queue.Size(); size != 10 {
		t.Error("expected: 10, got:", size)
	}
	for queue.Size() > 0 {
		n, err := queue.Peek()
		if err != nil {
			t.Error("unexpected error:", err)
		}
		if n < 0 {
			t.Error("unexpacted value of -1")
		}
		if err = queue.Pop(); err != nil {
			t.Error("unexpected error:", err)
		}
	}
	if size := queue.Size(); size != 0 {
		t.Error("expected: 0, got:", size)
	}
	for i := 0; i < 10; i++ {
		queue.Push(i)
	}
	if size := queue.Size(); size != 10 {
		t.Error("expected: 10, got:", size)
	}
	for queue.Size() > 0 {
		n, err := queue.Peek()
		if err != nil {
			t.Error("unexpected error:", err)
		}
		if n < 0 {
			t.Error("unexpacted value of -1")
		}
		if err = queue.Pop(); err != nil {
			t.Error("unexpected error:", err)
		}
	}
}

func TestQueue_Clear(t *testing.T) {
	var queue Queue
	for i := 0; i < 100; i++ {
		queue.Push(i)
	}
	queue.Clear()
	if size := queue.Size(); size != 0 {
		t.Error("expected: 0, got:", size)
	}
}
