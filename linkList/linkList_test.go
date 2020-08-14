package linkList

import (
	"reflect"
	"testing"
)

func TestNewLinkList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestNewLinkList case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			if this == nil {
				t.Errorf("got nil, want linkList")
				return
			}
			if !reflect.DeepEqual(this.head, (*Node)(nil)) {
				t.Errorf("NewLinkList() head = %v, want %v", this.head, nil)
			}
			if !reflect.DeepEqual(this.tail, (*Node)(nil)) {
				t.Errorf("NewLinkList() tail = %v, want %v", this.tail, nil)
			}
			if !reflect.DeepEqual(this.size, 0) {
				t.Errorf("NewLinkList() size = %v, size %v", this.size, 0)
			}
		})
	}
}

func TestNewNode(t *testing.T) {
	type args struct {
		newVal interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"TestNewNode case1", args{1}, interface{}(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewNode(tt.args.newVal, nil)
			if this == nil {
				t.Errorf("got nil, want Node")
				return
			}
			if !reflect.DeepEqual(this.Val, tt.want) {
				t.Errorf("NewNode() Val = %v, want %v", this.Val, tt.want)
			}
			if !reflect.DeepEqual(this.Next, (*Node)(nil)) {
				t.Errorf("NewNode() Next = %v, want %v", this.Next, nil)
			}
		})
	}
}

