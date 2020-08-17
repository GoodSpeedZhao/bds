package hashMap

import (
	"reflect"
	"testing"
)

func TestNewHashMap(t *testing.T) {
	tests := []struct {
		name   string
		keys   []interface{}
		values []interface{}
	}{
		{"TestNewHashMap case1", []interface{}{}, []interface{}{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewHashMap()
			if this == nil {
				t.Errorf("got nil, want HashMap")
				return
			}
			if this.dataStore == nil {
				t.Errorf("got nil, want HashMap dataStore")
			}
		})
	}
}

func TestMap_Put(t *testing.T) {
	tests := []struct {
		name   string
		keys   []interface{}
		values []interface{}
		size   int
	}{
		{"TestMap_Put case1", []interface{}{}, []interface{}{}, 0},
		{"TestMap_Put case2", []interface{}{1}, []interface{}{1}, 1},
		{"TestMap_Put case3", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, 3},
		{"TestMap_Put case4", []interface{}{1, 6, 6}, []interface{}{1, 6, 6}, 2},
		{"TestMap_Put case5", []interface{}{1, 6, 6}, []interface{}{1, 6, 12}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewHashMap()
			for i := 0; i < len(tt.keys); i++ {
				this.Put(tt.keys[i], tt.values[i])
			}
			if size := this.Size(); size != tt.size {
				t.Errorf("HashMap.Size() = %v, want %v", size, tt.size)
			}
		})
	}
}

