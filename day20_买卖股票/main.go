package main

import (
	"fmt"
)

/*
121. Best Time to Buy and Sell Stock
*/

func maxProfit(prices []int) int {
	lens := len(prices)
	if lens <= 1 {
		return 0
	}
	min, resulti := prices[0], 0
	for i := 1; i < lens; i++ {
		if prices[i]-min > resulti {
			resulti = prices[i] - min
		} else if prices[i] < min {
			min = prices[i]
		}

	}
	return resulti
}

func main() {
	// childs := []int{4, 7, 2, 1, 8, 5, 2}
	// cookies := []int{1, 2, 3, 4, 5}

	// childs := []int{10, 9, 8, 7, 10, 9, 8, 7}
	// cookies := []int{10, 9, 8, 7}

	childs := []int{7, 7, 8, 8, 9, 9, 10, 10}
	//cookies := []int{7, 8, 9, 10}

	result := maxProfit(childs)
	fmt.Println(result)
}
