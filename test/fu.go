package main

import "fmt"

func main() {
	//var stu = &Stu{1}
	//stu.do()
	//fmt.Println(stu.age)
	//
	//stu.do()
	arr := []int{1,2,3,4}
	//arr2 := []int{4}
	//swap(arr)
	//arr1 := arr[2:]
	//arr1[0] = 5
	//fmt.Println(arr)
	//arr = append(arr, arr2...)
	//fmt.Println(arr)
	fmt.Println(arr[0:3])
}

type Stu struct {
	age int
}

func (stu *Stu) do() {
	*stu = Stu{2}
	//fmt.Println(stu.age)
}
func swap(arr []int) {
	arr[0],arr[1] = arr[1],arr[0]
}