func TestMap_Get(t *testing.T) {
	type args struct {
		key interface{}
	}
	tests := []struct {
		name      string
		putKeys   []interface{}
		putValues []interface{}
		args      args
		wantVal   interface{}
		wantOk    bool
	}{
		{"TestMap_Get case1", []interface{}{}, []interface{}{}, args{1}, nil, false},
		{"TestMap_Get case2", []interface{}{1}, []interface{}{1}, args{1}, 1, true},
		{"TestMap_Get case3", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, args{9}, 9, true},
		{"TestMap_Get case4", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, args{12}, nil, false},
		{"TestMap_Get case5", []interface{}{1, 6, 6}, []interface{}{1, 6, 12}, args{6}, 12, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewHashMap()
			for i := 0; i < len(tt.putKeys); i++ {
				this.Put(tt.putKeys[i], tt.putValues[i])
			}
			gotVal, gotOk := this.Get(tt.args.key)
			if gotOk != tt.wantOk {
				t.Errorf("HashMap.Get() gotOk = %v, want %v", gotOk, tt.wantOk)
				return
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("HashMap.Get() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func TestMap_Delete(t *testing.T) {
	type args struct {
		key interface{}
	}
	tests := []struct {
		name      string
		putKeys   []interface{}
		putValues []interface{}
		args      args
		size      int
		get       interface{}
		wantVal   interface{}
		wantOk    bool
	}{
		{"TestMap_Delete case1", []interface{}{}, []interface{}{}, args{1}, 0, nil, nil, false},
		{"TestMap_Delete case2", []interface{}{1}, []interface{}{1}, args{12}, 1, nil, nil, false},
		{"TestMap_Delete case3", []interface{}{1}, []interface{}{1}, args{1}, 0, nil, nil, false},
		{"TestMap_Delete case4", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, args{12}, 3, nil, nil, false},
		{"TestMap_Delete case5", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, args{6}, 2, 6, nil, false},
		{"TestMap_Delete case6", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, args{6}, 2, 9, 9, true},
		{"TestMap_Delete case7", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, args{9}, 2, nil, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewHashMap()
			for i := 0; i < len(tt.putKeys); i++ {
				this.Put(tt.putKeys[i], tt.putValues[i])
			}
			this.Delete(tt.args.key)
			if size := this.Size(); size != tt.size {
				t.Errorf("HashMap.Size() = %v, want %v", size, tt.size)
			}
			if tt.get != nil {
				gotVal, gotOk := this.Get(tt.get)
				if gotOk != tt.wantOk {
					t.Errorf("HashMap.Get() gotOk = %v, want %v", gotOk, tt.wantOk)
					return
				}
				if !reflect.DeepEqual(gotVal, tt.wantVal) {
					t.Errorf("HashMap.Get() gotVal = %v, want %v", gotVal, tt.wantVal)
				}
			}
		})
	}
}

func TestMap_Size(t *testing.T) {
	tests := []struct {
		name      string
		putKeys   []interface{}
		putValues []interface{}
		del       []interface{}
		want      int
	}{
		{"TestMap_Size case1", []interface{}{}, []interface{}{}, []interface{}{}, 0},
		{"TestMap_Size case2", []interface{}{1}, []interface{}{1}, []interface{}{}, 1},
		{"TestMap_Size case3", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{}, 3},
		{"TestMap_Size case4", []interface{}{}, []interface{}{}, []interface{}{1}, 0},
		{"TestMap_Size case5", []interface{}{1}, []interface{}{1}, []interface{}{1}, 0},
		{"TestMap_Size case6", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, 0},
		{"TestMap_Size case7", []interface{}{1, 1, 1}, []interface{}{1, 6, 9}, []interface{}{}, 1},
		{"TestMap_Size case8", []interface{}{1, 1, 1}, []interface{}{1, 6, 9}, []interface{}{1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewHashMap()
			for i := 0; i < len(tt.putKeys); i++ {
				this.Put(tt.putKeys[i], tt.putValues[i])
			}
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if got := this.Size(); got != tt.want {
				t.Errorf("HashMap.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_IsEmpty(t *testing.T) {
	tests := []struct {
		name      string
		putKeys   []interface{}
		putValues []interface{}
		del       []interface{}
		want      bool
	}{
		{"TestMap_IsEmpty case1", []interface{}{}, []interface{}{}, []interface{}{}, true},
		{"TestMap_IsEmpty case2", []interface{}{1}, []interface{}{1}, []interface{}{}, false},
		{"TestMap_IsEmpty case3", []interface{}{1, 1}, []interface{}{1, 6}, []interface{}{}, false},
		{"TestMap_IsEmpty case4", []interface{}{1, 1}, []interface{}{1, 6}, []interface{}{1}, true},
		{"TestMap_IsEmpty case5", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewHashMap()
			for i := 0; i < len(tt.putKeys); i++ {
				this.Put(tt.putKeys[i], tt.putValues[i])
			}
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if got := this.IsEmpty(); got != tt.want {
				t.Errorf("HashMap.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_Keys(t *testing.T) {
	tests := []struct {
		name      string
		putKeys   []interface{}
		putValues []interface{}
		del       []interface{}
		size      int
		want      []interface{}
	}{
		{"TestMap_Keys case1", []interface{}{}, []interface{}{}, []interface{}{}, 0, []interface{}{}},
		{"TestMap_Keys case2", []interface{}{1}, []interface{}{1}, []interface{}{}, 1, []interface{}{1}},
		{"TestMap_Keys case3", []interface{}{1, 1, 1}, []interface{}{1, 6, 9}, []interface{}{}, 1, []interface{}{1}},
		{"TestMap_Keys case4", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{}, 3, []interface{}{1, 6, 9}},
		{"TestMap_Keys case5", []interface{}{1, 1, 1}, []interface{}{1, 6, 9}, []interface{}{1}, 0, []interface{}{}},
		{"TestMap_Keys case6", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, 0, []interface{}{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewHashMap()
			for i := 0; i < len(tt.putKeys); i++ {
				this.Put(tt.putKeys[i], tt.putValues[i])
			}
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			got := this.Keys()
			if len(got) != tt.size {
				t.Errorf("HashMap.Size() = %v, want %v", got, tt.want)
			}
			for _, val := range got {
				found := false
				for _, inner := range tt.want {
					if inner == val {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("HashMap %v not in the want.", val)
				}
			}
		})
	}
}

func TestMap_Values(t *testing.T) {
	tests := []struct {
		name      string
		putKeys   []interface{}
		putValues []interface{}
		del       []interface{}
		size      int
		want      []interface{}
	}{
		{"TestMap_Values case1", []interface{}{}, []interface{}{}, []interface{}{}, 0, []interface{}{}},
		{"TestMap_Values case2", []interface{}{1}, []interface{}{1}, []interface{}{}, 1, []interface{}{1}},
		{"TestMap_Values case3", []interface{}{1, 1, 1}, []interface{}{1, 6, 9}, []interface{}{}, 1, []interface{}{9}},
		{"TestMap_Values case4", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{}, 3, []interface{}{1, 6, 9}},
		{"TestMap_Values case5", []interface{}{1, 1, 1}, []interface{}{1, 6, 9}, []interface{}{1}, 0, []interface{}{}},
		{"TestMap_Values case6", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, 0, []interface{}{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewHashMap()
			for i := 0; i < len(tt.putKeys); i++ {
				this.Put(tt.putKeys[i], tt.putValues[i])
			}
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			got := this.Values()
			if len(got) != tt.size {
				t.Errorf("HashMap.Size() = %v, want %v", got, tt.want)
			}
			for _, val := range got {
				found := false
				for _, inner := range tt.want {
					if inner == val {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("HashMap %v not in the want.", val)
				}
			}
		})
	}
}

func TestMap_Clear(t *testing.T) {
	tests := []struct {
		name      string
		putKeys   []interface{}
		putValues []interface{}
		del       []interface{}
		size      int
		want      int
	}{
		{"TestMap_Clear case1", []interface{}{}, []interface{}{}, []interface{}{}, 0, 0},
		{"TestMap_Clear case2", []interface{}{1}, []interface{}{1}, []interface{}{}, 0, 0},
		{"TestMap_Clear case3", []interface{}{1, 1, 1}, []interface{}{1, 6, 9}, []interface{}{}, 0, 0},
		{"TestMap_Clear case4", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{}, 0, 0},
		{"TestMap_Clear case5", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, 0, 0},
		{"TestMap_Clear case6", []interface{}{1, 1, 1}, []interface{}{1, 6, 9}, []interface{}{1}, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewHashMap()
			for i := 0; i < len(tt.putKeys); i++ {
				this.Put(tt.putKeys[i], tt.putValues[i])
			}
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			this.Clear()
			if got := this.Size(); got != tt.want {
				t.Errorf("HashMap.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
