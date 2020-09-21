package dlinkList

import (
	"reflect"
	"testing"
)

func TestNewDLinkList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestNewDLinkList case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			if got == nil {
				t.Errorf("got nil, want DLinkList")
				return
			}
			if got.head != nil {
				t.Errorf("got not nil, want DLinkList head nil.")
			}
			if got.tail != nil {
				t.Errorf("got not nil, want DLinkList tail nil.")
			}
			if got.size != 0 {
				t.Errorf("got %v, want %v", got.size, 0)
			}
		})
	}
}

func TestNewNode(t *testing.T) {
	tests := []struct {
		name   string
		newVal interface{}
		want   interface{}
	}{
		{"TestNewNode case1", 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNode(tt.newVal)
			if got == nil {
				t.Errorf("got nil, want NewNode")
				return
			}
			if !reflect.DeepEqual(got.val, tt.want) {
				t.Errorf("DLinkList.NewNode() = %v, want %v", got.val, tt.want)
			}
			if got.pre != nil {
				t.Errorf("got %v, want %v", got.pre, nil)
			}
			if got.next != nil {
				t.Errorf("got %v, want %v", got.next, nil)
			}
		})
	}
}

func TestDLinkList_Get(t *testing.T) {
	tests := []struct {
		name    string
		append  []interface{}
		args    int
		want    interface{}
		wantErr bool
	}{
		{"TestDLinkList_Get case1", []interface{}{}, 0, nil, false},
		{"TestDLinkList_Get case2", []interface{}{2}, 0, 2, true},
		{"TestDLinkList_Get case3", []interface{}{2}, 1, nil, false},
		{"TestDLinkList_Get case4", []interface{}{2, 6, 8, 11}, 3, 11, true},
		{"TestDLinkList_Get case5", []interface{}{2, 6, 8, 11}, 0, 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			value, err := got.Get(tt.args)
			if (err != nil) == tt.wantErr {
				t.Errorf("DLinkList.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(value, tt.want) {
				t.Errorf("DLinkList.Get() = %v, want %v", value, tt.want)
			}
		})
	}
}

