package day1

import (
	"code.google.com/p/go-tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var r = make(map[string]int)
	for _, value := range strings.Fields(s) {
		svalue := string(value)
		r[svalue] = r[svalue] + 1
		//fmt.Println(svalue)
	}
	return r
}

func Test01() {
	fmt.Println("WordCount")
	wc.Test(WordCount)
}
