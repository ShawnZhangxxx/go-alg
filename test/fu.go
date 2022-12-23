package main

import "fmt"

func main() {
	var stu = &Stu{1}
	stu.do()
	fmt.Println(stu.age)

	stu.do()
}

type Stu struct {
	age int
}

func (stu *Stu) do() {
	*stu = Stu{2}
	//fmt.Println(stu.age)
}
