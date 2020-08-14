package arrayList

import (
	"reflect"
	"testing"
)

func TestNewArrayList(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want []interface{}
		size int
	}{
		{"TestNewArrayList case1", []interface{}{1, 2, 3}, []interface{}{1, 2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList(tt.args...)
			if this == nil {
				t.Errorf("got nil, want arrayList")
				return
			}
			if !reflect.DeepEqual(this.dataStore, tt.want) {
				t.Errorf("NewArrayList() = %v, want %v", this, tt.want)
			}
			if !reflect.DeepEqual(this.size, tt.size) {
				t.Errorf("NewArrayList() = %v, size %v", this.size, tt.size)
			}
		})
	}
}

func TestArrayList_Get(t *testing.T) {
	tests := []struct {
		name    string
		append  []interface{}
		del     []int
		get     int
		want    interface{}
		wantErr bool
	}{
		{"TestArrayList_Get case1", []interface{}{}, []int{}, 1, nil, false},
		{"TestArrayList_Get case2", []interface{}{1, 2, 6, 8}, []int{}, 0, 1, true},
		{"TestArrayList_Get case3", []interface{}{1, 2, 6, 8}, []int{}, 3, 8, true},
		{"TestArrayList_Get case4", []interface{}{1, 2, 6, 8, 9}, []int{}, 2, 6, true},
		{"TestArrayList_Get case5", []interface{}{1, 2, 6, 8, 9}, []int{0}, 0, 2, true},
		{"TestArrayList_Get case6", []interface{}{1, 2, 6, 8, 9}, []int{4}, 3, 8, true},
		{"TestArrayList_Get case7", []interface{}{1, 2, 6, 8, 9}, []int{2}, 2, 8, true},
		{"TestArrayList_Get case8", []interface{}{1, 6, 8}, []int{0, 0, 0}, 0, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			got, err := this.Get(tt.get)
			if (err != nil) == tt.wantErr {
				t.Errorf("ArrayList.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Set(t *testing.T) {
	type gots struct {
		index  int
		newVal interface{}
	}
	type args struct {
		index  int
		newVal interface{}
	}
	tests := []struct {
		name    string
		append  []interface{}
		del     []int
		get     gots
		args    args
		wantErr bool
	}{
		{"TestArrayList_Set case1", []interface{}{}, []int{}, gots{-1, nil}, args{1, 1}, false},
		{"TestArrayList_Set case2", []interface{}{0}, []int{}, gots{-1, nil}, args{1, 1}, false},
		{"TestArrayList_Set case3", []interface{}{0, 9}, []int{}, gots{1, 1}, args{1, 1}, true},
		{"TestArrayList_Set case4", []interface{}{0, 9, 16, 20}, []int{}, gots{0, 1}, args{0, 1}, true},
		{"TestArrayList_Set case5", []interface{}{0, 9, 16, 20}, []int{}, gots{-1, nil}, args{4, 1}, false},
		{"TestArrayList_Set case6", []interface{}{0, 9, 16, 20}, []int{}, gots{3, 1}, args{3, 1}, true},
		{"TestArrayList_Set case7", []interface{}{0, 9, 16, 20}, []int{0}, gots{0, 1}, args{0, 1}, true},
		{"TestArrayList_Set case8", []interface{}{0, 9, 16, 20}, []int{3}, gots{2, 1}, args{2, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if err := this.Set(tt.args.index, tt.args.newVal); (err != nil) == tt.wantErr {
				t.Errorf("ArrayList.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, err := this.Get(tt.args.index)
			if tt.get.index != -1 {
				if (err != nil) == tt.wantErr {
					t.Errorf("ArrayList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if (got != nil) && !reflect.DeepEqual(got, tt.args.newVal) {
					t.Errorf("ArrayList.Get() = %v, want %v", got, tt.args.newVal)
				}
			}
		})
	}
}

func TestArrayList_Insert(t *testing.T) {
	type args struct {
		index  int
		newVal []interface{}
	}
	tests := []struct {
		name    string
		append  []interface{}
		del     []int
		get     args
		args    args
		wantErr bool
	}{
		{"TestArrayList_Insert case1", []interface{}{}, []int{}, args{-1, nil}, args{0, []interface{}{1}}, false},
		{"TestArrayList_Insert case2", []interface{}{}, []int{}, args{-1, nil}, args{1, []interface{}{1}}, false},
		{"TestArrayList_Insert case3", []interface{}{}, []int{}, args{-1, nil}, args{0, []interface{}{}}, false},
		{"TestArrayList_Insert case4", []interface{}{}, []int{}, args{-1, nil}, args{1, []interface{}{}}, false},
		{"TestArrayList_Insert case5", []interface{}{1}, []int{}, args{-1, nil}, args{0, []interface{}{}}, true},
		{"TestArrayList_Insert case6", []interface{}{1, 2}, []int{}, args{1, []interface{}{12}}, args{0, []interface{}{12}}, true},
		{"TestArrayList_Insert case7", []interface{}{1, 2, 3}, []int{}, args{3, []interface{}{12}}, args{2, []interface{}{12}}, true},
		{"TestArrayList_Insert case8", []interface{}{1, 2, 3}, []int{2}, args{2, []interface{}{12}}, args{1, []interface{}{12}}, true},
		{"TestArrayList_Insert case9", []interface{}{1, 2, 3}, []int{0}, args{2, []interface{}{12}}, args{1, []interface{}{12}}, true},
		{"TestArrayList_Insert case10", []interface{}{1, 2, 3}, []int{0}, args{2, []interface{}{12}}, args{1, []interface{}{12}}, true},
		{"TestArrayList_Insert case11", []interface{}{1, 2, 3}, []int{0, 0, 0}, args{-1, nil}, args{0, []interface{}{12}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if err := this.Insert(tt.args.index, tt.args.newVal...); (err != nil) == tt.wantErr {
				t.Errorf("ArrayList.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.get.index != -1 {
				got, err := this.Get(tt.get.index)
				if (err != nil) == tt.wantErr {
					t.Errorf("ArrayList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.get.newVal[0]) {
					t.Errorf("ArrayList.Get() = %v, want %v", got, tt.get.newVal[0])
				}
			}
		})
	}
}

func TestArrayList_Delete(t *testing.T) {
	type gots struct {
		index  int
		newVal interface{}
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		append  []interface{}
		get     gots
		args    args
		wantErr bool
	}{
		{"TestArrayList_Delete case1", []interface{}{}, gots{-1, nil}, args{0}, false},
		{"TestArrayList_Delete case2", []interface{}{1}, gots{-1, nil}, args{2}, false},
		{"TestArrayList_Delete case3", []interface{}{1, 3, 6, 9}, gots{0, 3}, args{0}, true},
		{"TestArrayList_Delete case4", []interface{}{1, 3, 6, 9}, gots{2, 6}, args{3}, true},
		{"TestArrayList_Delete case5", []interface{}{1, 3, 6, 9, 12}, gots{2, 9}, args{2}, true},
		{"TestArrayList_Delete case6", []interface{}{1}, gots{-1, nil}, args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.append...)
			if err := this.Delete(tt.args.index); (err != nil) == tt.wantErr {
				t.Errorf("ArrayList.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.get.index != -1 {
				got, err := this.Get(tt.get.index)
				if (err != nil) == tt.wantErr {
					t.Errorf("ArrayList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.get.newVal) {
					t.Errorf("ArrayList.Get() = %v, want %v", got, tt.get.newVal)
				}
			}
		})
	}
}

func TestArrayList_Append(t *testing.T) {
	type gots struct {
		index  int
		newVal interface{}
	}
	type args struct {
		newVal []interface{}
	}
	tests := []struct {
		name    string
		del     []int
		get     gots
		args    args
		wantErr bool
	}{
		{"TestArrayList_Append case1", []int{}, gots{0, 1}, args{[]interface{}{1}}, true},
		{"TestArrayList_Append case2", []int{}, gots{2, 3}, args{[]interface{}{1, 2, 3}}, true},
		{"TestArrayList_Append case3", []int{0}, gots{0, 2}, args{[]interface{}{1, 2, 3}}, true},
		{"TestArrayList_Append case4", []int{2}, gots{0, 1}, args{[]interface{}{1, 2, 3}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.args.newVal...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if tt.get.index != -1 {
				got, err := this.Get(tt.get.index)
				if (err != nil) == tt.wantErr {
					t.Errorf("ArrayList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.get.newVal) {
					t.Errorf("ArrayList.Get() = %v, want %v", got, tt.get.newVal)
				}
			}
		})
	}
}

func TestArrayList_Prepend(t *testing.T) {
	type gots struct {
		index  int
		newVal interface{}
	}
	type args struct {
		newVal []interface{}
	}
	tests := []struct {
		name    string
		del     []int
		get     gots
		args    args
		wantErr bool
	}{
		{"TestArrayList_Prepend case1", []int{}, gots{0, 1}, args{[]interface{}{1}}, true},
		{"TestArrayList_Prepend case2", []int{}, gots{0, 6}, args{[]interface{}{1, 3, 6}}, true},
		{"TestArrayList_Prepend case3", []int{0}, gots{0, 3}, args{[]interface{}{1, 3, 6}}, true},
		{"TestArrayList_Prepend case4", []int{2}, gots{0, 6}, args{[]interface{}{1, 3, 6}}, true},
		{"TestArrayList_Prepend case5", []int{0}, gots{0, 3}, args{[]interface{}{1, 3, 6}}, true},
		{"TestArrayList_Prepend case6", []int{1}, gots{0, 6}, args{[]interface{}{1, 3, 6}}, true},
		{"TestArrayList_Prepend case7", []int{2}, gots{0, 6}, args{[]interface{}{1, 3, 6}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Prepend(tt.args.newVal...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if tt.get.index != -1 {
				got, err := this.Get(tt.get.index)
				if (err != nil) == tt.wantErr {
					t.Errorf("ArrayList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.get.newVal) {
					t.Errorf("ArrayList.Get() = %v, want %v", got, tt.get.newVal)
				}
			}
		})
	}
}

func TestArrayList_Contains(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		args   args
		want   bool
	}{
		{"TestArrayList_Contains case1", []interface{}{}, []int{}, args{[]interface{}{1}}, false},
		{"TestArrayList_Contains case2", []interface{}{3, 6, 9}, []int{}, args{[]interface{}{1}}, false},
		{"TestArrayList_Contains case3", []interface{}{3, 6, 9}, []int{}, args{[]interface{}{1, 2}}, false},
		{"TestArrayList_Contains case4", []interface{}{3, 6, 9}, []int{}, args{[]interface{}{3}}, true},
		{"TestArrayList_Contains case5", []interface{}{3, 6, 9}, []int{}, args{[]interface{}{9}}, true},
		{"TestArrayList_Contains case6", []interface{}{3, 6, 9}, []int{}, args{[]interface{}{3, 6, 9}}, true},
		{"TestArrayList_Contains case7", []interface{}{3, 6, 9}, []int{0}, args{[]interface{}{3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if got := this.Contains(tt.args.values...); got != tt.want {
				t.Errorf("ArrayList.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Swap(t *testing.T) {
	type gots struct {
		index  int
		newVal interface{}
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		append  []interface{}
		del     []int
		get     gots
		args    args
		wantErr bool
	}{
		{"TestArrayList_Swap case1", []interface{}{}, []int{}, gots{-1, nil}, args{0, 1}, false},
		{"TestArrayList_Swap case2", []interface{}{1, 3, 6, 9}, []int{}, gots{-1, nil}, args{0, 0}, true},
		{"TestArrayList_Swap case3", []interface{}{1, 3, 6, 9}, []int{}, gots{0, 3}, args{0, 1}, true},
		{"TestArrayList_Swap case4", []interface{}{1, 3, 6, 9}, []int{}, gots{0, 3}, args{0, 1}, true},
		{"TestArrayList_Swap case5", []interface{}{1, 3, 6, 9}, []int{0}, gots{0, 6}, args{0, 1}, true},
		{"TestArrayList_Swap case6", []interface{}{1, 3, 6, 9}, []int{3}, gots{1, 6}, args{2, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if err := this.Swap(tt.args.i, tt.args.j); (err != nil) == tt.wantErr {
				t.Errorf("ArrayList.Swap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.get.index != -1 {
				got, err := this.Get(tt.get.index)
				if (err != nil) == tt.wantErr {
					t.Errorf("ArrayList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.get.newVal) {
					t.Errorf("ArrayList.Get() = %v, want %v", got, tt.get.newVal)
				}
			}
		})
	}
}

func TestArrayList_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		want   bool
	}{
		{"TestArrayList_IsEmpty case1", []interface{}{}, []int{}, true},
		{"TestArrayList_IsEmpty case2", []interface{}{1}, []int{}, false},
		{"TestArrayList_IsEmpty case3", []interface{}{1}, []int{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if got := this.IsEmpty(); got != tt.want {
				t.Errorf("ArrayList.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Size(t *testing.T) {
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		want   int
	}{
		{"TestArrayList_Size case1", []interface{}{}, []int{}, 0},
		{"TestArrayList_Size case2", []interface{}{1, 3, 6}, []int{}, 3},
		{"TestArrayList_Size case3", []interface{}{1, 3, 6}, []int{0}, 2},
		{"TestArrayList_Size case4", []interface{}{1, 3, 6}, []int{2}, 2},
		{"TestArrayList_Size case5", []interface{}{1, 3, 6}, []int{0, 0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if got := this.Size(); got != tt.want {
				t.Errorf("ArrayList.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_withinRange(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		args   args
		want   bool
	}{
		{"TestArrayList_withinRange case1", []interface{}{}, []int{}, args{0}, false},
		{"TestArrayList_withinRange case2", []interface{}{1, 6, 9}, []int{}, args{0}, true},
		{"TestArrayList_withinRange case3", []interface{}{1, 6, 9}, []int{}, args{2}, true},
		{"TestArrayList_withinRange case4", []interface{}{1, 6, 9}, []int{2}, args{2}, false},
		{"TestArrayList_withinRange case5", []interface{}{1, 6, 9}, []int{0}, args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if got := this.withinRange(tt.args.index); got != tt.want {
				t.Errorf("ArrayList.withinRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Clear(t *testing.T) {
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		want   int
	}{
		{"TestArrayList_Clear case1", []interface{}{}, []int{}, 0},
		{"TestArrayList_Clear case2", []interface{}{1, 2, 3}, []int{}, 0},
		{"TestArrayList_Clear case3", []interface{}{1, 2, 3}, []int{0, 0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			this.Clear()
			if got := this.Size(); got != tt.want {
				t.Errorf("ArrayList.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Resize(t *testing.T) {
	type gots struct {
		index  int
		newVal interface{}
	}
	type args struct {
		cap int
	}
	tests := []struct {
		name    string
		append  []interface{}
		del     []int
		get     gots
		args    args
		wantErr bool
	}{
		{"TestArrayList_Resize case1", []interface{}{}, []int{}, gots{-1, nil}, args{-1}, false},
		{"TestArrayList_Resize case2", []interface{}{1, 3, 6, 9}, []int{}, gots{1, 3}, args{2}, true},
		{"TestArrayList_Resize case3", []interface{}{1, 3, 6, 9}, []int{}, gots{1, 3}, args{2}, true},
		{"TestArrayList_Resize case4", []interface{}{1, 3, 6, 9}, []int{}, gots{0, 1}, args{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewArrayList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if err := this.Resize(tt.args.cap); (err != nil) == tt.wantErr {
				t.Errorf("ArrayList.Resize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.get.index != -1 {
				got, err := this.Get(tt.get.index)
				if (err != nil) == tt.wantErr {
					t.Errorf("ArrayList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.get.newVal) {
					t.Errorf("ArrayList.Get() = %v, want %v", got, tt.get.newVal)
				}
			}
		})
	}
}
