package day1

import (
	"fmt"
)

// fibonacci 函数会返回一个返回 int 的函数。第3个数开始等于前两个数之和
func fibonacci() func() int {
	x1 := 1
	x2 := 1
	return func() int {
		x3 := x1 + x2
		x1 = x2
		x2 = x3
		return x3
	}
}

func Test02() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
