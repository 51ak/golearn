package day1

import "fmt"

type student struct {
	name                  string
	age, score01, score02 Myint
}

type Myint int

func (st *student) scores() Myint {
	return st.score01 + st.score02
}

func (i Myint) Check() Myint {
	if i < 0 {
		i = -i
	}
	if i > 200 {
		i = 199
	}
	return i
}

func Test04() {
	aa := &student{"张三", 12, 22, 23} //aa:=student{"张三",12,22,23}
	aa.score02 = Myint(-99).Check()
	aa.score01 = Myint(222).Check()
	fmt.Println(aa.scores())
	fmt.Println(aa)
}
