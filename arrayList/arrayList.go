package arrayList

import (
	"errors"
)

type ArrayList struct {
	dataStore []interface{}
	size      int
}

func NewArrayList(newVal ...interface{}) *ArrayList {
	return &ArrayList{
		dataStore: make([]interface{}, len(newVal)),
		size:      len(newVal),
	}
}

func (this *ArrayList) Get(index int) (interface{}, error) {
	if !this.withinRange(index) {
		return nil, errors.New("out of range.")
	}

	return this.dataStore[index], nil
}

func (this *ArrayList) Set(index int, newVal interface{}) error {
	if !this.withinRange(index) {
		return errors.New("out of range.")
	}

	this.dataStore[index] = newVal
	return nil
}

func (this *ArrayList) Insert(index int, newVal ...interface{}) error {
	if !this.withinRange(index) {
		return errors.New("out of range.")
	}

	if index == this.size-1 {
		this.Append(newVal...)
		return nil
	}

	newValLen := len(newVal)

	this.dataStore = append(this.dataStore, newVal...)
	for i := index + 1; i < this.size; i++ {
		this.dataStore[i+newValLen] = this.dataStore[i]
	}

	this.size += newValLen

	startIdx := index + 1
	for i := 0; i < newValLen; i++ {
		this.dataStore[startIdx+i] = newVal[i]
	}

	return nil
}

func (this *ArrayList) Delete(index int) error {
	if !this.withinRange(index) {
		return errors.New("out of range.")
	}

	this.dataStore = append(this.dataStore[:index], this.dataStore[index+1:]...)
	this.size--
	return nil
}

func (this *ArrayList) Append(newVal ...interface{}) {
	this.dataStore = append(this.dataStore, newVal...)
	this.size += len(newVal)
}

func (this *ArrayList) Prepend(newVal ...interface{}) {
	this.dataStore = append(newVal, this.dataStore...)
	this.size += len(newVal)
}

func (this *ArrayList) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}

	if this.size == 0 {
		return false
	}

	for _, val := range values {
		isExist := false
		for _, data := range this.dataStore {
			if val == data {
				isExist = true
				break
			}
		}
		if !isExist {
			return false
		}
	}
	return true
}

func (this *ArrayList) Swap(i, j int) error {
	if i == j {
		return nil
	}

	if !this.withinRange(i) || !this.withinRange(j) {
		return errors.New("out of range.")
	}

	this.dataStore[i], this.dataStore[j] = this.dataStore[j], this.dataStore[i]
	return nil
}

func (this *ArrayList) IsEmpty() bool {
	return this.size == 0
}

func (this *ArrayList) Size() int {
	return this.size
}

func (this *ArrayList) withinRange(index int) bool {
	return index >= 0 && index < this.size
}

func (this *ArrayList) Clear() {
	this.dataStore = []interface{}{}
	this.size = 0
}

func (this *ArrayList) Resize(cap int) error {
	if cap < 0 {
		return errors.New("cap less than zero.")
	}

	newDataStore := make([]interface{}, cap, cap)
	copy(newDataStore, this.dataStore)
	this.dataStore = newDataStore
	this.size = cap
	return nil
}

// func (this *ArrayList) Sort() {
// 	if this.size < 2 {
// 		return
// 	}
// }
