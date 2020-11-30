package main

import (
	"fmt"
)

/*
744. Find Smallest Letter Greater Than Target
*/

//二分法
func nextGreatestLetter(letters []byte, target byte) byte {
	lefti, righti := 0, len(letters)-1

	if target >= letters[righti] || target < letters[lefti] {
		return letters[0]
	}

	midi := 0
	for lefti < righti {
		midi = lefti + (righti-lefti)>>1

		fmt.Printf("lefti:%v,righti:%v,midi:%v,resulti:%v\n", lefti, righti, midi, letters[midi])
		if righti == midi {
			break
		}
		if lefti == midi {
			if lefti < righti {
				midi++
			}
			break
		}
		if target > letters[midi] {
			lefti = midi + 1
		} else if target < letters[midi] {
			righti = midi
		} else {
			midi = midi + 1
			break
		}
	}
	return letters[midi]
}

//暴力法
func nextGreatestLetterOld(letters []byte, target byte) byte {
	for _, key := range letters {
		if key > target {
			return key
		}
	}
	return letters[0]
}

func main() {
	var b []byte = make([]byte, 3)
	b[0] = 'c'
	b[1] = 'f'
	b[1] = 'g'
	a := [...]byte{'a', 'c', 'd', 'f', 'g', 'h'}
	for _, key := range a {
		fmt.Printf("%v-->%v\n", nextGreatestLetter(b, key), nextGreatestLetterOld(b, key))
	}

}
