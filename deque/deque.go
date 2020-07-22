package deque

import (
	"container/list"
	"errors"
	"math"
)

type Deque struct {
	dataStore *list.List
	capacity  int
}

func NewDeque() *Deque {
	return &Deque{
		dataStore: list.New(),
		capacity:  math.MaxUint32,
	}
}

func NewDequeWithCapacity(cap int) *Deque {
	return &Deque{
		dataStore: list.New(),
		capacity:  cap,
	}
}

func (this *Deque) Append(newVal interface{}) bool {
	if this.IsFull() {
		return false
	}

	this.dataStore.PushBack(newVal)
	return true
}

func (this *Deque) Prepend(newVal interface{}) bool {
	if this.IsFull() {
		return false
	}

	this.dataStore.PushFront(newVal)
	return true
}

func (this *Deque) Shift() (interface{}, error) {
	if this.IsEmpty() {
		return nil, errors.New("Deque is empty.")
	}

	return this.dataStore.Remove(this.dataStore.Front()), nil
}

func (this *Deque) Pop() (interface{}, error) {
	if this.IsEmpty() {
		return nil, errors.New("Deque is empty.")
	}

	return this.dataStore.Remove(this.dataStore.Back()), nil
}

func (this *Deque) Front() (interface{}, error) {
	if this.IsEmpty() {
		return nil, errors.New("Deque is empty.")
	}

	return this.dataStore.Front().Value, nil
}

func (this *Deque) Rear() (interface{}, error) {
	if this.IsEmpty() {
		return nil, errors.New("Deque is empty.")
	}

	return this.dataStore.Back().Value, nil
}

func (this *Deque) Size() int {
	return this.dataStore.Len()
}

func (this *Deque) Capacity() int {
	return this.capacity
}

func (this *Deque) IsEmpty() bool {
	return this.Size() == 0
}

func (this *Deque) IsFull() bool {
	return this.Size() >= this.capacity
}
