package collections

import (
	"strings"
	"testing"
)

func TestHeapq(t *testing.T) {
	heapq := NewHeapq()
	heapq.Push(15, "o")
	heapq.Push(10, "j")
	heapq.Push(22, "v")
	heapq.Push(23, "w")
	heapq.Push(9, "i")
	heapq.Push(6, "f")
	heapq.Push(20, "t")
	heapq.Push(21, "u")
	heapq.Push(14, "n")
	heapq.Push(8, "h")
	heapq.Push(13, "m")
	heapq.Push(19, "s")
	heapq.Push(1, "a")
	heapq.Push(4, "d")
	heapq.Push(25, "y")
	heapq.Push(11, "k")
	heapq.Push(17, "q")
	heapq.Push(2, "b")
	heapq.Push(12, "l")
	heapq.Push(7, "g")
	heapq.Push(18, "r")
	heapq.Push(24, "x")
	heapq.Push(3, "c")
	heapq.Push(26, "z")
	heapq.Push(5, "e")
	heapq.Push(16, "p")

	s := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	for _, c := range s {
		d, ok := heapq.Pop().(string)
		if !ok {
			t.FailNow()
		}

		if c != d {
			t.FailNow()
		}
	}
}

func TestIsEmpty(t *testing.T) {
	heapq := NewHeapq()
	if !heapq.isEmpty() {
		t.FailNow()
	}

	heapq.Push(1, "a")
	heapq.Push(2, "b")
	if heapq.isEmpty() {
		t.FailNow()
	}

	heapq.Pop()
	heapq.Pop()
	if !heapq.isEmpty() {
		t.FailNow()
	}
}
