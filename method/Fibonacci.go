package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	x,y := 0,1

	result := func() int {
		x,y = y,x+y
		return x
	}
	return result
}

func main() {
	// 函数的返回值相当于内层函数的地址
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
