package main

type UnionFind struct {
	parents []int
	sizes   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{}
	uf.parents = make([]int, n)
	uf.sizes = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parents[i] = i
		uf.sizes[i] = 0
	}
	return &uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parents[x] == x {
		return x
	}

	uf.parents[x] = uf.Find(uf.parents[x])
	return uf.parents[x]
}

func (uf *UnionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}

	if uf.sizes[x] < uf.sizes[y] {
		uf.parents[x] = y
		uf.sizes[y] += uf.sizes[x]
	} else {
		uf.parents[y] = x
		uf.sizes[x] += uf.sizes[y]
	}
}

func (uf *UnionFind) Size(n int) int {
	return uf.sizes[uf.Find(n)]
}
