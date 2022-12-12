package linkList

type LinkList struct {
	len   int
	start *Node
}

type Node struct {
	//prev unsafe.Pointer
	next *Node
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
			ll.start = ll.start.next
		}
		ll.start.data = data
	}

}
