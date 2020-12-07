package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	var MIN int = 0x80000000
	var MAX int = 0x7FFFFFFF

	y := x
	if x < 0 {
		y = -x
	}
	str_x := strconv.Itoa(y)

	// length := len(str_x)

	// result := ""
	// for i := length - 1; i >= 0; i-- {
	// 	result += string(str_x[i])
	// }
	var result []byte
	tmp := []byte(str_x)
	length := len(str_x)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}

	// ri, _ := strconv.Atoi(string(result))
	ri, _ := strconv.Atoi(string(result))

	if x < 0 {
		ri = -ri
	}
	if ri < -MIN || ri > MAX {

		ri = 0
	}
	return ri
}

func main() {
	s := reverse(-123410)
	fmt.Println(s)
}
