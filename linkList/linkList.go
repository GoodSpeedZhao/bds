package linkList

import (
	"errors"
)

type LinkList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	Val  interface{}
	Next *Node
}

func NewLinkList() *LinkList {
	return &LinkList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func NewNode(newVal interface{}, newNext *Node) *Node {
	return &Node{
		Val:  newVal,
		Next: newNext,
	}
}

func (this *LinkList) Get(index int) (interface{}, error) {
	if !this.withinRange(index) {
		return nil, errors.New("out of range.")
	}

	e := this.head
	for i := 0; i < index; i, e = i+1, e.Next {
	}
	return e.Val, nil
}

func (this *LinkList) Set(index int, newVal interface{}) error {
	if !this.withinRange(index) {
		return errors.New("out of range.")
	}

	if index == this.size-1 {
		this.tail.Val = newVal
		return nil
	}

	e := this.head
	for i := 0; i < index; i, e = i+1, e.Next {
	}
	e.Val = newVal
	return nil
}

func (this *LinkList) Insert(index int, newVal ...interface{}) error {
	if !this.withinRange(index) {
		return errors.New("out of range.")
	}

	if index == this.size-1 {
		this.Append(newVal...)
		return nil
	}

	e := this.head
	for i := 0; i < index; i, e = i+1, e.Next {
	}

	for _, val := range newVal {
		tmp := NewNode(val, e.Next)
		e.Next = tmp
		e = tmp
	}

	this.size += len(newVal)

	return nil
}

func (this *LinkList) Delete(index int) error {
	if !this.withinRange(index) {
		return errors.New("out of range.")
	}

	if this.size == 1 {
		this.Clear()
		return nil
	}

	var preElement *Node
	e := this.head
	for i := 0; i < index; i, e = i+1, e.Next {
		preElement = e
	}

	if e == this.head {
		this.head = this.head.Next
	} else if e == this.tail {
		this.tail = preElement
	} else {
		preElement.Next = e.Next
	}

	e.Next = nil
	e = nil

	this.size--

	return nil
}

func (this *LinkList) Append(newVal ...interface{}) {
	for _, val := range newVal {
		tmp := NewNode(val, nil)
		if this.head == nil && this.tail == nil {
			this.head = tmp
			this.tail = tmp
		} else {
			this.tail.Next = tmp
			this.tail = tmp
		}
	}

	this.size += len(newVal)
}

func (this *LinkList) Prepend(newVal ...interface{}) {
	for _, val := range newVal {
		tmp := NewNode(val, this.head)

		this.head = tmp
		if this.tail == nil {
			this.tail = tmp
		}
	}
	this.size += len(newVal)
}

func (this *LinkList) Clear() {
	this.head = nil
	this.tail = nil
	this.size = 0
}

func (this *LinkList) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}

	if this.size == 0 {
		return false
	}

	for _, val := range values {
		var e *Node
		for e = this.head; e != nil; e = e.Next {
			if val == e.Val {
				break
			}
		}
		if e == nil {
			return false
		}
	}

	return true
}

func (this *LinkList) Swap(i, j int) error {
	if i == j {
		return nil
	}

	if !this.withinRange(i) || !this.withinRange(j) {
		return errors.New("out of range.")
	}

	var elementI, elementJ *Node
	for count, e := 0, this.head; elementI == nil || elementJ == nil; count, e = count+1, e.Next {
		switch count {
		case i:
			elementI = e
		case j:
			elementJ = e
		}
	}
	elementI.Val, elementJ.Val = elementJ.Val, elementI.Val
	return nil
}

func (this *LinkList) IsEmpty() bool {
	return this.size == 0
}

func (this *LinkList) Size() int {
	return this.size
}

func (this *LinkList) withinRange(index int) bool {
	return index >= 0 && index < this.size
}

func (this *LinkList) HeadNode() *Node {
	return this.head
}

func (this *LinkList) TailNode() *Node {
	return this.tail
}
