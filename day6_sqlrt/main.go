package main

import (
	"fmt"
	"math"
)

/*
69. Sqrt(x)
*/
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	lefti := 1
	rigthi := x

	midi := 0
	resulti := 0
	for {
		if lefti >= rigthi {
			break
		}

		//do not  (lefti+righti)/2   int32+int32> int32max
		midi = lefti + (rigthi-lefti)>>1
		resulti = midi * midi

		//fmt.Printf("lefti:%v,righti:%v,midi:%v,resulti:%v\n", lefti, rigthi, midi, resulti)
		if midi == lefti {
			break
		}

		if resulti == x {
			return midi
		} else if resulti > x {
			rigthi = midi
		} else {
			lefti = midi
		}

	}
	return midi

}

func mySqrt2(x int) int {
	return int(math.Sqrt(float64(x)))
}

func main() {
	a := [...]int{2, 4, 8, 52, 12, 43, 557, 12341, 124314, 5555, 9999, 2222}
	//a := [...]int{4}
	for _, x := range a {
		fmt.Printf("%v:%v-->%v\n", x, mySqrt(x), mySqrt2(x))
	}

}
