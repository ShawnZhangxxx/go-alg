package main

import "fmt"

type LinkNode2 struct {
	Next *LinkNode2
	Data interface{}
}

/**
约瑟夫环
*/
func main() {
	list := new(LinkNode2)
	list.Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32)

	fmt.Println("原始数据：")
	list.Print()

	list.Delete(0)
	list.Print()

	//fmt.Println()
	//fmt.Println("删除数据：")
	//i := 1
	//for list.Length() > 2 {
	//	i += 3
	//	if i > list.Length() {
	//		i = list.Length()%3 + 1
	//	}
	//	list.Delete(i)
	//	list.Print()
	//	fmt.Println()
	//}
}

func (l *LinkNode2) Create(data ...interface{}) {
	if l == nil {
		return
	}
	head := l
	for i := 0; i < len(data); i++ {
		var newNode = &LinkNode2{}
		newNode.Data = data[i]
		l.Next = newNode
		l = l.Next
	}
	l.Next = head.Next
}

func (l *LinkNode2) Print() {
	if l == nil {
		return
	}
	start := l.Next
	var node = l.Next
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next

		if start == node {
			break
		}
	}
}

func (l *LinkNode2) Length() int {
	if l == nil {
		return 0
	}
	start := l.Next
	var node = l.Next
	var i int = 0
	for node != nil {
		i++
		node = node.Next
		if start == node {
			break
		}
	}
	return i
}

func (node *LinkNode2) Delete(index int) {
	if node == nil {
		return
	}
	if index < 0 {
		return
	}

	//起始节点
	start := node.Next

	//记录上一个节点
	preNode := node
	//循环找到删除数据的位置
	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next
	}

	//判断如果删除的是第一个节点
	if index == 1 {
		//temp记录最后一个位置
		temp := start
		for {

			if start == temp.Next {
				break
			}

			temp = temp.Next
		}
		//将最后一个节点指向新节点
		temp.Next = node.Next
	}

	//删除数据
	preNode.Next = node.Next

	//数据销毁
	//node.Data = nil
	//node.Next = nil
	node = nil
}

func (l *LinkNode2) Destroy() {
	l = nil
}
