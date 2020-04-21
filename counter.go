package main

import (
	"fmt"
	"strings"
)

type Counter map[interface{}]int

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment(key interface{}) {
	d := map[interface{}]int(*c)
	d[key] = c.Get(key) + 1
}

func (c *Counter) Get(key interface{}) int {
	d := map[interface{}]int(*c)
	if v, ok := d[key]; ok {
		return v
	} else {
		return 0
	}
}

func (c *Counter) Keys() []interface{} {
	d := map[interface{}]int(*c)
	keys := []interface{}{}
	for key := range d {
		keys = append(keys, key)
	}
	return keys
}

func (c *Counter) String() string {
	strs := []string{}
	for _, key := range c.Keys() {
		strs = append(strs, fmt.Sprintf("%v => %d", key, c.Get(key)))
	}
	return "[" + strings.Join(strs, ", ") + "]"
}
