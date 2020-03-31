package main

// Binary Indexed Tree (1 based indexed)
type BIT struct {
	size int
	tree []int
}

func NewBIT(size int) *BIT {
	bit := BIT{
		size: size,
		tree: make([]int, size+1),
	}
	return &bit
}

func (bit *BIT) Sum(i int) int {
	sum := 0
	for i > 0 {
		sum += bit.tree[i]
		i -= i & -i
	}
	return sum
}

func (bit *BIT) Add(i, x int) {
	for i <= bit.size {
		bit.tree[i] += x
		i += i & -i
	}
}
