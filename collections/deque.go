package collections

type Deque struct {
	head []interface{}
	tail []interface{}
}

func NewDeque() *Deque {
	return &Deque{
		head: []interface{}{},
		tail: []interface{}{},
	}
}

func (d *Deque) AppendLast(v interface{}) {
	d.tail = append(d.tail, v)
}

func (d *Deque) AppendFirst(v interface{}) {
	d.head = append(d.head, v)
}

func (d *Deque) PopLast() interface{} {
	if len(d.tail) == 0 {
		v := d.head[0]
		d.head = d.head[1:]
		return v
	} else {
		v := d.tail[len(d.tail)-1]
		d.tail = d.tail[:len(d.tail)-1]
		return v
	}
}

func (d *Deque) PopFirst() interface{} {
	if len(d.head) == 0 {
		v := d.tail[0]
		d.tail = d.tail[1:]
		return v
	} else {
		v := d.head[len(d.head)-1]
		d.head = d.head[:len(d.head)-1]
		return v
	}
}

func (d *Deque) IsEmpty() bool {
	return len(d.head) == 0 && len(d.tail) == 0
}

func (d *Deque) PeekFirst() interface{} {
	if len(d.head) == 0 {
		return d.tail[0]
	} else {
		return d.head[len(d.head)-1]
	}
}

func (d *Deque) PeekLast() interface{} {
	if len(d.tail) == 0 {
		return d.head[0]
	} else {
		return d.tail[len(d.tail)-1]
	}
}
