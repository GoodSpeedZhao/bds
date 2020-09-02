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

func (m *Map) Put(key interface{}, value interface{}) {
	if val, ok := m.dataStore[key]; ok {
		val.value = value
	} else {
		e := &entry{
			key:   key,
			value: value,
		}

		m.dataStore[key] = e
		if m.head == nil && m.tail == nil {
			m.head = e
			m.tail = e
		} else {
			e.pre = m.tail
			m.tail.next = e
			m.tail = e
		}
	}
}

func (m *Map) Get(key interface{}) (val interface{}, ok bool) {
	e, ok := m.dataStore[key]
	if ok {
		val = e.value
	}
	return
}

func (m *Map) Delete(key interface{}) {
	val, ok := m.dataStore[key]
	if !ok {
		return
	}

	if val == m.head {
		m.head = m.head.next
	} else if val == m.tail {
		m.tail = m.tail.pre
	} else {
		val.pre.next = val.next
		val.next.pre = val.pre
	}

	val.pre = nil
	val.next = nil
	delete(m.dataStore, key)
}

func (m *Map) Keys() []interface{} {
	keys := make([]interface{}, m.Size())
	count := 0
	for e := m.head; e != nil; e = e.next {
		keys[count] = e.key
		count++
	}
	return keys
}

func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	for e := m.head; e != nil; e = e.next {
		values[count] = e.value
		count++
	}
	return values
}

func (m *Map) Size() int {
	return len(m.dataStore)
}

func (m *Map) IsEmpty() bool {
	return m.Size() == 0
}

func (m *Map) Clear() {
	m.dataStore = make(map[interface{}]*entry)
	m.head = nil
	m.tail = nil
}
