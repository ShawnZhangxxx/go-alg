package linkList

import (
	"unsafe"
)

type LinkList struct {
	len   int
	start *Node
}

type Node struct {
	//prev unsafe.Pointer
	next unsafe.Pointer
	data int
}

func (ll *LinkList) Create() {
	var node = &Node{
		next: nil,
		data: nil,
	}
	ll.start = node
}
func (ll *LinkList) append(data int) {
	if ll.start == nil {
		var node = &Node{
			next: nil,
			data: data,
		}
		ll.start = node
	} else {
		for ll.start.next != nil {

		}
	}

}
