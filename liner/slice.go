package main

import "C"
import (
	"fmt"
	"unsafe"
)
type Slice struct {
	Data unsafe.Pointer //万能指针类型，对应c语言中void*
	len int
	cap int
}

const TAG  = 8

func main()  {
	var s []int
	fmt.Println(unsafe.Sizeof(s))//
}

func (s *Slice)Create(l int,c int,Data ...int)  {
	if len(Data) == 0{
		return
	}

	if l < 0 || c < 0 || l > c || len(Data) > l {
		return
	}

	s.Data = C.malloc(C.ulonglong(c) * 8)
	s.len = l
	s.cap = c
	//转换成可计算的指针类型
	p := uintptr(s.Data)

	for _,v := range Data {
		*(* int)(unsafe.Pointer(p)) = v
		p += TAG
	}

}