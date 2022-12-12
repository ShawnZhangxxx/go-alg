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

<<<<<<< HEAD

//Print 打印切片
func (s *Slice) Print() {
	if s == nil {
		return
	}

	//将万能指针转成可以计算的指针
	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		//获取内存中的数据
		fmt.Print(*(*int)(unsafe.Pointer(p)), " ")
		p += TAG
	}

}

//切片追加
func (s *Slice) Append(Data ...int) {
	if s == nil {
		return
	}
	if len(Data) == 0 {
		return
	}

	//如果添加的数据超出了容量
	if s.len+len(Data) > s.cap {
		//扩充容量
		//C.realloc(指针,字节大小)
		s.Data = C.realloc(s.Data, C.ulonglong(s.cap)*2*8)
		//改变容量的值
		s.cap = s.cap * 2
	}

	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		//指针偏移
		p += TAG
	}

	//添加数据
	for _, v := range Data {
		*(*int)(unsafe.Pointer(p)) = v
		p += TAG
	}
	//更新有效数据（长度）
	s.len = s.len + len(Data)
}

//获取元素 GetData(下标)  返回值为int 元素
func (s *Slice) GetdData(index int) int {
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

//查找元素 Search(元素)返回值为int 下标
func (s *Slice) Search(Data int) int {
	if s == nil || s.Data == nil {
		return -1
	}

	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		//查找数据  返回第一次元素出现的位置
		if *(*int)(unsafe.Pointer(p)) == Data {
			return i
		}
		//指针偏移
		p += TAG
	}
	return -1
}

//删除元素 Delete(下标)
func (s *Slice) Delete(index int) {
	if s == nil || s.Data == nil {
		return
	}
	if index < 0 || index > s.len-1 {
		return
	}

	//将指针指向需要删除的下标位置
	p := uintptr(s.Data)
=======
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

>>>>>>> f827a68fa064ac5025d1d968adb3a95a65344d76
	for i := 0; i < index; i++ {
		p += TAG
	}

<<<<<<< HEAD
	//用下一个指针对应的值为当前指针对应的值进行赋值
	temp := p
	for i := index; i < s.len; i++ {
		temp += TAG
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(temp))
=======
	for i := 0; i < s.len-index; i++ {
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(p + TAG))
>>>>>>> f827a68fa064ac5025d1d968adb3a95a65344d76
		p += TAG
	}

	s.len--
<<<<<<< HEAD
}

//插入元素 Insert(下标 元素)
=======
	return 0
}

// 插入元素 Insert(下标 元素)
>>>>>>> f827a68fa064ac5025d1d968adb3a95a65344d76
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

<<<<<<< HEAD
//销毁切片
=======
// 销毁切片
>>>>>>> f827a68fa064ac5025d1d968adb3a95a65344d76
func (s *Slice) Destroy() {
	//调用C语言  适释放堆空间
	C.free(s.Data)
	s.Data = nil
	s.len = 0
	s.cap = 0
	s = nil
}
