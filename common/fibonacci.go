package common

func Fib(n int) int {
	f, s := 1, 1
	for i := 1; i < n; i++ {
		f, s = s, f+s
	}
	return f
}

func GenFib(a, b int) func(n int) int {
	return func(n int) int {
		f, s := a, b
		for i := 1; i < n; i++ {
			f, s = s, f+s
		}
		return f
	}
}
