package collections

import "testing"

func TestIncrementAndGet(t *testing.T) {
	c := NewCounter()
	c.Increment("1")
	c.Increment(2)
	c.Increment(2)

	assert(t, c.Get("0"), 0)
	assert(t, c.Get(0), 0)
	assert(t, c.Get("1"), 1)
	assert(t, c.Get(2), 2)
}

func TestKeys(t *testing.T) {
	c := NewCounter()
	c.Increment("1")
	c.Increment(2)
	c.Increment(2)

	keys := c.Keys()
	if len(keys) != 2 {
		t.Errorf("length of Counter.Keys() --- expected=2, but actual=%d", len(keys))
		return
	}

	ok1 := false
	ok2 := false
	for _, key := range keys {
		if key == "1" {
			ok1 = true
		}
		if key == 2 {
			ok2 = true
		}
	}

	if !ok1 {
		t.Error(`Counter.Keys() do not contains "1"`)
	}
	if !ok2 {
		t.Error(`Counter.Keys() do not contains 2`)
	}
}

func assert(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("expected=%d, but actual=%d", actual, expected)
	}
}
