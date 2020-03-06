package ints

import "testing"

func TestPow(t *testing.T) {
	tests := []struct {
		x int
		y int
		z int
	}{
		{x: 2, y: 0, z: 1},
		{x: 2, y: 2, z: 4},
		{x: 3, y: 3, z: 27},
	}

	for _, test := range tests {
		if Pow(test.x, test.y) != test.z {
			t.FailNow()
		}
	}
}

func TestCeilDiv(t *testing.T) {
	tests := []struct {
		x int
		y int
		z int
	}{
		{x: 10, y: 2, z: 5},
		{x: 10, y: 3, z: 4},
	}
	for _, test := range tests {
		if CeilDiv(test.x, test.y) != test.z {
			t.Errorf("CeilDiv(%d, %d) - expected=%d, but actual=%d", test.x, test.y, test.z, CeilDiv(test.x, test.y))
		}
	}
}

func TestMax(t *testing.T) {
	if Max(1) != 1 {
		t.Errorf("Max(1): expected 1, actual %d", Max(1))
	}

	if Max(2, 3, 1) != 3 {
		t.Errorf("Max(2,3,1): expected 3, actual %d", Max(2, 3, 1))
	}
}

func TestMin(t *testing.T) {
	if Min(1) != 1 {
		t.Errorf("Min(1): expected 1, actual %d", Min(1))
	}

	if Min(2, 3, 1) != 1 {
		t.Errorf("Min(2,3,1): expected 1, actual %d", Min(2, 3, 1))
	}
}
