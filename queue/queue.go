package queue

import (
	"errors"
	"math"
)

type Queue struct {
	dataStore []interface{}
	front     int
	rear      int
	size      int
	capacity  int
}

func NewQueue(capacity int) *Queue {
	if capacity < 0 {
		capacity = 0
	} else if capacity > math.MaxUint16 {
		capacity = math.MaxUint16
	}

	return &Queue{
		dataStore: make([]interface{}, capacity, capacity),
		front:     0,
		rear:      -1,
		size:      0,
		capacity:  capacity,
	}
}

func (this *Queue) Push(newVal interface{}) error {
	if this.IsFull() {
		return errors.New("Queue is full.")
	}

	this.size++
	this.rear = (this.rear + 1) % this.capacity
	this.dataStore[this.rear] = newVal
	return nil
}

func (this *Queue) Pop() (value interface{}, err error) {
	if this.IsEmpty() {
		return nil, errors.New("Queue is empty.")
	}

	this.size--
	value, err = this.dataStore[this.front], nil
	this.front = (this.front + 1) % this.capacity
	return
}

func (this *Queue) Front() (value interface{}, err error) {
	if this.IsEmpty() {
		return nil, errors.New("Queue is empty.")
	}

	return this.dataStore[this.front], nil
}

func (this *Queue) Rear() (value interface{}, err error) {
	if this.IsEmpty() {
		return nil, errors.New("Queue is empty.")
	}

	return this.dataStore[this.rear], nil
}

func (this *Queue) IsEmpty() bool {
	return this.size == 0
}

func (this *Queue) IsFull() bool {
	return this.size == this.capacity
}

func (this *Queue) Clear() {
	this.front = 0
	this.rear = 0
	this.size = 0
}

func (this *Queue) Size() int {
	return this.size
}
