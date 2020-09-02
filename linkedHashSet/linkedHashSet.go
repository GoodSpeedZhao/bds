package linkedHashSet

import "github.com/GoodSpeedZhao/bds/linkedHashMap"

type Set struct {
	dataStore *linkedHashMap.Map
}

func NewLinkedHashSet() *Set {
	return &Set{
		dataStore: linkedHashMap.NewLinkedHashMap(),
	}
}

func (s *Set) Add(newVal ...interface{}) {
	for _, val := range newVal {
		s.dataStore.Put(val, struct{}{})
	}
}

func (s *Set) Delete(values ...interface{}) {
	for _, val := range values {
		s.dataStore.Delete(val)
	}
}

func (s *Set) IsEmpty() bool {
	return s.dataStore.IsEmpty()
}

func (s *Set) Size() int {
	return s.dataStore.Size()
}

func (s *Set) Clear() {
	s.dataStore.Clear()
}
