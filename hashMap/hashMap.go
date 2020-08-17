package hashMap

type Map struct {
	dataStore map[interface{}]interface{}
}

func NewHashMap() *Map {
	return &Map{
		dataStore: make(map[interface{}]interface{}),
	}
}

func (this *Map) Put(key interface{}, value interface{}) {
	this.dataStore[key] = value
}

func (this *Map) Get(key interface{}) (val interface{}, ok bool) {
	val, ok = this.dataStore[key]
	return
}

func (this *Map) Delete(key interface{}) {
	delete(this.dataStore, key)
}

func (this *Map) Size() int {
	return len(this.dataStore)
}

func (this *Map) IsEmpty() bool {
	return this.Size() == 0
}

func (this *Map) Keys() []interface{} {
	keys := make([]interface{}, this.Size())
	count := 0
	for key := range this.dataStore {
		keys[count] = key
		count++
	}
	return keys
}

func (this *Map) Values() []interface{} {
	values := make([]interface{}, this.Size())
	count := 0
	for _, val := range this.dataStore {
		values[count] = val
		count++
	}
	return values
}

func (this *Map) Clear() {
	this.dataStore = make(map[interface{}]interface{})
}
