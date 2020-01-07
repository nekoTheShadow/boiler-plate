package main

type Deque struct {
	deque []interface{}
}

func NewDeque() *Deque {
	return &Deque{deque: []interface{}{}}
}

func (d *Deque) AppendFirst(v interface{}) {
	e := make([]interface{}, len(d.deque)+1)
	copy(e[1:], d.deque)
	d.deque = e
}

func (d *Deque) AppendLast(v interface{}) {
	d.deque = append(d.deque, v)
}

func (d *Deque) PopLast() interface{} {
	n := len(d.deque)
	v := d.deque[n-1]
	d.deque = d.deque[0 : n-1]
	return v
}

func (d *Deque) PopFirst() interface{} {
	v := d.deque[0]
	d.deque = d.deque[1:]
	return v
}

func (d *Deque) IsEmpty() bool {
	return len(d.deque) == 0
}
