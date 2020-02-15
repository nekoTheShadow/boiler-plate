package floats

import "testing"

func TestMax(t *testing.T) {
	assertTrue(t, Max(1.0) == 1.0)
	assertTrue(t, Max(1.0, 2.0) == 2.0)
	assertTrue(t, Max(1.0, 2.0, 3.0) == 3.0)
}

func TestMin(t *testing.T) {
	assertTrue(t, Min(1.0) == 1.0)
	assertTrue(t, Min(1.0, 2.0) == 1.0)
	assertTrue(t, Min(1.0, 2.0, 3.0) == 1.0)
}

func assertTrue(t *testing.T, condition bool) {
	if !condition {
		t.Fail()
	}
}
