package main

import (
	"fmt"
	"sort"
)

/*
455. Assign Cookies
*/

func findContentChildren(g []int, s []int) int {
	lensg := len(g)
	lenss := len(s)
	if lensg == 0 || lenss == 0 {
		return 0
	}
	//对孩子胃口数组从小到大排序
	sort.Ints(g)
	//对饼干尺寸数组从小到大排序
	sort.Ints(s)
	//fmt.Println(s)
	j := 0
	okcount := 0
	for i := 0; i < lensg; i++ {
		for ; j < lenss; j++ {
			if g[i] <= s[j] {
				//fmt.Printf("g[%v]=%v,s[%v]=%v\n", i, g[i], j, s[j])
				okcount++
				j++
				break
			}
		}

		if j == lenss {
			break
		}

	}
	return okcount
}

func main() {
	// childs := []int{4, 7, 2, 1, 8, 5, 2}
	// cookies := []int{1, 2, 3, 4, 5}

	// childs := []int{10, 9, 8, 7, 10, 9, 8, 7}
	// cookies := []int{10, 9, 8, 7}

	childs := []int{7, 7, 8, 8, 9, 9, 10, 10}
	cookies := []int{7, 8, 9, 10}

	result := findContentChildren(childs, cookies)
	fmt.Println(result)
}
