package collections

type UnionFind struct {
	parents []int
}

func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{}
	uf.parents = []int{}
	for i := 0; i < n; i++ {
		uf.parents = append(uf.parents, i)
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
	if x != y {
		uf.parents[y] = x
	}
}
