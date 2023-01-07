package grind75

import (
	"testing"

	ds "github.com/davidn5013/leetcode/tools/ds"
)

func TestQueue(t *testing.T) {
	q := ds.NewMyQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Pop()
	q.Pop()
	ans := q.Pop()
	if ans != 3 {
		t.Errorf("Pop was %d; want 3", ans)
	}
}

func TestQueue2(t *testing.T) {
	q := ds.NewMyQueue()
	q.Push(1)
	q.Pop()
	ans := q.Empty()
	if ans != true {
		t.Errorf("Check if empty was %v; want true", ans)
	}
}

func TestQueue3(t *testing.T) {
	q := ds.NewMyQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Pop()
	ans := q.Peek()
	if ans != 2 {
		t.Errorf("Check if empty was %d; want 2", ans)
	}
}

func TestMyQueue4(t *testing.T) {
	q := ds.NewMyQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Pop()
	ans := q.Peek()
	if ans != 2 {
		t.Errorf("Check if empty was %d; want 2", ans)
	}
}
