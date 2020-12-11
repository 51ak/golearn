package main

import (
	"fmt"
	"math/rand"
	"time"
)

var errori int = 1

/*
70. Climbing Stairs
*/
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	resulti, prei, prerprei := 2, 2, 1
	for i := 3; i < n+1; i++ {
		resulti = prei + prerprei
		prerprei = prei
		prei = resulti
	}
	return resulti
}

func main() {
	rand.Seed(time.Now().Unix())
	//var a []int = []int{1, 2, 3, 4, 9, 11, 22, 1001, 9999, 9999999}

	var a []int = []int{1, 2, 3, 4, 9, 11, 22, 30, 1001, 9999, 9999999, 777777}
	for _, n := range a {
		result := climbStairs(n)
		fmt.Printf("%v-->%v\n", n, result)

	}

}
