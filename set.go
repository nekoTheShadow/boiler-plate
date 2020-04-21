package main

type Set struct {
	set map[interface{}]bool
}

func NewSet() *Set {
	return &Set{
		set: map[interface{}]bool{},
	}
}

func (s *Set) Add(v interface{}) {
	s.set[v] = true
}

func (s *Set) Remove(v interface{}) {
	if _, ok := s.set[v]; ok {
		delete(s.set, v)
	}
}

func (s *Set) Contains(v interface{}) bool {
	_, ok := s.set[v]
	return ok
}

func (s *Set) Values() []interface{} {
	values := []interface{}{}
	for value := range s.set {
		values = append(values, value)
	}
	return values
}
