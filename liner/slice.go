package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Slice struct {
	Data unsafe.Pointer //万能指针类型，对应c语言中void*
	len  int
	cap  int
}

const TAG = 8

//func main() {
//	//fmt.Println(111)
//	s := new(Slice)
//	s.Create(3,3,1,2,3)
//	s.Append(4)
//	s.Print()
//	s.Search(2)
//	s.Delete(1)
//	fmt.Println("delete--")
//	s.Print()
//}

func (s *Slice) Create(l int, c int, Data ...int) {
	if len(Data) == 0 {
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

	for _, v := range Data {
		*(*int)(unsafe.Pointer(p)) = v
		p += TAG
	}

}

func (s *Slice) Print() {
	if s == nil {
		return
	}

	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		//fmt.println(*p)
		fmt.Print(*(*int)(unsafe.Pointer(p)))
		p += TAG
	}
}

func (s *Slice) Append(Data ...int) {
	if s == nil {
		return
	}
	if len(Data) == 0 {
		return
	}
	if s.len+len(Data) > s.cap {
		//扩充容量
		s.Data = C.realloc(s.Data, C.ulonglong(s.cap)*2*8)

		s.cap = s.cap * 2
	}

	p := uintptr(s.Data)
	//p = p + TAG * s.len
	for i := 0; i < s.len; i++ {
		//指针偏移
		p += TAG
	}
	for i := 0; i < len(Data); i++ {
		*(*int)(unsafe.Pointer(p)) = Data[i]
		p += TAG
	}
	s.len = s.len + len(Data)
}

func (s *Slice) GetData(index int) int {
	if s == nil || s.Data == nil {
		return 0
	}

	if index < 0 || index > s.len-1 {
		return 0
	}

	p := uintptr(s.Data)

	for i := 0; i < index; i++ {
		p += TAG
	}
	return *(*int)(unsafe.Pointer(p))

}

func (s *Slice) Search(value int) int {
	if s == nil || s.Data == nil {
		return 0
	}

	p := uintptr(s.Data)

	for i := 0; i < s.len; i++ {
		if *(*int)(unsafe.Pointer(p)) == value {
			fmt.Println("search----")
			fmt.Println(i)
			return i
		}
		p += TAG

	}
	return -1

}

func (s *Slice) Delete(index int) int {
	if s == nil || s.Data == nil {
		return 0
	}

	if index < -1 || index > s.len-1 {
		return 0
	}

	p := uintptr(s.Data)

	for i := 0; i < index; i++ {
		p += TAG
	}

	for i := 0; i < s.len-index; i++ {
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(p + TAG))
		p += TAG
	}

	s.len--
	return 0
}

// 插入元素 Insert(下标 元素)
func (s *Slice) Insert(index int, Data int) {
	if s == nil || s.Data == nil {
		return
	}
	if index < 0 || index > s.len-1 {
		return
	}

	//如果插入数据是最后一个
	//if index == s.len-1 {
	//	p := uintptr(s.Data)
	//
	//	//for i := 0; i < s.len; i++ {
	//	//	p += TAG
	//	//}
	//	p += TAG * uintptr(s.len-1)
	//	*(*int)(unsafe.Pointer(p)) = Data
	//	s.len++
	//	return
	//}
	//调用追加方法
	if index == s.len-1 {
		s.Append(Data)
		return
	}

	//获取插入数据的位置
	p := uintptr(s.Data)
	for i := 0; i < index; i++ {
		p += TAG
	}
	//获取末尾的指针位置

	temp := uintptr(s.Data)
	temp += TAG * uintptr(s.len)

	//将后面数据依次向后移动
	for i := s.len; i > index; i-- {
		//用前一个数据为当前数据赋值
		*(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(temp - TAG))
		temp -= TAG
	}

	//修改插入下标的数据
	*(*int)(unsafe.Pointer(p)) = Data
	s.len++
}

// 销毁切片
func (s *Slice) Destroy() {
	//调用C语言  适释放堆空间
	C.free(s.Data)
	s.Data = nil
	s.len = 0
	s.cap = 0
	s = nil
}
