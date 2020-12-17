package main

import (
	"fmt"
)

/*
78. Subsets
*/
var sublist [][]int

func subsets(nums []int) [][]int {
	sublist = [][]int{}
	lens := len(nums)
	resulti := []int{}
	for resultlens := 0; resultlens <= lens; resultlens++ {
		//i--> the length of ther resulti []int
		subsetsLoop(resultlens, nums, resulti, 0)
		//fmt.Printf("%v:--->%v \n", resultlens, sublist)
	}
	return sublist
}

func subsetsLoop(resultlens int, nums []int, reuslti []int, ci int) {
	if len(reuslti) == resultlens {
		tmp := make([]int, resultlens)
		copy(tmp, reuslti)
		//fmt.Println(tmp)
		sublist = append(sublist, tmp)
		return
	}
	for j := ci; j < len(nums); j++ {
		reuslti = append(reuslti, nums[j])
		subsetsLoop(resultlens, nums, reuslti, j+1)
		reuslti = reuslti[:len(reuslti)-1]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	resutl := subsets(nums)
	// resa := fmt.Sprintf("%v", resutl)
	// resa = strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(resa, "1", "飞", -1), "2", "羽", -1), "3", "云", -1), "4", "超", -1), "5", "忠", -1), "[]", "", -1)
	// fmt.Printf("%v种排法:%v\n", len(resutl)-1, resa)
	fmt.Println(resutl)
}
