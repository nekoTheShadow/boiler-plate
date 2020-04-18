package main

import (
	"fmt"
	"reflect"
)

func isPrime(x int) bool {
	if x < 2 {
		return false
	}

	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func Gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}

	for y > 0 {
		x, y = y, x%y
	}

	return x
}

func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}

func CeilDiv(x, y int) int {
	if x%y == 0 {
		return x / y
	} else {
		return x/y + 1
	}
}

func LCS(s, t string) string {
	n := len(s)
	m := len(t)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	x := n
	y := m
	u := make([]byte, dp[n][m])
	i := dp[n][m] - 1
	for x > 0 && y > 0 {
		if dp[x][y] == dp[x-1][y] {
			x--
		} else if dp[x][y] == dp[x][y-1] {
			y--
		} else {
			x--
			y--
			u[i] = s[x]
			i--
		}
	}

	return string(u)
}

// [0..n) から r個選んだ組合せを生成する。
// r < 0 の場合はすべての組合せを生成する。
func Combinations(n int, r int) [][]int {
	a := [][]int{}
	for i := uint(0); i < (1 << uint(n)); i++ {
		b := []int{}
		for j := uint(0); j < uint(n); j++ {
			if i&(uint(1)<<j) != 0 {
				b = append(b, int(j))
			}
		}

		if r < 0 || len(b) == r {
			a = append(a, b)
		}
	}
	return a
}

// [0..n)の順列を生成する。
func Permutations(n int) [][]int {
	type e struct {
		bit uint
		a   []int
	}

	q := []*e{}
	for i := 0; i < n; i++ {
		q = append(q, &e{bit: 1 << uint(i), a: []int{i}})
	}

	p := [][]int{}
	for x := 0; x < len(q); x++ {
		cur := q[x]
		if cur.bit == (1<<uint(n) - 1) {
			p = append(p, cur.a)
			continue
		}

		for i := 0; i < n; i++ {
			if cur.bit&(1<<uint(i)) != 0 {
				continue
			}

			m := len(cur.a) + 1
			b := make([]int, m)
			copy(b, cur.a)
			b[m-1] = i
			q = append(q, &e{bit: cur.bit | (1 << uint(i)), a: b})
		}
	}
	return p
}

func ReverseSlice(v interface{}) {
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		slice := reflect.ValueOf(v)
		for i1, n := 0, slice.Len(); i1 < n/2; i1++ {
			i2 := n - i1 - 1
			v1 := reflect.ValueOf(slice.Index(i1).Interface())
			slice.Index(i1).Set(slice.Index(i2))
			slice.Index(i2).Set(v1)
		}
	}
}

func ReverseString(s string) string {
	t := []rune(s)
	ReverseSlice(t)
	return string(t)
}

func ToStringSlice(slice interface{}) []string {
	// 本来はerrを返却すべき。 利用する側(=自分)は引数がsliceであることを知っている。
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v is not slice", slice))
	}

	src := reflect.ValueOf(slice)
	n := src.Len()
	dst := make([]string, n)
	for i := 0; i < n; i++ {
		dst[i] = fmt.Sprintf("%v", src.Index(i))
	}
	return dst
}
