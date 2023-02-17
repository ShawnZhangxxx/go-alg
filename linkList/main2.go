package main

import "fmt"

type LinkNode2 struct {
	Next *LinkNode2
	Data interface{}
}

/**
	自己实现的
 */
//func main() {
//	var ll = &LinkNode2{}
//	ll.Create(1)
//	//ll.Print()
//	ll.Append(4,5,6)
//	ll.Print()
//	ll.Destroy()  //并不能destroy掉
//	ll.Print()
//}

func (l *LinkNode2) Create(data ...interface{})  {
	if l == nil{
		return
	}

	//var node = l
	for i := 0; i < len(data); i++ {
		var newNode = &LinkNode2{}
		newNode.Data = data[i]
		l.Next = newNode
		l = l.Next
	}
}

func (l *LinkNode2) Print()  {
	if l == nil{
		return
	}
	var node = l.Next
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}

func (l *LinkNode2) Append(data ...interface{})  {
	if l == nil{
		return
	}

	for l.Next != nil {
		l = l.Next
	}
	for i := 0; i < len(data); i++ {
		var newNode = &LinkNode2{}
		newNode.Data = data[i]
		l.Next = newNode
		l = l.Next
	}
}


func (l *LinkNode2) Destroy()  {
	l = nil
}