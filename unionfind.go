package main

type UnionFind struct {
	parents map[int]int
	sizes   map[int]int
}

func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{
		parents: map[int]int{},
		sizes:   map[int]int{},
	}
	return &uf
}

func (uf *UnionFind) Find(x int) int {
	if _, ok := uf.parents[x]; !ok {
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

	if _, ok := uf.sizes[x]; !ok {
		uf.sizes[x] = 1
	}
	if _, ok := uf.sizes[y]; !ok {
		uf.sizes[y] = 1
	}

	if uf.sizes[x] < uf.sizes[y] {
		uf.parents[x] = y
		uf.sizes[y] += uf.sizes[x]
	} else {
		uf.parents[y] = x
		uf.sizes[x] += uf.sizes[y]
	}
}

func (uf *UnionFind) Size(x int) int {
	x = uf.Find(x)
	if _, ok := uf.sizes[x]; !ok {
		uf.sizes[x] = 1
	}
	return uf.sizes[x]
}
