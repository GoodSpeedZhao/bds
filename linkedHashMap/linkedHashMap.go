package linkedHashMap

type entry struct {
	key   interface{}
	value interface{}
	pre   *entry
	next  *entry
}

type Map struct {
	dataStore map[interface{}]*entry
	head      *entry
	tail      *entry
}

func NewLinkedHashMap() *Map {
	return &Map{
		dataStore: make(map[interface{}]*entry),
		head:      nil,
		tail:      nil,
	}
}

func (this *Map) Put(key interface{}, value interface{}) {
	if val, ok := this.dataStore[key]; ok {
		val.value = value
	} else {
		e := &entry{
			key:   key,
			value: value,
		}

		this.dataStore[key] = e
		if this.head == nil && this.tail == nil {
			this.head = e
			this.tail = e
		} else {
			e.pre = this.tail
			this.tail.next = e
			this.tail = e
		}
	}
}

func (this *Map) Get(key interface{}) (val interface{}, ok bool) {
	e, ok := this.dataStore[key]
	if ok {
		val = e.value
	}
	return
}

func (this *Map) Delete(key interface{}) {
	val, ok := this.dataStore[key]
	if !ok {
		return
	}

	if val == this.head {
		this.head = this.head.next
	} else if val == this.tail {
		this.tail = this.tail.pre
	} else {
		val.pre.next = val.next
		val.next.pre = val.pre
	}

	val.pre = nil
	val.next = nil
	delete(this.dataStore, key)
}

func (this *Map) Keys() []interface{} {
	keys := make([]interface{}, this.Size())
	count := 0
	for e := this.head; e != nil; e = e.next {
		keys[count] = e.key
		count++
	}
	return keys
}

func (this *Map) Values() []interface{} {
	values := make([]interface{}, this.Size())
	count := 0
	for e := this.head; e != nil; e = e.next {
		values[count] = e.value
		count++
	}
	return values
}

func (this *Map) Size() int {
	return len(this.dataStore)
}

func (this *Map) IsEmpty() bool {
	return this.Size() == 0
}

func (this *Map) Clear() {
	this.dataStore = make(map[interface{}]*entry)
	this.head = nil
	this.tail = nil
}
