package hashMap

type Map struct {
	dataStore map[interface{}]interface{}
}

func NewHashMap() *Map {
	return &Map{
		dataStore: make(map[interface{}]interface{}),
	}
}

func (m *Map) Put(key interface{}, value interface{}) {
	m.dataStore[key] = value
}

func (m *Map) Get(key interface{}) (val interface{}, ok bool) {
	val, ok = m.dataStore[key]
	return
}

func (m *Map) Delete(key interface{}) {
	delete(m.dataStore, key)
}

func (m *Map) Size() int {
	return len(m.dataStore)
}

func (m *Map) IsEmpty() bool {
	return m.Size() == 0
}

func (m *Map) Keys() []interface{} {
	keys := make([]interface{}, m.Size())
	count := 0
	for key := range m.dataStore {
		keys[count] = key
		count++
	}
	return keys
}

func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	for _, val := range m.dataStore {
		values[count] = val
		count++
	}
	return values
}

func (m *Map) Clear() {
	m.dataStore = make(map[interface{}]interface{})
}
