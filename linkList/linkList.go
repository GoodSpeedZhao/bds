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

func (linkList *LinkList) Get(index int) (interface{}, error) {
	if !linkList.withinRange(index) {
		return nil, errors.New("out of range.")
	}

	e := linkList.head
	for i := 0; i < index; i, e = i+1, e.Next {
	}
	return e.Val, nil
}

func (linkList *LinkList) Set(index int, newVal interface{}) error {
	if !linkList.withinRange(index) {
		return errors.New("out of range.")
	}

	if index == linkList.size-1 {
		linkList.tail.Val = newVal
		return nil
	}

	e := linkList.head
	for i := 0; i < index; i, e = i+1, e.Next {
	}
	e.Val = newVal
	return nil
}

func (linkList *LinkList) Insert(index int, newVal ...interface{}) error {
	if !linkList.withinRange(index) {
		return errors.New("out of range.")
	}

	if index == linkList.size-1 {
		linkList.Append(newVal...)
		return nil
	}

	e := linkList.head
	for i := 0; i < index; i, e = i+1, e.Next {
	}

	for _, val := range newVal {
		tmp := NewNode(val, e.Next)
		e.Next = tmp
		e = tmp
	}

	linkList.size += len(newVal)

	return nil
}

func (linkList *LinkList) Delete(index int) error {
	if !linkList.withinRange(index) {
		return errors.New("out of range.")
	}

	if linkList.size == 1 {
		linkList.Clear()
		return nil
	}

	var preElement *Node
	e := linkList.head
	for i := 0; i < index; i, e = i+1, e.Next {
		preElement = e
	}

	if e == linkList.head {
		linkList.head = linkList.head.Next
	} else if e == linkList.tail {
		linkList.tail = preElement
	} else {
		preElement.Next = e.Next
	}

	e.Next = nil
	e = nil

	linkList.size--

	return nil
}

func (linkList *LinkList) Append(newVal ...interface{}) {
	for _, val := range newVal {
		tmp := NewNode(val, nil)
		if linkList.head == nil && linkList.tail == nil {
			linkList.head = tmp
			linkList.tail = tmp
		} else {
			linkList.tail.Next = tmp
			linkList.tail = tmp
		}
	}

	linkList.size += len(newVal)
}

func (linkList *LinkList) Prepend(newVal ...interface{}) {
	for _, val := range newVal {
		tmp := NewNode(val, linkList.head)

		linkList.head = tmp
		if linkList.tail == nil {
			linkList.tail = tmp
		}
	}
	linkList.size += len(newVal)
}

func (linkList *LinkList) Clear() {
	linkList.head = nil
	linkList.tail = nil
	linkList.size = 0
}

func (linkList *LinkList) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}

	if linkList.size == 0 {
		return false
	}

	for _, val := range values {
		var e *Node
		for e = linkList.head; e != nil; e = e.Next {
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

func (linkList *LinkList) Swap(i, j int) error {
	if i == j {
		return nil
	}

	if !linkList.withinRange(i) || !linkList.withinRange(j) {
		return errors.New("out of range.")
	}

	var elementI, elementJ *Node
	for count, e := 0, linkList.head; elementI == nil || elementJ == nil; count, e = count+1, e.Next {
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

func (linkList *LinkList) IsEmpty() bool {
	return linkList.size == 0
}

func (linkList *LinkList) Size() int {
	return linkList.size
}

func (linkList *LinkList) withinRange(index int) bool {
	return index >= 0 && index < linkList.size
}

func (linkList *LinkList) HeadNode() *Node {
	return linkList.head
}

func (linkList *LinkList) TailNode() *Node {
	return linkList.tail
}
