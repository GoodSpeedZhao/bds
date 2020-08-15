package stack

import (
	"errors"
	"math"

	"github.com/GoodSpeedZhao/bds/arrayList"
)

type Stack struct {
	dataStore *arrayList.ArrayList
	capacity  int
}

func NewStack() *Stack {
	return &Stack{
		dataStore: arrayList.NewArrayList(),
		capacity:  math.MaxUint32,
	}
}

func NewStackWithCapacity(capacity int) *Stack {
	if capacity < 0 {
		capacity = 0
	} else if capacity > math.MaxUint32 {
		capacity = math.MaxUint32
	}

	return &Stack{
		dataStore: arrayList.NewArrayList(),
		capacity:  capacity,
	}
}

func (this *Stack) Push(value interface{}) error {
	if this.IsFull() {
		return errors.New("Stack is full.")
	}

	this.dataStore.Prepend(value)
	return nil
}

func (this *Stack) Pop() (value interface{}, err error) {
	if this.IsEmpty() {
		return nil, errors.New("Stack is empty.")
	}

	value, err = this.dataStore.Get(0)
	this.dataStore.Delete(0)
	return
}

func (this *Stack) Peek() (interface{}, error) {
	if this.IsEmpty() {
		return nil, errors.New("Stack is empty.")
	}

	return this.dataStore.Get(0)
}

func (this *Stack) IsEmpty() bool {
	return this.dataStore.IsEmpty()
}

func (this *Stack) IsFull() bool {
	return this.dataStore.Size() == this.capacity
}

func (this *Stack) Clear() {
	this.dataStore.Clear()
}

func (this *Stack) Size() int {
	return this.dataStore.Size()
}
