package main

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
