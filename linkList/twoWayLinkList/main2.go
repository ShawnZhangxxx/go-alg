package main

import "fmt"

type LinkNode2 struct {
	Next *LinkNode2
	Prev *LinkNode2
	Data interface{}
}

/**
自己实现的
*/
func main() {
	var ll = &LinkNode2{}
	ll.Create(1, 2, 4)
	//ll.Print()
	ll.Print()
	//ll.Destroy()  //并不能destroy掉
	//ll.Print()
}

func (l *LinkNode2) Create(data ...interface{}) {
	if l == nil {
		return
	}
	for i := 0; i < len(data); i++ {
		var newNode = &LinkNode2{}
		newNode.Data = data[i]
		newNode.Prev = l
		l.Next = newNode
		l = l.Next
	}
}

func (l *LinkNode2) Print() {
	if l == nil {
		return
	}
	var node = l.Next
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}

func (l *LinkNode2) Destroy() {
	l = nil
}
