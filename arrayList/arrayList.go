package arrayList

import (
	"errors"
)

type ArrayList struct {
	dataStore []interface{}
	size      int
}

func NewArrayList(newVal ...interface{}) *ArrayList {
	arrayList := &ArrayList{
		size: len(newVal),
	}

	if newVal != nil {
		arrayList.dataStore = make([]interface{}, len(newVal))
		copy(arrayList.dataStore, newVal)
	} else {
		arrayList.dataStore = []interface{}{}
	}

	return arrayList
}

func (array *ArrayList) Get(index int) (interface{}, error) {
	if !array.withinRange(index) {
		return nil, errors.New("out of range.")
	}

	return array.dataStore[index], nil
}

func (array *ArrayList) Set(index int, newVal interface{}) error {
	if !array.withinRange(index) {
		return errors.New("out of range.")
	}

	array.dataStore[index] = newVal
	return nil
}

func (array *ArrayList) Insert(index int, newVal ...interface{}) error {
	if !array.withinRange(index) {
		return errors.New("out of range.")
	}

	if index == array.size-1 {
		array.Append(newVal...)
		return nil
	}

	newValLen := len(newVal)

	array.dataStore = append(array.dataStore, newVal...)
	for i := index + 1; i < array.size; i++ {
		array.dataStore[i+newValLen] = array.dataStore[i]
	}

	array.size += newValLen

	startIdx := index + 1
	for i := 0; i < newValLen; i++ {
		array.dataStore[startIdx+i] = newVal[i]
	}

	return nil
}

func (array *ArrayList) Delete(index int) error {
	if !array.withinRange(index) {
		return errors.New("out of range.")
	}

	array.dataStore = append(array.dataStore[:index], array.dataStore[index+1:]...)
	array.size--
	return nil
}

func (array *ArrayList) Append(newVal ...interface{}) {
	array.dataStore = append(array.dataStore, newVal...)
	array.size += len(newVal)
}

func (array *ArrayList) Prepend(newVal ...interface{}) {
	for i := 0; i < len(newVal); i++ {
		array.dataStore = append([]interface{}{newVal[i]}, array.dataStore...)
	}
	array.size += len(newVal)
}

func (array *ArrayList) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}

	if array.size == 0 {
		return false
	}

	for _, val := range values {
		isExist := false
		for _, data := range array.dataStore {
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

func (array *ArrayList) Swap(i, j int) error {
	if !array.withinRange(i) || !array.withinRange(j) {
		return errors.New("out of range.")
	}

	if i == j {
		return nil
	}

	array.dataStore[i], array.dataStore[j] = array.dataStore[j], array.dataStore[i]
	return nil
}

func (array *ArrayList) IsEmpty() bool {
	return array.size == 0
}

func (array *ArrayList) Size() int {
	return array.size
}

func (array *ArrayList) withinRange(index int) bool {
	return index >= 0 && index < array.size
}

func (array *ArrayList) Clear() {
	array.dataStore = []interface{}{}
	array.size = 0
}

func (array *ArrayList) Resize(cap int) error {
	if cap < 0 {
		return errors.New("cap less than zero.")
	}

	newDataStore := make([]interface{}, cap, cap)
	copy(newDataStore, array.dataStore)
	array.dataStore = newDataStore
	array.size = cap
	return nil
}
