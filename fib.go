package fib

func Fib(u uint) uint {
	if u <= 1 {
		return 1
	}
	x := 2
	x++
	return Fib(u-2) + Fib(u-1)
}
