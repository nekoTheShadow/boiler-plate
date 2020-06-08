package main

type Ruiseki2Builder struct {
	a [][]int
	h int
	w int
}

type Ruiseki2Result struct {
	s [][]int
}

func NewRuiseki2Builder(h, w int) *Ruiseki2Builder {
	a := createMatrix(h, w)
	return &Ruiseki2Builder{a: a, h: h, w: w}
}

func (r *Ruiseki2Builder) Set(x, y, v int) {
	r.a[x][y] = v
}

func (r *Ruiseki2Builder) Build() *Ruiseki2Result {
	s := createMatrix(r.h+1, r.w+1)
	for i := 0; i < r.h; i++ {
		for j := 0; j < r.w; j++ {
			s[i+1][j+1] = s[i][j+1] + s[i+1][j] - s[i][j] + r.a[i][j]
		}
	}
	return &Ruiseki2Result{s: s}
}

// [x1, x1) [y1, y2)の累積和を求める
func (r *Ruiseki2Result) Get(x1, y1, x2, y2 int) int {
	return r.s[x2][y2] - r.s[x1][y2] - r.s[x2][y1] + r.s[x1][y1]
}

func createMatrix(h, w int) [][]int {
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
	}
	return a
}