func TestDLinkList_Set(t *testing.T) {
	type args struct {
		index  int
		newVal interface{}
	}
	tests := []struct {
		name     string
		append   []interface{}
		args     []args
		wantErr  bool
		get      int
		getValue interface{}
		getErr   bool
	}{
		{"TestDLinkList_Set case1", []interface{}{}, []args{{0, 10}}, false, 0, nil, false},
		{"TestDLinkList_Set case2", []interface{}{1}, []args{{0, 10}}, true, 0, 10, true},
		{"TestDLinkList_Set case3", []interface{}{1}, []args{{0, 10}, {0, 20}}, true, 0, 20, true},
		{"TestDLinkList_Set case4", []interface{}{1, 6, 8, 11}, []args{{3, 10}}, true, 3, 10, true},
		{"TestDLinkList_Set case5", []interface{}{1, 6, 8, 11}, []args{{0, 10}}, true, 0, 10, true},
		{"TestDLinkList_Set case6", []interface{}{1, 6, 8, 11}, []args{{6, 10}}, false, 6, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			for i := 0; i < len(tt.args); i++ {
				if err := got.Set(tt.args[i].index, tt.args[i].newVal); (err != nil) == tt.wantErr {
					t.Errorf("DLinkList.Set() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			value, err := got.Get(tt.get)
			if (err != nil) == tt.getErr {
				t.Errorf("DLinkList.Get() error = %v, getErr %v", err, tt.getErr)
				return
			}
			if err == nil && !reflect.DeepEqual(value, tt.getValue) {
				t.Errorf("DLinkList.Get() = %v, getValue %v", value, tt.getValue)
			}
		})
	}
}

func TestDLinkList_Insert(t *testing.T) {
	type args struct {
		index  int
		newVal []interface{}
	}
	tests := []struct {
		name     string
		append   []interface{}
		args     args
		wantErr  bool
		get      int
		getValue interface{}
		getErr   bool
	}{
		{"TestDLinkList_Insert case1", []interface{}{}, args{0, []interface{}{10}}, false, 0, nil, false},
		{"TestDLinkList_Insert case2", []interface{}{1}, args{0, []interface{}{10}}, true, 0, 1, true},
		{"TestDLinkList_Insert case3", []interface{}{1, 6, 8}, args{0, []interface{}{10}}, true, 1, 10, true},
		{"TestDLinkList_Insert case4", []interface{}{1, 6, 8}, args{1, []interface{}{10}}, true, 2, 10, true},
		{"TestDLinkList_Insert case5", []interface{}{1, 6, 8}, args{2, []interface{}{10}}, true, 3, 10, true},
		{"TestDLinkList_Insert case6", []interface{}{1, 6, 8}, args{2, []interface{}{10}}, true, 2, 8, true},

		{"TestDLinkList_Insert case7", []interface{}{1, 6, 8, 9}, args{2, []interface{}{10}}, true, 3, 10, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			if err := got.Insert(tt.args.index, tt.args.newVal...); (err != nil) == tt.wantErr {
				t.Errorf("DLinkList.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			value, err := got.Get(tt.get)
			if (err != nil) == tt.getErr {
				t.Errorf("DLinkList.Get() error = %v, getErr %v", err, tt.getErr)
				return
			}
			if err == nil && !reflect.DeepEqual(value, tt.getValue) {
				t.Errorf("DLinkList.Get() = %v, getValue %v", value, tt.getValue)
			}
		})
	}
}

func TestDLinkList_Delete(t *testing.T) {
	tests := []struct {
		name     string
		append   []interface{}
		args     int
		wantErr  bool
		get      int
		getValue interface{}
		getErr   bool
	}{
		{"TestDLinkList_Delete case1", []interface{}{}, 0, false, 0, nil, false},
		{"TestDLinkList_Delete case2", []interface{}{1}, 0, true, 0, nil, false},
		{"TestDLinkList_Delete case3", []interface{}{1, 6, 8}, 0, true, 0, 6, true},
		{"TestDLinkList_Delete case4", []interface{}{1, 6, 8}, 2, true, 0, 1, true},
		{"TestDLinkList_Delete case5", []interface{}{1, 6, 8}, 1, true, 1, 8, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			if err := got.Delete(tt.args); (err != nil) == tt.wantErr {
				t.Errorf("DLinkList.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			value, err := got.Get(tt.get)
			if (err != nil) == tt.getErr {
				t.Errorf("DLinkList.Get() error = %v, getErr %v", err, tt.getErr)
				return
			}
			if err == nil && !reflect.DeepEqual(value, tt.getValue) {
				t.Errorf("DLinkList.Get() = %v, getValue %v", value, tt.getValue)
			}
		})
	}
}

func TestDLinkList_DeleteNode(t *testing.T) {
	tests := []struct {
		name     string
		append   []interface{}
		args     int
		wantErr  bool
		get      int
		getValue interface{}
		getErr   bool
	}{
		{"TestDLinkList_DeleteNode case1", []interface{}{}, 0, false, 0, nil, false},
		{"TestDLinkList_DeleteNode case2", []interface{}{1}, 0, true, 0, nil, false},
		{"TestDLinkList_DeleteNode case3", []interface{}{1, 6, 8}, 0, true, 0, 6, true},
		{"TestDLinkList_DeleteNode case4", []interface{}{1, 6, 8}, 1, true, 1, 8, true},
		{"TestDLinkList_DeleteNode case5", []interface{}{1, 6, 8}, 2, true, 1, 6, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			target := got.HeadNode()
			for i := 0; i < tt.args; i++ {
				if target != nil {
					target = target.next
				}
			}
			if err := got.DeleteNode(target); (err != nil) == tt.wantErr {
				t.Errorf("DLinkList.DeleteNode() error = %v, wantErr %v", err, tt.wantErr)
			}
			value, err := got.Get(tt.get)
			if (err != nil) == tt.getErr {
				t.Errorf("DLinkList.Get() error = %v, getErr %v", err, tt.getErr)
				return
			}
			if err == nil && !reflect.DeepEqual(value, tt.getValue) {
				t.Errorf("DLinkList.Get() = %v, getValue %v", value, tt.getValue)
			}
		})
	}
}

func TestDLinkList_Append(t *testing.T) {
	tests := []struct {
		name     string
		args     []interface{}
		get      int
		getValue interface{}
		getErr   bool
		want     int
	}{
		{"TestDLinkList_Append case1", []interface{}{}, 0, nil, false, 0},
		{"TestDLinkList_Append case2", []interface{}{1}, 1, nil, false, 1},
		{"TestDLinkList_Append case3", []interface{}{1, 6, 8}, 0, 1, true, 3},
		{"TestDLinkList_Append case4", []interface{}{1, 6, 8}, 1, 6, true, 3},
		{"TestDLinkList_Append case5", []interface{}{1, 6, 8}, 2, 8, true, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.args...)
			value, err := got.Get(tt.get)
			if (err != nil) == tt.getErr {
				t.Errorf("DLinkList.Get() error = %v, getErr %v", err, tt.getErr)
				return
			}
			if err == nil && !reflect.DeepEqual(value, tt.getValue) {
				t.Errorf("DLinkList.Get() = %v, getValue %v", value, tt.getValue)
			}
			if size := got.Size(); size != tt.want {
				t.Errorf("DLinkList.Size() = %v, want %v", size, tt.want)
			}
		})
	}
}

func TestDLinkList_Prepend(t *testing.T) {
	tests := []struct {
		name     string
		args     []interface{}
		get      int
		getValue interface{}
		getErr   bool
		want     int
	}{
		{"TestDLinkList_Prepend case1", []interface{}{}, 0, nil, false, 0},
		{"TestDLinkList_Prepend case2", []interface{}{1}, 1, nil, false, 1},
		{"TestDLinkList_Prepend case3", []interface{}{1, 6, 8}, 0, 8, true, 3},
		{"TestDLinkList_Prepend case4", []interface{}{1, 6, 8}, 1, 6, true, 3},
		{"TestDLinkList_Prepend case5", []interface{}{1, 6, 8}, 2, 1, true, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Prepend(tt.args...)
			value, err := got.Get(tt.get)
			if (err != nil) == tt.getErr {
				t.Errorf("DLinkList.Get() error = %v, getErr %v", err, tt.getErr)
				return
			}
			if err == nil && !reflect.DeepEqual(value, tt.getValue) {
				t.Errorf("DLinkList.Get() = %v, getValue %v", value, tt.getValue)
			}
			if size := got.Size(); size != tt.want {
				t.Errorf("DLinkList.Size() = %v, want %v", size, tt.want)
			}
		})
	}
}

