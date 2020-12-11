package main

import (
	"fmt"
)

/*
9. 桌子上有5个盒子，从左到右分别装有8,4,2,1,5个相同的小球。每次操作都从其余4个盒子中各拿出1个小球，放入小球数最少的盒子中；如果有2个盒子中小球数相同，就认为其中左侧盒子的小球数较少；如果某个盒子中只有1个小球，则不从其中拿出小球。经过10000次操作后，第1个盒子中还有多少个小球？
*/
func changebox(nums []int, lens int, t_minindex int) int {
	t_mini := 1000000
	last_miniindex := t_minindex
	for i := 0; i < lens; i++ {
		if i == last_miniindex {
			nums[i] = nums[i] + lens - 1
		} else if nums[i] == 0 {
			nums[last_miniindex]--
			if t_mini > 0 {
				t_minindex = i
				t_mini = 1
			}
		} else {
			nums[i]--
			if nums[i] < t_mini {
				t_minindex = i
				t_mini = nums[i]
			}
		}
	}
	return t_minindex
}

func main() {
	nums := []int{8, 4, 2, 1, 5}
	//nums := []int{8, 8, 7, 4, 2, 1, 5, 8}
	lens := len(nums)
	lastminiindex := 3
	//fmt.Println(nums)
	// fmt.Println("=", nums[0]+nums[1]+nums[2]+nums[3]+nums[4])
	for i := 0; i < 20; i++ {
		lastminiindex = changebox(nums, lens, lastminiindex)
		if i < 20 {
			fmt.Printf("第%v次结果：%v\n", i+1, nums)
		}
	}
	fmt.Println("....")
	fmt.Printf("第%v次结果：%v\n", 1000, nums)
}
