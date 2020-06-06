package main

type Dijkstra2 struct {
	h     int
	w     int
	nexts map[from][]to
}

type from struct {
	x int
	y int
}

type to struct {
	x    int
	y    int
	cost int
}

func NewDijkstra2(h, w int) *Dijkstra2 {
	return &Dijkstra2{
		h:     h,
		w:     w,
		nexts: make(map[from][]to),
	}
}

func (d *Dijkstra2) Add(x1, y1, x2, y2, cost int) {
	k := from{x: x1, y: y1}
	if _, ok := d.nexts[k]; !ok {
		d.nexts[k] = []to{}
	}
	d.nexts[k] = append(d.nexts[k], to{
		x:    x2,
		y:    y2,
		cost: cost,
	})
}

func (d *Dijkstra2) Run(x, y int) [][]int {
	score := make([][]int, d.h)
	for i := 0; i < d.h; i++ {
		score[i] = make([]int, d.w)
		for j := 0; j < d.w; j++ {
			score[i][j] = INFINITY
		}
	}
	score[x][y] = 0

	q := NewHeapq()
	q.Push(0, from{x: x, y: y})
	for !q.isEmpty() {
		cur := q.Pop().(from)
		if _, ok := d.nexts[cur]; !ok {
			continue
		}
		for _, next := range d.nexts[cur] {
			if score[cur.x][cur.y]+next.cost < score[next.x][next.y] {
				score[next.x][next.y] = score[cur.x][cur.y] + next.cost
				q.Push(score[next.x][next.y], from{x: next.x, y: next.y})
			}
		}
	}

	return score
}
