package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
有两个一模一样的小盒子，1号盒子有30颗糖果和10个乒乓球。2号盒子里有20颗糖果和20个乒乓球。随机选择一个盒子，从里面随机拿出一样东西出来，发现是颗糖果。问这颗糖果来自1号盒子的概率是多少？
暴力计算法
*/
func randbox() (bool, bool) {
	var si, sj int
	rand.Seed(time.Now().UnixNano())
	si = rand.Intn(99)
	sj = rand.Intn(399)
	if si < 50 {
		return true, sj < 300
	} else {
		return false, sj < 200
	}
}

func main() {
	var isone, istang bool
	okcount, oneokcount := 0, 0
	maxcount := 1000000000
	for i := 1; i <= maxcount; i++ {
		isone, istang = randbox()
		if istang {
			okcount++
			if isone {
				oneokcount++
			}
		}
		if i%100000 == 0 {
			fmt.Printf("计算%v次结果：有效次数（是糖）%v次,在第一个盒子里%v次，概率：0.%v\n", i, okcount, oneokcount, oneokcount*1000000/okcount)
		}
	}
	// fmt.Println("....")
	// fmt.Printf("第%v次结果：%v\n", 1000, nums)
}
