package main

import (
	"fmt"
	"reflect"
)

type BinaryNode struct {
	Data      interface{}
	LeftNode  *BinaryNode
	RightNode *BinaryNode
}

func main() {
	tree := &BinaryNode{0, nil, nil}
	node1 := BinaryNode{1, nil, nil}
	node2 := BinaryNode{2, nil, nil}
	node3 := BinaryNode{3, nil, nil}
	node4 := BinaryNode{4, nil, nil}
	node5 := BinaryNode{5, nil, nil}
	node6 := BinaryNode{6, nil, nil}
	node7 := BinaryNode{7, nil, nil}

	tree.LeftNode = &node1
	tree.RightNode = &node2
	node1.LeftNode = &node3
	node1.RightNode = &node4
	node2.LeftNode = &node5
	node2.RightNode = &node6
	node3.LeftNode = &node7
	//tree.preOrder()
	//tree.middleOrder()
	//tree.RearOrder()
	fmt.Println(tree.TreeHeight())
	fmt.Println(tree.LeafCount())
}

func (tree *BinaryNode) preOrder() {
	if tree == nil {
		return
	}

	fmt.Println(tree.Data)
	tree.LeftNode.preOrder()
	tree.RightNode.preOrder()

}

func (tree *BinaryNode) middleOrder() {
	if tree == nil {
		return
	}

	tree.LeftNode.middleOrder()
	fmt.Println(tree.Data)

	tree.RightNode.middleOrder()

}

func (tree *BinaryNode) RearOrder() {
	if tree == nil {
		return
	}
	tree.LeftNode.RearOrder()
	tree.RightNode.RearOrder()
	fmt.Println(tree.Data)
}

func (tree *BinaryNode) TreeHeight() int {
	if tree == nil {
		return 0
	}
	lt := tree.LeftNode.TreeHeight()
	rt := tree.RightNode.TreeHeight()

	if lt > rt {
		lt++
		return lt
	} else {
		rt++
		return rt
	}

}

//二叉树叶子节点个数
//叶子节点是没有后继的节点
func (tree *BinaryNode) LeafCount() int {
	if tree == nil {
		return 0
	}
	leftC := tree.LeftNode.LeafCount()
	rightC := tree.RightNode.LeafCount()

	if tree.LeftNode == nil && tree.RightNode == nil {
		return leftC + rightC + 1
	}

	return leftC + rightC
}

//二叉树数据查找
func (node *BinaryNode) Search(Data interface{}) {
	if node == nil {
		return
	}

	//== 不支持slice 和 map
	//reflect.DeepEqual()
	if reflect.TypeOf(node.Data) == reflect.TypeOf(Data) && node.Data == Data {
		fmt.Println("找到数据", node.Data)
		return
	}
	node.LeftNode.Search(Data)
	node.RightNode.Search(Data)

}

//二叉树销毁
func (node *BinaryNode) Destroy() {
	if node == nil {
		return
	}

	node.LeftNode.Destroy()
	node.LeftNode = nil
	node.RightNode.Destroy()
	node.RightNode = nil
	node.Data = nil

}

//二叉树反转
//如果想反转二叉树 二叉树必须是一个满二叉树
func (node *BinaryNode) Reverse() {
	if node == nil {
		return
	}

	//交换节点  多重赋值
	node.LeftNode, node.RightNode = node.RightNode, node.LeftNode

	node.LeftNode.Reverse()
	node.RightNode.Reverse()

}

//二叉树拷贝
func (node *BinaryNode) Copy() *BinaryNode {
	if node == nil {
		return nil
	}
	l := node.LeftNode.Copy()
	r := node.RightNode.Copy()

	newNode := new(BinaryNode)
	newNode.RightNode = r
	newNode.LeftNode = l
	newNode.Data = node.Data

	return newNode

}
