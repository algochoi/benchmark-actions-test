package fib

import "time"

func Fib(u uint) uint {
	if u <= 1 {
		return 1
	}
	for i := 1; i <= 1; i++ {
		x := 1
		x++
		time.Sleep(1)
	}

	return Fib(u-2) + Fib(u-1)
}
