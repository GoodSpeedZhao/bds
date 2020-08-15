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

func (this *Set) Add(newVal ...interface{}) {
	for _, val := range newVal {
		this.dataStore[val] = struct{}{}
	}
}

func (this *Set) Delete(values ...interface{}) {
	for _, val := range values {
		delete(this.dataStore, val)
	}
}

func (this *Set) Size() int {
	return len(this.dataStore)
}

func (this *Set) Clear() {
	this.dataStore = make(map[interface{}]struct{})
}

func (this *Set) IsEmpty() bool {
	return len(this.dataStore) == 0
}

func (this *Set) Contains(values ...interface{}) bool {
	for _, val := range values {
		if _, ok := this.dataStore[val]; !ok {
			return false
		}
	}
	return true
}

func (this *Set) Intersection(other *Set) *Set {
	if this.IsEmpty() || other.IsEmpty() {
		return nil
	}

	newHashSet := NewHashSet()
	if this.Size() > other.Size() {
		for key := range other.dataStore {
			if _, ok := this.dataStore[key]; ok {
				newHashSet.Add(key)
			}
		}
	} else {
		for key := range this.dataStore {
			if _, ok := other.dataStore[key]; ok {
				newHashSet.Add(key)
			}
		}
	}
	return newHashSet
}

func (this *Set) Union(other *Set) *Set {
	newHashSet := NewHashSet()
	for key := range this.dataStore {
		newHashSet.Add(key)
	}
	for key := range other.dataStore {
		newHashSet.Add(key)
	}
	return newHashSet
}

func (this *Set) Diff(other *Set) *Set {
	newHashSet := NewHashSet()
	for key := range this.dataStore {
		if _, ok := other.dataStore[key]; !ok {
			newHashSet.Add(key)
		}
	}
	return newHashSet
}
