package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	a := []int{1, 2, 3, 4}
	ReverseSlice(a)
	assert(t, a, []int{4, 3, 2, 1})

	b := []string{"あ", "い", "う"}
	ReverseSlice(b)
	assert(t, b, []string{"う", "い", "あ"})

	c := "かきくけこ"
	ReverseSlice(c)
	assert(t, c, "かきくけこ") // not Reversed
}

func TestReverseString(t *testing.T) {
	assert(t, ReverseString("かきくけこ"), "こけくきか")
	assert(t, ReverseString("ABCD"), "DCBA")
	assert(t, ReverseString("A1𩸽"), "𩸽1A")
}

func TestReverseToStringSlice(t *testing.T) {
	assert(t, ToStringSlice([]int{1, 2, 3}), []string{"1", "2", "3"})
	assert(t, ToStringSlice([]int{1, 2, 3, 4}), []string{"1", "2", "3", "4"})
}

func assert(t *testing.T, result, expected interface{}) {
	fmt.Println(result)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result: %v, expected: %v", result, expected)
	}
}
