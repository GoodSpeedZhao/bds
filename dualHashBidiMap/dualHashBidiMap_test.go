package dualHashBidiMap

import (
	"reflect"
	"testing"
)

func TestNewDualHashBidiMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestNewDualHashBidiMap case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDualHashBidiMap()
			if got == nil {
				t.Errorf("got nil, want DualHashBidiMap")
				return
			}
			if got.keyValue == nil {
				t.Errorf("got nil, want DualHashBidiMap keyValue Map.")
			}
			if got.valueKey == nil {
				t.Errorf("got nil, want DualHashBidiMap valueKey Map.")
			}
		})
	}
}

func TestMap_Put(t *testing.T) {
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args []args
		size int
	}{
		{"TestMap_Put case1", []args{{1, 10}}, 1},
		{"TestMap_Put case2", []args{{1, 10}, {10, 100}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDualHashBidiMap()
			for i := 0; i < len(tt.args); i++ {
				got.Put(tt.args[i].key, tt.args[i].value)
			}
			if size := got.Size(); size != tt.size {
				t.Errorf("DualHashBidiMap.Size() = %v, want %v", size, tt.size)
			}
		})
	}
}

func TestMap_GetValue(t *testing.T) {
	type puts struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name    string
		put     []puts
		key     interface{}
		wantVal interface{}
		wantOk  bool
	}{
		{"TestMap_GetValue case1", []puts{{1, 10}, {10, 100}}, 1, 10, true},
		{"TestMap_GetValue case2", []puts{{1, 10}, {10, 100}}, 10, 100, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDualHashBidiMap()
			for i := 0; i < len(tt.put); i++ {
				got.Put(tt.put[i].key, tt.put[i].value)
			}
			value, ok := got.GetValue(tt.key)
			if !reflect.DeepEqual(value, tt.wantVal) {
				t.Errorf("Map.GetValue() gotVal = %v, want %v", value, tt.wantVal)
			}
			if ok != tt.wantOk {
				t.Errorf("Map.GetValue() gotOk = %v, want %v", ok, tt.wantOk)
			}
		})
	}
}

func TestMap_GetKey(t *testing.T) {
	type puts struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name    string
		put     []puts
		value   interface{}
		wantKey interface{}
		wantOk  bool
	}{
		{"TestMap_GetKey case1", []puts{{1, 10}, {10, 100}}, 10, 1, true},
		{"TestMap_GetKey case2", []puts{{1, 10}, {10, 100}}, 100, 10, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDualHashBidiMap()
			for i := 0; i < len(tt.put); i++ {
				got.Put(tt.put[i].key, tt.put[i].value)
			}
			key, ok := got.GetKey(tt.value)
			if !reflect.DeepEqual(key, tt.wantKey) {
				t.Errorf("Map.GetKey() gotKey = %v, want %v", key, tt.wantKey)
			}
			if ok != tt.wantOk {
				t.Errorf("Map.GetKey() gotOk = %v, want %v", ok, tt.wantOk)
			}
		})
	}
}

func TestMap_Delete(t *testing.T) {
	type puts struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		put  []puts
		del  []interface{}
		size int
	}{
		{"TestMap_Delete case1", []puts{{1, 10}, {10, 100}}, []interface{}{1}, 1},
		{"TestMap_Delete case2", []puts{{1, 10}, {10, 100}}, []interface{}{1, 10}, 0},
		{"TestMap_Delete case3", []puts{}, []interface{}{1, 10}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDualHashBidiMap()
			for i := 0; i < len(tt.put); i++ {
				got.Put(tt.put[i].key, tt.put[i].value)
			}
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			if size := got.Size(); size != tt.size {
				t.Errorf("DualHashBidiMap.Size() = %v, want %v", size, tt.size)
			}
		})
	}
}

func TestMap_Size(t *testing.T) {
	type puts struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		put  []puts
		del  []interface{}
		size int
	}{
		{"TestMap_Size case1", []puts{{1, 10}, {10, 100}}, []interface{}{1}, 1},
		{"TestMap_Size case2", []puts{{1, 10}, {10, 100}}, []interface{}{1, 10}, 0},
		{"TestMap_Size case3", []puts{}, []interface{}{1, 10}, 0},
		{"TestMap_Size case4", []puts{}, []interface{}{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDualHashBidiMap()
			for i := 0; i < len(tt.put); i++ {
				got.Put(tt.put[i].key, tt.put[i].value)
			}
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			if size := got.Size(); size != tt.size {
				t.Errorf("Map.Size() = %v, want %v", size, tt.size)
			}
		})
	}
}

