package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	var stack Stack
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	if size := stack.Size(); size != 10 {
		t.Error("expected: 10, got:", size)
	}
	for stack.Size() > 0 {
		n, err := stack.Peek()
		if err != nil {
			t.Error("unexpected error:", err)
		}
		if n < 0 {
			t.Error("unexpacted value of -1")
		}
		if err = stack.Pop(); err != nil {
			t.Error("unexpected error:", err)
		}
	}
	if size := stack.Size(); size != 0 {
		t.Error("expected: 0, got:", size)
	}
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	if size := stack.Size(); size != 10 {
		t.Error("expected: 10, got:", size)
	}
	for stack.Size() > 0 {
		n, err := stack.Peek()
		if err != nil {
			t.Error("unexpected error:", err)
		}
		if n < 0 {
			t.Error("unexpacted value of -1")
		}
		if err = stack.Pop(); err != nil {
			t.Error("unexpected error:", err)
		}
	}
}
