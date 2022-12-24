package main

import "fmt"

type StackNode struct {
	Data interface{}
	Next *StackNode
}

func main() {
	s := CreateStack(1, 2, 3)
	PrintStack(s)
	//fmt.Println(LengthStack(s))
	s = Push(s,4)
	 PrintStack(s)
	s = Pop(s)
	 PrintStack(s)

}

//创建链栈
func CreateStack(Data ...interface{}) *StackNode {
	if len(Data) == 0 {
		return nil
	}
	s := new(StackNode)

	//记录下一个节点
	var nextNode *StackNode = nil
	for _, v := range Data { //1,2,3,4,5
		//创建新节点存储数据
		newNode := new(StackNode)
		newNode.Data = v
		s = newNode
		//如果下一个节点不为空 将当前节点的下一个节点设置上一次节点
		//if nextNode != nil {
		s.Next = nextNode
		//}
		//下一个节点为当前节点
		nextNode = s
	}
	return s
}

//打印链栈
func PrintStack(s *StackNode) {
	if s == nil {
		return
	}
	for s != nil {
		fmt.Println(s.Data)
		s = s.Next
	}
}

//链栈个数
func LengthStack(s *StackNode) int {
	if s == nil {
		return 0
	}
	len := 0
	for s != nil {
		len++
		s = s.Next
	}
	return len
}

//入栈
func Push(s *StackNode, Data interface{}) *StackNode {
	if s == nil {
		return nil
	}
	newNode := &StackNode{
		Data: Data,
		Next: s,
	}
	return newNode

}
//入栈
func Pop(s *StackNode) *StackNode {
	if s == nil {
		return nil
	}
	node := s.Next
	s.Data = nil
	s = nil
	return node
}

//清空链栈
func Clear(s *StackNode) *StackNode {
	return nil
}
