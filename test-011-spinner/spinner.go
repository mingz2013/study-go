package main

import (
	"time"
	"fmt"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}

	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

/*
同时运行的两个goroutine，


当主goroutine完成后，先强制结束所有其他goroutine，然后程序退出。


*/

func main() {

	go spinner(100 * time.Millisecond)

	const n = 45

	fibN := fib(n) // slow

	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