func TestMap_IsEmpty(t *testing.T) {
	type puts struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		put  []puts
		del  []interface{}
		want bool
	}{
		{"TestMap_IsEmpty case1", []puts{}, []interface{}{}, true},
		{"TestMap_IsEmpty case2", []puts{{1, 10}}, []interface{}{}, false},
		{"TestMap_IsEmpty case3", []puts{{1, 10}}, []interface{}{12}, false},
		{"TestMap_IsEmpty case4", []puts{{1, 10}}, []interface{}{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDualHashBidiMap()
			for i := 0; i < len(tt.put); i++ {
				got.Put(tt.put[i].key, tt.put[i].value)
			}
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			if isEmpty := got.IsEmpty(); isEmpty != tt.want {
				t.Errorf("Map.IsEmpty() = %v, want %v", isEmpty, tt.want)
			}
		})
	}
}

func TestMap_Keys(t *testing.T) {
	type puts struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		put  []puts
		del  []interface{}
		want []interface{}
	}{
		{"TestMap_Keys case1", []puts{}, []interface{}{}, []interface{}{}},
		{"TestMap_Keys case2", []puts{{1, 10}}, []interface{}{}, []interface{}{1}},
		{"TestMap_Keys case3", []puts{{1, 10}, {1, 10}}, []interface{}{}, []interface{}{1}},
		{"TestMap_Keys case4", []puts{{1, 10}, {100, 50}, {10, 100}}, []interface{}{}, []interface{}{1, 10, 100}},
		{"TestMap_Keys case5", []puts{{1, 10}, {100, 50}, {10, 100}}, []interface{}{100}, []interface{}{1, 10}},
		{"TestMap_Keys case6", []puts{{1, 10}, {100, 50}, {10, 100}}, []interface{}{10, 100}, []interface{}{1}},
		{"TestMap_Keys case7", []puts{{1, 10}, {100, 50}, {10, 100}}, []interface{}{1, 10, 100}, []interface{}{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDualHashBidiMap()
			for i := 0; i < len(tt.put); i++ {
				got.Put(tt.put[i].key, tt.put[i].value)
			}
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			keys := got.Keys()
			if len(keys) != len(tt.want) {
				t.Errorf("Map.keys() len = %v, want len %v", len(keys), len(tt.want))
				return
			}
			for i := 0; i < len(tt.want); i++ {
				flag := true
				for j := 0; j < len(keys); j++ {
					if tt.want[i] == keys[j] {
						flag = false
						break
					}
				}
				if flag {
					t.Errorf("Map.keys() not equal")
					return
				}
			}
		})
	}
}

func TestMap_Values(t *testing.T) {
	type puts struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		put  []puts
		del  []interface{}
		want []interface{}
	}{
		{"TestMap_Values case1", []puts{}, []interface{}{}, []interface{}{}},
		{"TestMap_Values case2", []puts{{1, 10}}, []interface{}{}, []interface{}{10}},
		{"TestMap_Values case3", []puts{{1, 10}, {1, 10}}, []interface{}{}, []interface{}{10}},
		{"TestMap_Values case4", []puts{{1, 10}, {10, 100}, {100, 50}}, []interface{}{}, []interface{}{10, 100, 50}},
		{"TestMap_Values case5", []puts{{1, 10}, {10, 100}, {100, 50}}, []interface{}{100}, []interface{}{100, 10}},
		{"TestMap_Values case6", []puts{{1, 10}, {10, 100}, {100, 50}}, []interface{}{10, 100}, []interface{}{10}},
		{"TestMap_Values case7", []puts{{1, 10}, {10, 100}, {100, 50}}, []interface{}{1, 10, 100}, []interface{}{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDualHashBidiMap()
			for i := 0; i < len(tt.put); i++ {
				got.Put(tt.put[i].key, tt.put[i].value)
			}
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			values := got.Values()
			if len(values) != len(tt.want) {
				t.Errorf("Map.Values() len = %v, want len %v", len(values), len(tt.want))
				return
			}
			for i := 0; i < len(tt.want); i++ {
				flag := true
				for j := 0; j < len(values); j++ {
					if tt.want[i] == values[j] {
						flag = false
						break
					}
				}
				if flag {
					t.Errorf("Map.Values() not equal")
					return
				}
			}
		})
	}
}

func TestMap_Clear(t *testing.T) {
	type puts struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		put  []puts
		del  []interface{}
		size int
	}{
		{"TestMap_Clear case1", []puts{}, []interface{}{}, 0},
		{"TestMap_Clear case2", []puts{{1, 10}}, []interface{}{}, 0},
		{"TestMap_Clear case3", []puts{{1, 10}}, []interface{}{1}, 0},
		{"TestMap_Clear case4", []puts{}, []interface{}{1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDualHashBidiMap()
			for i := 0; i < len(tt.put); i++ {
				got.Put(tt.put[i].key, tt.put[i].value)
			}
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			got.Clear()
			if size := got.Size(); size != tt.size {
				t.Errorf("Map.Size() = %v, want %v", size, tt.size)
			}
		})
	}
}