func TestLinkList_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		append  []interface{}
		del     []int
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestLinkList_Get case1", []interface{}{}, []int{}, args{1}, nil, false},
		{"TestLinkList_Get case2", []interface{}{1, 2, 6, 8}, []int{}, args{0}, 1, true},
		{"TestLinkList_Get case3", []interface{}{1, 2, 6, 8}, []int{}, args{3}, 8, true},
		{"TestLinkList_Get case4", []interface{}{1, 2, 6, 8, 9}, []int{}, args{2}, 6, true},
		{"TestLinkList_Get case5", []interface{}{1, 2, 6, 8, 9}, []int{0}, args{0}, 2, true},
		{"TestLinkList_Get case6", []interface{}{1, 2, 6, 8, 9}, []int{4}, args{3}, 8, true},
		{"TestLinkList_Get case7", []interface{}{1, 2, 6, 8, 9}, []int{2}, args{2}, 8, true},
		{"TestLinkList_Get case8", []interface{}{1, 6, 8}, []int{0, 0, 0}, args{0}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			got, err := this.Get(tt.args.index)
			if (err != nil) == tt.wantErr {
				t.Errorf("LinkList.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_Set(t *testing.T) {
	type args struct {
		index  int
		newVal interface{}
	}
	tests := []struct {
		name    string
		append  []interface{}
		del     []int
		args    args
		wantErr bool
	}{
		{"TestLinkList_Set case1", []interface{}{}, []int{}, args{1, 1}, false},
		{"TestLinkList_Set case2", []interface{}{0}, []int{}, args{1, 1}, false},
		{"TestLinkList_Set case3", []interface{}{0, 9}, []int{}, args{1, 1}, true},
		{"TestLinkList_Set case4", []interface{}{0, 9, 16, 20}, []int{}, args{0, 1}, true},
		{"TestLinkList_Set case5", []interface{}{0, 9, 16, 20}, []int{}, args{4, 1}, false},
		{"TestLinkList_Set case6", []interface{}{0, 9, 16, 20}, []int{}, args{3, 1}, true},
		{"TestLinkList_Set case7", []interface{}{0, 9, 16, 20}, []int{0}, args{0, 1}, true},
		{"TestLinkList_Set case8", []interface{}{0, 9, 16, 20}, []int{3}, args{2, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if err := this.Set(tt.args.index, tt.args.newVal); (err != nil) == tt.wantErr {
				t.Errorf("LinkList.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkList_Insert(t *testing.T) {
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
		{"TestLinkList_Insert case1", []interface{}{}, []int{}, args{-1, nil}, args{0, []interface{}{1}}, false},
		{"TestLinkList_Insert case2", []interface{}{}, []int{}, args{-1, nil}, args{1, []interface{}{1}}, false},
		{"TestLinkList_Insert case3", []interface{}{}, []int{}, args{-1, nil}, args{0, []interface{}{}}, false},
		{"TestLinkList_Insert case4", []interface{}{}, []int{}, args{-1, nil}, args{1, []interface{}{}}, false},
		{"TestLinkList_Insert case5", []interface{}{1}, []int{}, args{-1, nil}, args{0, []interface{}{}}, true},
		{"TestLinkList_Insert case6", []interface{}{1, 2}, []int{}, args{1, []interface{}{12}}, args{0, []interface{}{12}}, true},
		{"TestLinkList_Insert case7", []interface{}{1, 2, 3}, []int{}, args{3, []interface{}{12}}, args{2, []interface{}{12}}, true},
		{"TestLinkList_Insert case8", []interface{}{1, 2, 3}, []int{2}, args{2, []interface{}{12}}, args{1, []interface{}{12}}, true},
		{"TestLinkList_Insert case9", []interface{}{1, 2, 3}, []int{0}, args{2, []interface{}{12}}, args{1, []interface{}{12}}, true},
		{"TestLinkList_Insert case10", []interface{}{1, 2, 3}, []int{0}, args{2, []interface{}{12}}, args{1, []interface{}{12}}, true},
		{"TestLinkList_Insert case11", []interface{}{1, 2, 3}, []int{0, 0, 0}, args{-1, nil}, args{0, []interface{}{12}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if err := this.Insert(tt.args.index, tt.args.newVal...); (err != nil) == tt.wantErr {
				t.Errorf("LinkList.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.get.index != -1 {
				got, err := this.Get(tt.get.index)
				if (err != nil) == tt.wantErr {
					t.Errorf("LinkList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.get.newVal[0]) {
					t.Errorf("LinkList.Get() = %v, want %v", got, tt.get.newVal[0])
				}
			}
		})
	}
}

func TestLinkList_Delete(t *testing.T) {
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
		{"TestLinkList_Delete case1", []interface{}{}, gots{-1, nil}, args{0}, false},
		{"TestLinkList_Delete case2", []interface{}{1}, gots{-1, nil}, args{2}, false},
		{"TestLinkList_Delete case3", []interface{}{1, 3, 6, 9}, gots{0, 3}, args{0}, true},
		{"TestLinkList_Delete case4", []interface{}{1, 3, 6, 9}, gots{2, 6}, args{3}, true},
		{"TestLinkList_Delete case5", []interface{}{1, 3, 6, 9, 12}, gots{2, 9}, args{2}, true},
		{"TestLinkList_Delete case6", []interface{}{1}, gots{-1, nil}, args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			if err := this.Delete(tt.args.index); (err != nil) == tt.wantErr {
				t.Errorf("LinkList.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.get.index != -1 {
				got, err := this.Get(tt.get.index)
				if (err != nil) == tt.wantErr {
					t.Errorf("LinkList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.get.newVal) {
					t.Errorf("LinkList.Get() = %v, want %v", got, tt.get.newVal)
				}
			}
		})
	}
}

func TestLinkList_Append(t *testing.T) {
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
		{"TestLinkList_Append case1", []int{}, gots{0, 1}, args{[]interface{}{1}}, true},
		{"TestLinkList_Append case2", []int{}, gots{2, 3}, args{[]interface{}{1, 2, 3}}, true},
		{"TestLinkList_Append case3", []int{0}, gots{0, 2}, args{[]interface{}{1, 2, 3}}, true},
		{"TestLinkList_Append case4", []int{2}, gots{0, 1}, args{[]interface{}{1, 2, 3}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.args.newVal...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if tt.get.index != -1 {
				got, err := this.Get(tt.get.index)
				if (err != nil) == tt.wantErr {
					t.Errorf("NewLinkList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.get.newVal) {
					t.Errorf("NewLinkList.Get() = %v, want %v", got, tt.get.newVal)
				}
			}
		})
	}
}

func TestLinkList_Prepend(t *testing.T) {
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
		opt     int
	}{
		{"TestLinkList_Prepend case1", []int{}, gots{0, 1}, args{[]interface{}{1}}, true, 0},
		{"TestLinkList_Prepend case2", []int{}, gots{0, 6}, args{[]interface{}{1, 3, 6}}, true, 0},
		{"TestLinkList_Prepend case3", []int{0}, gots{0, 3}, args{[]interface{}{1, 3, 6}}, true, 0},
		{"TestLinkList_Prepend case4", []int{2}, gots{0, 6}, args{[]interface{}{1, 3, 6}}, true, 0},
		{"TestLinkList_Prepend case5", []int{0}, gots{0, 3}, args{[]interface{}{1, 3, 6}}, true, 1},
		{"TestLinkList_Prepend case6", []int{1}, gots{0, 6}, args{[]interface{}{1, 3, 6}}, true, 1},
		{"TestLinkList_Prepend case7", []int{2}, gots{0, 6}, args{[]interface{}{1, 3, 6}}, true, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			if tt.opt == 0 {
				this.Prepend(tt.args.newVal...)
			} else {
				for i := 0; i < len(tt.args.newVal); i++ {
					this.Prepend(tt.args.newVal[i])
				}
			}
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if tt.get.index != -1 {
				got, err := this.Get(tt.get.index)
				if (err != nil) == tt.wantErr {
					t.Errorf("NewLinkList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.get.newVal) {
					t.Errorf("NewLinkList.Get() = %v, want %v", got, tt.get.newVal)
				}
			}
		})
	}
}

func TestLinkList_Clear(t *testing.T) {

	tests := []struct {
		name   string
		append []interface{}
		del    []int
		want   int
	}{
		{"TestLinkList_Clear case1", []interface{}{}, []int{}, 0},
		{"TestLinkList_Clear case2", []interface{}{1, 2, 3}, []int{}, 0},
		{"TestLinkList_Clear case3", []interface{}{1, 2, 3}, []int{0, 0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			this.Clear()
			if got := this.Size(); got != tt.want {
				t.Errorf("LinkList.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_Contains(t *testing.T) {
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
		{"TestLinkList_Contains case1", []interface{}{}, []int{}, args{[]interface{}{1}}, false},
		{"TestLinkList_Contains case2", []interface{}{3, 6, 9}, []int{}, args{[]interface{}{1}}, false},
		{"TestLinkList_Contains case3", []interface{}{3, 6, 9}, []int{}, args{[]interface{}{1, 2}}, false},
		{"TestLinkList_Contains case4", []interface{}{3, 6, 9}, []int{}, args{[]interface{}{3}}, true},
		{"TestLinkList_Contains case5", []interface{}{3, 6, 9}, []int{}, args{[]interface{}{9}}, true},
		{"TestLinkList_Contains case6", []interface{}{3, 6, 9}, []int{}, args{[]interface{}{3, 6, 9}}, true},
		{"TestLinkList_Contains case7", []interface{}{3, 6, 9}, []int{0}, args{[]interface{}{3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if got := this.Contains(tt.args.values...); got != tt.want {
				t.Errorf("LinkList.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_Swap(t *testing.T) {
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
		{"TestLinkList_Swap case1", []interface{}{}, []int{}, gots{-1, nil}, args{0, 1}, false},
		{"TestLinkList_Swap case2", []interface{}{1, 3, 6, 9}, []int{}, gots{-1, nil}, args{0, 0}, true},
		{"TestLinkList_Swap case3", []interface{}{1, 3, 6, 9}, []int{}, gots{0, 3}, args{0, 1}, true},
		{"TestLinkList_Swap case4", []interface{}{1, 3, 6, 9}, []int{}, gots{0, 3}, args{0, 1}, true},
		{"TestLinkList_Swap case5", []interface{}{1, 3, 6, 9}, []int{0}, gots{0, 6}, args{0, 1}, true},
		{"TestLinkList_Swap case6", []interface{}{1, 3, 6, 9}, []int{3}, gots{1, 6}, args{2, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if err := this.Swap(tt.args.i, tt.args.j); (err != nil) == tt.wantErr {
				t.Errorf("LinkList.Swap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.get.index != -1 {
				got, err := this.Get(tt.get.index)
				if (err != nil) == tt.wantErr {
					t.Errorf("LinkList.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.get.newVal) {
					t.Errorf("LinkList.Get() = %v, want %v", got, tt.get.newVal)
				}
			}
		})
	}
}

func TestLinkList_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		want   bool
	}{
		{"TestLinkList_IsEmpty case1", []interface{}{}, []int{}, true},
		{"TestLinkList_IsEmpty case2", []interface{}{1}, []int{}, false},
		{"TestLinkList_IsEmpty case3", []interface{}{1}, []int{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if got := this.IsEmpty(); got != tt.want {
				t.Errorf("LinkList.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_Size(t *testing.T) {
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		want   int
	}{
		{"TestLinkList_Size case1", []interface{}{}, []int{}, 0},
		{"TestLinkList_Size case2", []interface{}{1, 3, 6}, []int{}, 3},
		{"TestLinkList_Size case3", []interface{}{1, 3, 6}, []int{0}, 2},
		{"TestLinkList_Size case4", []interface{}{1, 3, 6}, []int{2}, 2},
		{"TestLinkList_Size case5", []interface{}{1, 3, 6}, []int{0, 0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if got := this.Size(); got != tt.want {
				t.Errorf("LinkList.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_withinRange(t *testing.T) {
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
		{"TestLinkList_withinRange case1", []interface{}{}, []int{}, args{0}, false},
		{"TestLinkList_withinRange case2", []interface{}{1, 6, 9}, []int{}, args{0}, true},
		{"TestLinkList_withinRange case3", []interface{}{1, 6, 9}, []int{}, args{2}, true},
		{"TestLinkList_withinRange case4", []interface{}{1, 6, 9}, []int{2}, args{2}, false},
		{"TestLinkList_withinRange case5", []interface{}{1, 6, 9}, []int{0}, args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			if got := this.withinRange(tt.args.index); got != tt.want {
				t.Errorf("LinkList.withinRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_HeadNode(t *testing.T) {
	tests := []struct {
		name    string
		append  []interface{}
		prepend []interface{}
		del     []int
		want    interface{}
	}{
		{"TestLinkList_HeadNode case1", []interface{}{}, []interface{}{}, []int{}, (*Node)(nil)},
		{"TestLinkList_HeadNode case2", []interface{}{1, 3, 6}, []interface{}{}, []int{}, 1},
		{"TestLinkList_HeadNode case3", []interface{}{1, 3, 6}, []interface{}{}, []int{0, 0}, 6},
		{"TestLinkList_HeadNode case4", []interface{}{}, []interface{}{1, 3, 6}, []int{}, 6},
		{"TestLinkList_HeadNode case5", []interface{}{}, []interface{}{1, 3, 6}, []int{0, 0}, 1},
		{"TestLinkList_HeadNode case6", []interface{}{2, 4, 9}, []interface{}{1, 3, 6}, []int{}, 6},
		{"TestLinkList_HeadNode case7", []interface{}{2, 4, 9}, []interface{}{1, 3, 6}, []int{0, 0}, 1},
		{"TestLinkList_HeadNode case8", []interface{}{2, 4, 9}, []interface{}{1, 3, 6}, []int{0, 0, 0}, 2},
		{"TestLinkList_HeadNode case9", []interface{}{2, 4, 9}, []interface{}{1, 3, 6}, []int{0, 0, 0, 0, 0, 0}, (*Node)(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			this.Prepend(tt.prepend...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			got := this.HeadNode()
			if got == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkList.HeadNode() = %v, want %v", got, tt.want)
				}
			} else if !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("LinkList.HeadNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_TailNode(t *testing.T) {
	tests := []struct {
		name    string
		append  []interface{}
		prepend []interface{}
		del     []int
		want    interface{}
	}{
		{"TestLinkList_TailNode case1", []interface{}{}, []interface{}{}, []int{}, (*Node)(nil)},
		{"TestLinkList_TailNode case2", []interface{}{1, 3, 6}, []interface{}{}, []int{}, 6},
		{"TestLinkList_TailNode case3", []interface{}{1, 3, 6}, []interface{}{}, []int{0, 0}, 6},
		{"TestLinkList_TailNode case4", []interface{}{}, []interface{}{1, 3, 6}, []int{}, 1},
		{"TestLinkList_TailNode case5", []interface{}{}, []interface{}{1, 3, 6}, []int{0, 0}, 1},
		{"TestLinkList_TailNode case6", []interface{}{2, 4, 9}, []interface{}{1, 3, 6}, []int{}, 9},
		{"TestLinkList_TailNode case7", []interface{}{2, 4, 9}, []interface{}{1, 3, 6}, []int{0, 0}, 9},
		{"TestLinkList_TailNode case8", []interface{}{2, 4, 9}, []interface{}{1, 3, 6}, []int{0, 0, 0}, 9},
		{"TestLinkList_TailNode case9", []interface{}{2, 4, 9}, []interface{}{1, 3, 6}, []int{0, 0, 0, 0, 0, 0}, (*Node)(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinkList()
			this.Append(tt.append...)
			this.Prepend(tt.prepend...)
			for i := 0; i < len(tt.del); i++ {
				this.Delete(tt.del[i])
			}
			got := this.TailNode()
			if got == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkList.TailNode() = %v, want %v", got, tt.want)
				}
			} else if !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("LinkList.TailNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
