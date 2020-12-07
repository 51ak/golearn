package main

import "fmt"

/*Given a string s, return the longest palindromic substring in s.



Example 1:

Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
Example 3:

Input: s = "a"
Output: "a"
Example 4:

Input: s = "ac"
Output: "a"


Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters (lower-case and/or upper-case),
*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	slice := make([]string, 0, 100)
	for _, c := range s {
		slice = append(slice, ",")
		slice = append(slice, string(c))
	}
	slice = append(slice, ",")
	// abc --> ,a,b,c,   abcd -->,a,b,c,d,
	// fmt.Printf("%s", slice)
	maxr, maxr_index, lens := 0, 0, len(slice)
	for i := 1; i < lens; i++ {
		//main loop
		r, lefti, righti := 0, i-1, i+1
		for {
			if lefti < 0 || righti >= lens {
				break
			}

			if slice[lefti] == slice[righti] {

				//fmt.Printf("%s|%s|%d|%d\n", slice[lefti], slice[righti], lefti, righti)
				r++
				lefti--
				righti++
			} else {
				break
			}
		}

		if r > maxr {
			maxr = r
			maxr_index = i
		}
	}
	//slice_new :=slice[maxr_index-maxr:maxr+maxr_index]
	s_new := ""
	// fmt.Println(maxr)
	// fmt.Println(maxr_index)
	for j := maxr_index - maxr; j < maxr+maxr_index; j++ {
		if j < lens {
			if slice[j] != "," {
				s_new += slice[j]
			}
		}
	}
	// 	//for _,s := range slice[maxr_index-maxr,j<:maxr+maxr_index-1]{
	// 	// if s != ","{
	// 	//     s_new+=s
	// 	// }
	// }
	return s_new
}
func main() {
	s := longestPalindrome("ab")
	fmt.Println(s)
}
