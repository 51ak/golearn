package main

import (
	"fmt"
	"strings"
)

/*
402. Remove K Digits
*/
func removeKdigits(num string, k int) string {
	lensNum := len(num)
	returnLens := lensNum - k
	if k == 0 || lensNum == 0 {
		return num
	}
	if returnLens <= 0 {
		return "0"
	}

	stackStr := make([]byte, lensNum)
	stackIndex := 0
	for i := 0; i < lensNum; i++ {
		for k > 0 && stackIndex > 0 && stackStr[stackIndex-1] > num[i] {
			k--
			stackIndex--
		}
		stackStr[stackIndex] = num[i]
		stackIndex++
	}

	resultstr := strings.TrimLeft(string(stackStr[0:returnLens]), "0")
	if len(resultstr) == 0 {
		return "0"
	}
	return resultstr

}

//这是网上找的一段代码，比我写的工整，简洁，参照
func removeKdigits2(num string, k int) string {
	digits := len(num) - k
	stack := make([]byte, len(num))
	top := 0

	for i := range num {
		for top > 0 && stack[top-1] > num[i] && k > 0 {
			top--
			k--
		}

		stack[top] = num[i]
		top++
	}

	i := 0
	for i < digits && stack[i] == '0' {
		i++
	}

	if i == digits {
		return "0"
	}

	return string(stack[i:digits])
}

func main() {
	nums := []string{"14325709", "10200", "310837726325978"}
	lens := len(nums)
	for i := 0; i < lens; i++ {
		for j := 0; j < len(nums[i]); j++ {
			reult := removeKdigits(nums[i], j)
			fmt.Printf("   - %v擦除%v次，最小结果%v\n", nums[i], j, reult)
		}
	}

}
