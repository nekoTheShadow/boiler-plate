package collections

import "testing"

func TestIsEmpty(t *testing.T) {
	d := NewDeque()
	if !d.IsEmpty() {
		t.FailNow()
	}

	d.AppendFirst(0)
	if d.IsEmpty() {
		t.FailNow()
	}

	d.PopLast()
	if !d.IsEmpty() {
		t.FailNow()
	}
}

func TestForStack(t *testing.T) {
	d := NewDeque()
	d.AppendLast(1)
	d.AppendLast(2)
	d.AppendLast(3)

	for _, expected := range []int{3, 2, 1} {
		if v := d.PopLast(); v != expected {
			t.FailNow()
		}
	}
}

func TestForQueue(t *testing.T) {
	d := NewDeque()
	d.AppendLast(1)
	d.AppendLast(2)
	d.AppendLast(3)

	for _, expected := range []int{1, 2, 3} {
		if v := d.PopFirst(); v != expected {
			t.FailNow()
		}
	}
}

func TestPeekFunctions(t *testing.T) {
	d := NewDeque()
	d.AppendLast(1)
	d.AppendLast(2)
	d.AppendLast(3)

	if d.PeekFirst() != 1 {
		t.FailNow()
	}

	if d.PeekLast() != 3 {
		t.FailNow()
	}
}
