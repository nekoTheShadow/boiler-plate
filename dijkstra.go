package main

type Dijkstra struct {
	n     int
	nexts [][]node
}

type node struct {
	from int
	to   int
	cost int
}

func NewDijkstra(n int) *Dijkstra {
	nexts := make([][]node, n)
	for i := 0; i < n; i++ {
		nexts[i] = []node{}
	}
	return &Dijkstra{
		n:     n,
		nexts: nexts,
	}
}

func (d *Dijkstra) Add(from, to, cost int) {
	d.nexts[from] = append(d.nexts[from], node{from: from, to: to, cost: cost})
}

func (d *Dijkstra) Run(start int) []int {
	score := make([]int, d.n)
	for i := 0; i < d.n; i++ {
		score[i] = INFINITY
	}
	score[start] = 0

	q := NewHeapq()
	q.Push(0, start)
	for !q.isEmpty() {
		cur := q.Pop().(int)
		for _, next := range d.nexts[cur] {
			if score[cur]+next.cost < score[next.to] {
				score[next.to] = score[cur] + next.cost
				q.Push(score[next.to], next.to)
			}
		}
	}

	return score
}