func TestDLinkList_Contains(t *testing.T) {
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		args   []interface{}
		want   bool
	}{
		{"TestDLinkList_Contains case1", []interface{}{}, []int{}, []interface{}{}, true},
		{"TestDLinkList_Contains case2", []interface{}{}, []int{}, []interface{}{1}, false},
		{"TestDLinkList_Contains case3", []interface{}{1}, []int{}, []interface{}{2}, false},
		{"TestDLinkList_Contains case4", []interface{}{1, 6, 8}, []int{}, []interface{}{1}, true},
		{"TestDLinkList_Contains case5", []interface{}{1, 6, 8}, []int{}, []interface{}{6}, true},
		{"TestDLinkList_Contains case6", []interface{}{1, 6, 8}, []int{}, []interface{}{8}, true},
		{"TestDLinkList_Contains case7", []interface{}{1, 6, 8}, []int{}, []interface{}{9}, false},
		{"TestDLinkList_Contains case8", []interface{}{1, 6, 8}, []int{0}, []interface{}{1}, false},
		{"TestDLinkList_Contains case9", []interface{}{1, 6, 8}, []int{1}, []interface{}{6}, false},
		{"TestDLinkList_Contains case10", []interface{}{1, 6, 8}, []int{2}, []interface{}{8}, false},
		{"TestDLinkList_Contains case11", []interface{}{1, 6, 8}, []int{2}, []interface{}{6}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			if isContain := got.Contains(tt.args...); isContain != tt.want {
				t.Errorf("DLinkList.Contains() = %v, want %v", isContain, tt.want)
			}
		})
	}
}

func TestDLinkList_Clear(t *testing.T) {
	tests := []struct {
		name     string
		append   []interface{}
		del      []int
		contains []interface{}
		wantErr  bool
		want     int
	}{
		{"TestDLinkList_Clear case1", []interface{}{}, []int{}, []interface{}{}, true, 0},
		{"TestDLinkList_Clear case2", []interface{}{1, 6, 8}, []int{}, []interface{}{}, true, 0},
		{"TestDLinkList_Clear case3", []interface{}{1, 6, 8}, []int{0}, []interface{}{}, true, 0},
		{"TestDLinkList_Clear case4", []interface{}{1, 6, 8}, []int{1}, []interface{}{}, true, 0},
		{"TestDLinkList_Clear case5", []interface{}{1, 6, 8}, []int{2}, []interface{}{}, true, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			got.Clear()
			if isContain := got.Contains(tt.contains...); isContain != tt.wantErr {
				t.Errorf("DLinkList.Contains() = %v, wantErr %v", isContain, tt.wantErr)
			}
			if size := got.Size(); size != tt.want {
				t.Errorf("DLinkList.Size() = %v, want %v", size, tt.want)
			}
		})
	}
}

