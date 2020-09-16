package dualHashBidiMap

import "github.com/GoodSpeedZhao/bds/hashMap"

type Map struct {
	keyValue *hashMap.Map
	valueKey *hashMap.Map
}

func NewDualHashBidiMap() *Map {
	return &Map{
		keyValue: hashMap.NewHashMap(),
		valueKey: hashMap.NewHashMap(),
	}
}

func (this *Map) Put(key interface{}, value interface{}) {
	this.keyValue.Put(key, value)
	this.valueKey.Put(value, key)
}

func (this *Map) GetValue(key interface{}) (val interface{}, ok bool) {
	val, ok = this.keyValue.Get(key)
	return
}

func (this *Map) GetKey(value interface{}) (key interface{}, ok bool) {
	key, ok = this.valueKey.Get(value)
	return
}

func (this *Map) Delete(key interface{}) {
	if value, ok := this.keyValue.Get(key); ok {
		this.keyValue.Delete(key)
		this.valueKey.Delete(value)
	}
}

func (this *Map) Size() int {
	return this.keyValue.Size()
}

func (this *Map) IsEmpty() bool {
	return this.keyValue.Size() == 0
}

func (this *Map) Keys() []interface{} {
	return this.keyValue.Keys()
}

func (this *Map) Values() []interface{} {
	return this.keyValue.Values()
}

func (this *Map) Clear() {
	this.keyValue.Clear()
	this.valueKey.Clear()
}
