package main

import "fmt"

type LinkNode2 struct {
	Next *LinkNode2
	Data interface{}
}

/**
自己实现的
*/
func main() {
	var ll = &LinkNode2{}
	ll.Create(1)
	//ll.Print()
	ll.Append(2, 3, 4)
	ll.Print()
	//ll.Destroy()  //并不能destroy掉 l在方法里设置为nil并不能,真正设置为nil
	//ll.Print()
	ll = ll.Reverse()
	ll.Print()
}

func (l *LinkNode2) Create(data ...interface{}) {
	if l == nil {
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

func (l *LinkNode2) Append(data ...interface{}) {
	if l == nil {
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

/**
没用
*/
func (l *LinkNode2) Destroy() {
	l = nil
}

/**
反转链表
*/
func (l *LinkNode2) Reverse() *LinkNode2 {
	if l == nil {
		return nil
	}

	l = l.Next

	var prev *LinkNode2 = nil
	var curNode *LinkNode2
	for l != nil {
		curNode = l
		l = l.Next
		curNode.Next = prev
		prev = curNode
	}

	var newNode = &LinkNode2{}
	newNode.Next = curNode

	return newNode
}