func TestDLinkList_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name     string
		append   []interface{}
		del      []int
		args     args
		wantErr  bool
		get      int
		getValue interface{}
		getErr   bool
	}{
		{"TestDLinkList_Swap case1", []interface{}{}, []int{}, args{0, 1}, false, 0, nil, false},
		{"TestDLinkList_Swap case2", []interface{}{1}, []int{}, args{0, 0}, true, 0, 1, true},
		{"TestDLinkList_Swap case3", []interface{}{1, 2}, []int{}, args{0, 1}, true, 0, 2, true},
		{"TestDLinkList_Swap case4", []interface{}{1, 2}, []int{}, args{1, 0}, true, 0, 2, true},
		{"TestDLinkList_Swap case5", []interface{}{1, 6, 8}, []int{}, args{2, 0}, true, 0, 8, true},
		{"TestDLinkList_Swap case6", []interface{}{1, 6, 8}, []int{1}, args{1, 0}, true, 0, 8, true},
		{"TestDLinkList_Swap case7", []interface{}{1, 6, 8}, []int{1}, args{1, 0}, true, 1, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			if err := got.Swap(tt.args.i, tt.args.j); (err != nil) == tt.wantErr {
				t.Errorf("DLinkList.Swap() error = %v, wantErr %v", err, tt.wantErr)
			}
			value, err := got.Get(tt.get)
			if (err != nil) == tt.getErr {
				t.Errorf("DLinkList.Get() error = %v, getErr %v", err, tt.getErr)
				return
			}
			if err == nil && !reflect.DeepEqual(value, tt.getValue) {
				t.Errorf("DLinkList.Get() = %v, getValue %v", value, tt.getValue)
			}
		})
	}
}

func TestDLinkList_Size(t *testing.T) {
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		want   int
	}{
		{"TestDLinkList_Size case1", []interface{}{}, []int{}, 0},
		{"TestDLinkList_Size case2", []interface{}{1}, []int{0}, 0},
		{"TestDLinkList_Size case3", []interface{}{1, 6, 8}, []int{0}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			if size := got.Size(); size != tt.want {
				t.Errorf("DLinkList.Size() = %v, want %v", size, tt.want)
			}
		})
	}
}

func TestDLinkList_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		want   bool
	}{
		{"TestDLinkList_IsEmpty case1", []interface{}{}, []int{}, true},
		{"TestDLinkList_IsEmpty case2", []interface{}{1}, []int{}, false},
		{"TestDLinkList_IsEmpty case3", []interface{}{1}, []int{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			if isEmpty := got.IsEmpty(); isEmpty != tt.want {
				t.Errorf("DLinkList.IsEmpty() = %v, want %v", isEmpty, tt.want)
			}
		})
	}
}

func TestDLinkList_HeadNode(t *testing.T) {
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		want   interface{}
	}{
		{"TestDLinkList_HeadNode case1", []interface{}{}, []int{}, (*Node)(nil)},
		{"TestDLinkList_HeadNode case2", []interface{}{1}, []int{}, 1},
		{"TestDLinkList_HeadNode case3", []interface{}{1}, []int{0}, (*Node)(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			headNode := got.HeadNode()
			if headNode == nil {
				if !reflect.DeepEqual(headNode, tt.want) {
					t.Errorf("DLinkList.HeadNode() = %v, want %v", headNode, tt.want)
				}
			} else if !reflect.DeepEqual(headNode.val, tt.want) {
				t.Errorf("DLinkList.HeadNode() = %v, want %v", headNode, tt.want)
			}
		})
	}
}

func TestDLinkList_TailNode(t *testing.T) {
	tests := []struct {
		name   string
		append []interface{}
		del    []int
		want   interface{}
	}{
		{"TestDLinkList_TailNode case1", []interface{}{}, []int{}, (*Node)(nil)},
		{"TestDLinkList_TailNode case2", []interface{}{1}, []int{}, 1},
		{"TestDLinkList_TailNode case3", []interface{}{1}, []int{0}, (*Node)(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDLinkList()
			got.Append(tt.append...)
			for i := 0; i < len(tt.del); i++ {
				got.Delete(tt.del[i])
			}
			tailNode := got.TailNode()
			if tailNode == nil {
				if !reflect.DeepEqual(tailNode, tt.want) {
					t.Errorf("DLinkList.TailNode() = %v, want %v", tailNode, tt.want)
				}
			} else if !reflect.DeepEqual(tailNode.val, tt.want) {
				t.Errorf("DLinkList.TailNode() = %v, want %v", tailNode, tt.want)
			}
		})
	}
}
