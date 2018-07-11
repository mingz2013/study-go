package main

import (
	"fmt"
	"os"
	"strings"
)

func doArg1() {
	var s string
	for i := 1; i < len(os.Args); i++ {
		//fmt.Println(os.Args[i])
		s += os.Args[i] + " "
	}
	fmt.Println(s)
}

func doArg2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func doArg3() {
	// 高效的方式, 上面两种方式性能有问题
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	doArg1()

	doArg2()

	doArg3()

}
