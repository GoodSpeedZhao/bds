package dlinkList

import (
	"errors"
)

type DLinkList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  interface{}
	next *Node
	pre  *Node
}

func NewDLinkList() *DLinkList {
	return &DLinkList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func NewNode(newVal interface{}) *Node {
	return &Node{
		val:  newVal,
		next: nil,
		pre:  nil,
	}
}

func (dll *DLinkList) Get(index int) (interface{}, error) {
	if !dll.withinRange(index) {
		return nil, errors.New("out of range.")
	}

	if dll.size-index > index {
		e := dll.head
		for i := 0; i < index; i, e = i+1, e.next {
		}
		return e.val, nil
	} else {
		e := dll.tail
		for i := dll.size - 1; i > index; i, e = i-1, e.pre {
		}
		return e.val, nil
	}
}

func (dll *DLinkList) Set(index int, newVal interface{}) error {
	if !dll.withinRange(index) {
		return errors.New("out of range.")
	}

	if dll.size-index > index {
		e := dll.head
		for i := 0; i < index; i, e = i+1, e.next {
		}
		e.val = newVal
	} else {
		e := dll.tail
		for i := dll.size - 1; i > index; i, e = i-1, e.pre {
		}
		e.val = newVal
	}
	return nil
}

func (dll *DLinkList) Insert(index int, newVal ...interface{}) error {
	if !dll.withinRange(index) {
		return errors.New("out of range.")
	}

	if index == dll.size-1 {
		dll.Append(newVal...)
		return nil
	}

	var e *Node
	if dll.size-index > index {
		e = dll.head
		for i := 0; i < index; i, e = i+1, e.next {
		}
	} else {
		e = dll.tail
		for i := dll.size - 1; i > index; i, e = i-1, e.pre {
		}
	}

	for _, val := range newVal {
		tmp := new(Node)
		tmp.val = val
		tmp.pre = e
		tmp.next = e.next

		e.next.pre = tmp
		e.next = tmp
		e = tmp
	}

	dll.size += len(newVal)

	return nil
}

func (dll *DLinkList) Delete(index int) error {
	if !dll.withinRange(index) {
		return errors.New("out of range.")
	}

	if dll.size == 1 {
		dll.Clear()
		return nil
	}

	var e *Node
	if dll.size-index > index {
		e = dll.head
		for i := 0; i < index; i, e = i+1, e.next {
		}
	} else {
		e = dll.tail
		for i := dll.size - 1; i > index; i, e = i-1, e.pre {
		}
	}
	if e == dll.head {
		dll.head = e.next
		dll.head.pre = nil

		e.pre = nil
		e.next = nil
	} else if e == dll.tail {
		dll.tail = e.pre
		dll.tail.next = nil

		e.pre = nil
		e.next = nil
	} else {
		e.pre.next = e.next
		e.next.pre = e.pre

		e.pre = nil
		e.next = nil
	}
	e = nil

	dll.size--

	return nil
}

func (dll *DLinkList) DeleteNode(node *Node) error {
	if node == nil {
		return errors.New("node is nil.")
	}

	if node == dll.head {
		dll.head = node.next
		if dll.head != nil {
			dll.head.pre = nil
		}

		node.pre = nil
		node.next = nil
	} else if node == dll.tail {
		dll.tail = node.pre
		dll.tail.next = nil

		node.pre = nil
		node.next = nil
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre

		node.pre = nil
		node.next = nil
	}

	dll.size--

	return nil
}

func (dll *DLinkList) Append(newVal ...interface{}) {
	e := dll.tail
	for _, val := range newVal {
		tmp := new(Node)
		tmp.val = val
		if dll.head == nil && dll.tail == nil {
			dll.head = tmp
			dll.tail = tmp
			e = tmp
		} else {
			tmp.pre = e
			e.next = tmp
			e = tmp
		}
	}
	dll.tail = e

	dll.size += len(newVal)
}

func (dll *DLinkList) Prepend(newVal ...interface{}) {
	e := dll.head
	for _, val := range newVal {
		tmp := new(Node)
		tmp.val = val
		if dll.head == nil && dll.tail == nil {
			dll.head = tmp
			dll.tail = tmp
			e = tmp
		} else {
			tmp.next = e
			e.pre = tmp
			e = tmp
		}
	}
	dll.head = e

	dll.size += len(newVal)
}

func (dll *DLinkList) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}

	if dll.size == 0 {
		return false
	}

	for _, val := range values {
		var e *Node
		for e = dll.head; e != nil; e = e.next {
			if val == e.val {
				break
			}
		}
		if e == nil {
			return false
		}
	}

	return true
}

func (dll *DLinkList) Clear() {
	dll.head = nil
	dll.tail = nil
	dll.size = 0
}

func (dll *DLinkList) Swap(i, j int) error {
	if !dll.withinRange(i) || !dll.withinRange(j) {
		return errors.New("out of range.")
	}

	if i == j {
		return nil
	}

	var elementI, elementJ *Node
	for count, e := 0, dll.head; elementI == nil || elementJ == nil; count, e = count+1, e.next {
		switch count {
		case i:
			elementI = e
		case j:
			elementJ = e
		}
	}
	elementI.val, elementJ.val = elementJ.val, elementI.val
	return nil
}

func (dll *DLinkList) Size() int {
	return dll.size
}

func (dll *DLinkList) IsEmpty() bool {
	return dll.size == 0
}

func (dll *DLinkList) withinRange(index int) bool {
	return index >= 0 && index < dll.size
}

func (dll *DLinkList) HeadNode() *Node {
	return dll.head
}

func (dll *DLinkList) TailNode() *Node {
	return dll.tail
}
