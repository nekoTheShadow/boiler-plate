package codeforces

var MOD = int64(1000000007)

var FCT = []int64{1}

func ModPow(x, y int64) int64 {
	z := int64(1)
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

func ModInv(x int64) int64 {
	return ModPow(x, MOD-2)
}

func C(n, r int64) int64 {
	for i := int64(len(FCT)); i <= n; i++ {
		FCT = append(FCT, FCT[i-1]*i%MOD)
	}
	return FCT[n] * (ModInv(FCT[n-r]) * ModInv(FCT[r]) % MOD) % MOD
}

func H(n, r int64) int64 {
	return C(n+r-1, r)
}

func NoCacheC(n, r int64) int64 {
	numer := int64(1)
	denom := int64(1)
	for i := int64(0); i < r; i++ {
		numer = numer * (n - i) % MOD
		denom = denom * (i + 1) % MOD
	}
	return numer * ModInv(denom) % MOD
}

func NoCacheH(n, r int64) int64 {
	return NoCacheC(n+r-1, r)
}
