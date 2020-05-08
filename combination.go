package main

const MOD = 1000000007

var FCT = []int{1}

func ModPow(x, y int) int {
	z := 1
	for y > 0 {
		if y%2 == 0 {
			x = (x * x) % MOD
			y /= 2
		} else {
			z = (z * x) % MOD
			y--
		}
	}
	return z
}

func ModInv(x int) int {
	return ModPow(x, MOD-2)
}

func C(n, r int) int {
	for i := len(FCT); i <= n; i++ {
		FCT = append(FCT, FCT[i-1]*i%MOD)
	}
	return FCT[n] * (ModInv(FCT[n-r]) * ModInv(FCT[r]) % MOD) % MOD
}

func H(n, r int) int {
	return C(n+r-1, r)
}

func NoCacheH(n, r int) int {
	return NoCacheC(n+r-1, r)
}

func NoCacheC(n, r int) int {
	numer := 1
	denom := 1
	for i := 0; i < r; i++ {
		numer = numer * (n - i) % MOD
		denom = denom * (i + 1) % MOD
	}
	return numer * ModInv(denom) % MOD
}
