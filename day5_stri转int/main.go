package main

import (
	"fmt"
	"math"
	"strings"
)

/*
8. String to Integer (atoi)
*/
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	iresult := int64(0)
	isbegin := false
	signnum := int64(1)
	for _, v := range s {

		if v >= '0' && v <= '9' {
			isbegin = true
			//fmt.Println(int(v) - 48)
			iresult = iresult*int64(10) + int64(v) - int64(48)
			if iresult > math.MaxInt32 {
				break
			}
		} else {
			if isbegin {
				break
			}
			if v == '+' {
				isbegin = true
			} else if v == '-' {
				isbegin = true
				signnum = -1
			} else {

				break

			}
		}

	}
	//fmt.Printf("%v,%v:%v,%v", signnum, iresult, math.MinInt32, math.MaxInt32)

	iresult = iresult * signnum
	//fmt.Printf("%v,%v:%v,%v", signnum, iresult, math.MinInt32, math.MaxInt32)
	if iresult < math.MinInt32 {
		return math.MinInt32
	}

	return int(iresult)

}
func main() {
	// s := myAtoi("words and 987")
	// fmt.Println(s)
	s := myAtoi("922337203685477580822")
	fmt.Println(s)
}
