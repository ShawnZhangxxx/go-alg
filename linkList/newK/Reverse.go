package main

 type ListNode struct{
   Val int
   Next *ListNode
 }


/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList( pHead *ListNode ) *ListNode {
	// write code here
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	var newHead  *ListNode
	for pHead != nil {
		nextNode := pHead.Next
		pHead.Next = newHead
		newHead  = pHead
		pHead  = nextNode
	}
	return newHead
}

func ReverseList( pHead *ListNode ) *ListNode {
	// write code here
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	var newNode *ListNode
	if pHead != nil {
		nextNode := pHead.Next
		pHead.Next = newNode
		newNode = pHead
		pHead = nextNode
	}
	return newNode
}
func ReverseList( pHead *ListNode ) *ListNode {
	// write code here
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	var newHead  *ListNode
	for pHead != nil {
		nextNode := pHead.Next
		pHead.Next = newHead
		newHead  = pHead
		pHead  = nextNode
	}

	for i := 0; i < ; i++ {
		
	}
	return newHead
}