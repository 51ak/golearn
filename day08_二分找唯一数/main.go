package main

import (
	"fmt"
)

/*
540. Single Element in a Sorted Array
*/

//二分法，因这题目指定要求logn ，只能2分
func singleNonDuplicate(nums []int) int {
	lens := len(nums)
	if lens == 1 {
		return lens
	}

	if nums[0] != nums[1] {
		return nums[0]
	}

	if nums[lens-1] != nums[lens-2] {
		return nums[lens-1]
	}

	lefti, righti := 2, lens-2
	midi := 0

	for lefti < righti {
		midi = lefti + (righti-lefti)>>1
		fmt.Printf("%v(%v-->%v)        %v == %v\n", midi, lefti, righti, nums[midi], nums[midi+1])
		if nums[midi] != nums[midi-1] && nums[midi] != nums[midi+1] {
			fmt.Println("t")
			return nums[midi]
		}

		if midi%2 == 0 {
			if nums[midi] == nums[midi-1] {
				righti = midi - 1
			}

			if nums[midi] == nums[midi+1] {
				lefti = midi + 1
			}
		} else {
			if nums[midi] == nums[midi-1] {
				lefti = midi + 1
			}
			if nums[midi] == nums[midi+1] {
				righti = midi - 1
			}
		}

	}
	return nums[lefti]
}

func main() {

	//var a []int = []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	//var a []int = []int{0, 1, 1, 2, 2, 3, 3}
	var a []int = []int{1, 2, 2, 3, 3}
	//var a []int = []int{3, 3, 7, 7, 10, 11, 11}

	fmt.Printf("%v\n", singleNonDuplicate(a))

}
