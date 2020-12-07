package day1

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for x := range a {
		b := make([]uint8, dx)
		for y := range b {
			b[y] = uint8(x*y - 1)
		}
		a[x] = b
	}
	//fmt.Println(a)
	return a
}

func Test03() {
	pic.Show(Pic)
}
