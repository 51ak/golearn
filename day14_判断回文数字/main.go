package main

import (
	"fmt"
	"strconv"
)

/*
9. Palindrome Number
*/
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	lens := len(s)
	if lens < 2 {
		return true
	}
	if s[0] == '-' {
		return false
	}
	slens := lens / 2

	for i := 0; i < slens; i++ {
		//fmt.Printf("%v,%v,i:%v,lens:%v,slens:%v\n", s[i], s[lens-i-1], i, lens, slens)
		if s[i] != s[lens-i-1] {
			return false
		}
	}
	return true

}
func main() {
	s := isPalindrome(-12521)
	fmt.Println(s)
}
