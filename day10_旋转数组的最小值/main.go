package main

import (
	"fmt"
	"math/rand"
	"time"
)

var errori int = 1

/*
153. Find Minimum in Rotated Sorted Array
*/
func findMin(nums []int) int {
	lens := len(nums)

	// if lens == 0 {
	// 	return -1
	// }
	if lens < 5 {
		tempi := nums[0]
		for _, i := range nums {
			if i < tempi {
				tempi = i
			}
		}
		return tempi
	}
	lefti, righti := 0, lens-1
	i := 0
	for lefti < righti {
		midi := lefti + (righti-lefti)/2
		if midi == lefti || midi == righti {
			return nums[midi]
		}
		//fmt.Println(midi)
		if nums[midi] < nums[midi-1] {
			return nums[midi]
		}
		if nums[midi] > nums[midi+1] {
			return nums[midi+1]
		}
		if nums[midi] < nums[righti] {
			righti = midi
		} else {
			lefti = midi
		}
		i++
		if i > 1000 {
			break
		}
	}

	return -1
}

func changenum(mini int, maxi int, stepi int) []int {
	nums := make([]int, 0)

	i := mini
	for i < maxi {

		nums = append(nums, i)
		i = i + stepi
		// fmt.Println(nums)
	}
	lens := len(nums)

	changei := rand.Intn(lens)
	if changei < 0 {
		changei = 0
	}
	//num2 := make([]int, 5)
	copyData := make([]int, 0)
	//copy(copyData, nums[changei:lens])
	copyData = append(copyData, nums[changei:lens]...)
	copyData = append(copyData, nums[0:changei]...)
	//newnums := append(nums[changei-1:], nums[0:changei-1])
	return copyData
}

func main() {
	rand.Seed(time.Now().Unix())
	//var a []int = []int{1, 2, 3, 4, 9, 11, 22, 1001, 9999, 9999999}

	var a []int = []int{1, 2, 3, 4, 9, 11, 22, 1001, 9999, 9999999, 777777}
	for _, n := range a {

		mini := rand.Intn(n)
		diffi := n - mini

		if diffi > 50 {
			diffi = rand.Intn(50)
		}
		stepi := rand.Intn((n - mini))
		if stepi < 2 {
			stepi = 1
		}
		if n < 20 {
			stepi = 1
		}
		if stepi > 1363475 {
			stepi = 1363475
		}
		nums := changenum(mini, n, stepi)
		resutl := findMin(nums)
		fmt.Printf("min:%v,max:%v,step:%v,resutl:%v \n", mini, n, stepi, resutl)
		//fmt.Println(nums)

		//fmt.Printf("n:%v,errori:%v,result:%v\n", n, changei, firstBadVersion(n))
	}

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(len(nums))
	resutl := findMin(nums)
	fmt.Println(resutl)
}
