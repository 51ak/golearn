package main

import (
	"fmt"
	"math/rand"
	"time"
)

var errori int = 1

/*
198. House Robber
*/
func rob(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	} else if lens == 1 {
		return nums[0]
	}
	//  else if lens == 2 {
	// 	return MaxInt(nums[0], nums[1])
	// }
	result := make([]int, 0)
	last1 := MaxInt(nums[0], nums[1])
	tmp, last2 := 0, nums[0]
	for i := 2; i < lens; i++ {
		tmp = last1
		last1 = MaxInt(nums[i]+last2, tmp)

		if nums[i]+last2 > tmp {
			result = append(result, nums[i])
		}
		last2 = tmp
	}
	fmt.Println(result)
	return last1
}

/* MaxInt
golang没有math.maxint自己民实现一个
*/
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maketest(mini int, maxi int, stepi int, counts int) []int {
	nums := make([]int, 0)
	rand.Seed(time.Now().Unix())
	tmpi := 0
	for i := 0; i < counts; i++ {
		if stepi == 1 {
			tmpi = i
		} else if stepi == 0 {
			tmpi = mini + rand.Intn(maxi-mini)
		} else {
			tmpi = mini + stepi*i
		}
		nums = append(nums, tmpi)
	}
	return nums
}

func main() {
	// rand.Seed(time.Now().Unix())
	// //var a []int = []int{1, 2, 3, 4, 9, 11, 22, 1001, 9999, 9999999}

	// var a []int = []int{1, 2, 3, 4, 9, 11, 22}
	// for _, n := range a {

	// 	nums := maketest(0, 400, 0, n)
	// 	resutl := rob(nums)
	// 	fmt.Printf("nums:%v,resutl:%v \n", nums, resutl)
	// 	//fmt.Println(nums)

	// 	//fmt.Printf("n:%v,errori:%v,result:%v\n", n, changei, firstBadVersion(n))
	// }

	nums := []int{4, 3, 5, 4, 5, 23, 7, 8, 9, 7, 7, 8, 12, 15, 17, 13, 4, 9, 12, 17, 13, 5, 0, 7, 14, 7, 9, 8, 9, 30}
	fmt.Println(len(nums))
	resutl := rob(nums)
	fmt.Println(resutl)
}
