package main

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
// Related Topics 递归 链表 数学 👍 8618 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var carryBit int = 0
	l3 := &LinkNode{}
	head := l3
	for l1 != nil && l2 != nil {
		if l1 != nil && l2 != nil {
			l3.Data = (l1.Data.(int) + l2.Data.(int) + carryBit) % 10
			carryBit = (l1.Data.(int) + l2.Data.(int) + carryBit) / 10
			l1 = l1.Next
			l2 = l2.Next
			newNode := &LinkNode{}
			l3.Next = newNode
			l3 = l3.Next
		} else { //处理长短不一致的部分

		}
	}
	return head
}
/**
一个聪明做法是加了两个0,这样很好处理一条链结束问题,也避免开头判断问题
复用了sum =  n1 + n2 + carry

*/
//func addTwoNumbers(l1, l2 *LinkNode) (head *LinkNode) {
//	var tail *LinkNode
//	carry := 0
//	for l1 != nil || l2 != nil {
//		n1, n2 := 0, 0
//		if l1 != nil {
//			n1 = l1.Data
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			n2 = l2.Data
//			l2 = l2.Next
//		}
//		sum := n1 + n2 + carry
//		sum, carry = sum%10, sum/10
//		if head == nil {
//			head = &LinkNode{Data: sum}
//			tail = head
//		} else {
//			tail.Next = &LinkNode{Data: sum}
//			tail = tail.Next
//		}
//	}
//	if carry > 0 {
//		tail.Next = &LinkNode{Data: carry}
//	}
//	return
//}

//func main() {
//	var list1 *LinkNode = new(LinkNode)
//	var list2 *LinkNode = new(LinkNode)
//	//list.Create(1,2,3,4,5)
//	//创建链表
//	list1.Create(2, 4, 3)
//	//list1.Print()
//	list2.Create(5, 6, 4)
//	list3 := addTwoNumbers(list1.Next, list2.Next)
//	list3.Print()
//}

//leetcode submit region end(Prohibit modification and deletion)
