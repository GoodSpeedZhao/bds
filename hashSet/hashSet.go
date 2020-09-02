package hashSet

type Set struct {
	dataStore map[interface{}]struct{}
}

func NewHashSet(newVal ...interface{}) *Set {
	set := &Set{
		dataStore: make(map[interface{}]struct{}),
	}
	if len(newVal) > 0 {
		set.Add(newVal...)
	}

	return set
}

func (s *Set) Add(newVal ...interface{}) {
	for _, val := range newVal {
		s.dataStore[val] = struct{}{}
	}
}

func (s *Set) Delete(values ...interface{}) {
	for _, val := range values {
		delete(s.dataStore, val)
	}
}

func (s *Set) Size() int {
	return len(s.dataStore)
}

func (s *Set) Clear() {
	s.dataStore = make(map[interface{}]struct{})
}

func (s *Set) IsEmpty() bool {
	return len(s.dataStore) == 0
}

func (s *Set) Contains(values ...interface{}) bool {
	for _, val := range values {
		if _, ok := s.dataStore[val]; !ok {
			return false
		}
	}
	return true
}

func (s *Set) Intersection(other *Set) *Set {
	if s.IsEmpty() || other.IsEmpty() {
		return nil
	}

	newHashSet := NewHashSet()
	if s.Size() > other.Size() {
		for key := range other.dataStore {
			if _, ok := s.dataStore[key]; ok {
				newHashSet.Add(key)
			}
		}
	} else {
		for key := range s.dataStore {
			if _, ok := other.dataStore[key]; ok {
				newHashSet.Add(key)
			}
		}
	}
	return newHashSet
}

func (s *Set) Union(other *Set) *Set {
	newHashSet := NewHashSet()
	for key := range s.dataStore {
		newHashSet.Add(key)
	}
	for key := range other.dataStore {
		newHashSet.Add(key)
	}
	return newHashSet
}

func (s *Set) Diff(other *Set) *Set {
	newHashSet := NewHashSet()
	for key := range s.dataStore {
		if _, ok := other.dataStore[key]; !ok {
			newHashSet.Add(key)
		}
	}
	return newHashSet
}
