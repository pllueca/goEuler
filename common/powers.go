package common

func Pow(a,b int64) int64 {
	var r int64
	r = 0
	for b != 0 {
		if b&1 == 1 {
			r *= a
			b--
		}	
		// here b is even
		a *= a
		b /= 2
	}
	return r
}
