package fib

func Fib(u uint) uint {
	if u <= 1 {
		return 1
	}
	for i := 1; i <= 10; i++ {
		x := 1
		x++
	}

	return Fib(u-2) + Fib(u-1)
}
