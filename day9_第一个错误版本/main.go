package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var errori int = 1

/*
278. First Bad Version
*/
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {

	lefti, righti := 1, n
	midi := 1
	for lefti < righti {
		// if lefti >= righti+1 {
		// 	if isBadVersion(lefti + 1) {
		// 		fmt.Println("left+1")
		// 		return lefti + 1
		// 	}
		// 	fmt.Println("left")
		// 	return lefti
		// }
		midi = lefti + (righti-lefti)/2
		//fmt.Printf("lefti:%v, righti:%v, midi:%v, errori:%v\n", lefti, righti, midi, errori)
		if isBadVersion(midi) {
			righti = midi
		} else {
			lefti = midi + 1
		}

	}

	return lefti

}

func isBadVersion(version int) bool {
	return version >= errori
}

func main() {
	rand.Seed(time.Now().Unix())
	var a []int = []int{1, 2, 3, 4, 9, 11, 22, 1001, 9999, 9999999, 777777, math.MaxInt32 - 1, math.MaxInt32 - 2}
	for _, n := range a {
		errori = rand.Intn(n)
		if errori == 0 {
			errori = 1
		}
		fmt.Printf("n:%v,errori:%v,result:%v\n", n, errori, firstBadVersion(n))
	}
	errori = 2
	fmt.Printf("n:%v,errori:%v,result:%v\n", 2, errori, firstBadVersion(2))
}
