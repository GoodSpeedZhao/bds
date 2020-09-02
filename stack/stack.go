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

func (s *Stack) Push(value interface{}) error {
	if s.IsFull() {
		return errors.New("Stack is full.")
	}

	s.dataStore.Prepend(value)
	return nil
}

func (s *Stack) Pop() (value interface{}, err error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty.")
	}

	value, err = s.dataStore.Get(0)
	s.dataStore.Delete(0)
	return
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty.")
	}

	return s.dataStore.Get(0)
}

func (s *Stack) IsEmpty() bool {
	return s.dataStore.IsEmpty()
}

func (s *Stack) IsFull() bool {
	return s.dataStore.Size() == s.capacity
}

func (s *Stack) Clear() {
	s.dataStore.Clear()
}

func (s *Stack) Size() int {
	return s.dataStore.Size()
}
