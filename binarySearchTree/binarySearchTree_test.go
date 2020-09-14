package binarySearchTree

import (
	"reflect"

	"testing"

	"github.com/GoodSpeedZhao/bds/comparable"
)

func TestNewBinarySearchTree(t *testing.T) {
	tests := []struct {
		name string
		cmp  comparable.Comparator
	}{
		{"TestNewBinarySearchTree case1", comparable.IntComparator},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.cmp)
			if got == nil {
				t.Errorf("got nil, want BinarySearchTree")
				return
			}
			if got.root != nil {
				t.Errorf("got not nil, want BinarySearchTree root nil.")
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
			if !reflect.DeepEqual(got.value, tt.want) {
				t.Errorf("BinarySearchTree.NewNode() = %v, want %v", got.value, tt.want)
			}
			if got.left != nil {
				t.Errorf("got %v, want %v", got.left, nil)
			}
			if got.right != nil {
				t.Errorf("got %v, want %v", got.right, nil)
			}
		})
	}
}

func TestBinarySearchTree_Insert(t *testing.T) {
	tests := []struct {
		name    string
		cmp     comparable.Comparator
		args    []interface{}
		wantErr []bool
		size    int
	}{
		{"TestBinarySearchTree_Insert case1", comparable.IntComparator, []interface{}{1}, []bool{true}, 1},
		{"TestBinarySearchTree_Insert case2", comparable.IntComparator, []interface{}{1, 1}, []bool{true, false}, 1},
		{"TestBinarySearchTree_Insert case3", comparable.IntComparator, []interface{}{6, 1}, []bool{true, true}, 2},
		{"TestBinarySearchTree_Insert case4", comparable.IntComparator, []interface{}{6, 9}, []bool{true, true}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.cmp)
			for i := 0; i < len(tt.args); i++ {
				if err := got.Insert(tt.args[i]); (err != nil) == tt.wantErr[i] {
					t.Errorf("BinarySearchTree.Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			if size := got.Size(); size != tt.size {
				t.Errorf("BinarySearchTree.Size() = %v, want %v", size, tt.size)
			}
		})
	}
}

func TestBinarySearchTree_Delete(t *testing.T) {
	tests := []struct {
		name    string
		cmp     comparable.Comparator
		insert  []interface{}
		args    []interface{}
		wantErr []bool
		size    int
	}{
		{"TestBinarySearchTree_Delete case1", comparable.IntComparator, []interface{}{}, []interface{}{1}, []bool{false}, 0},
		{"TestBinarySearchTree_Delete case2", comparable.IntComparator, []interface{}{1}, []interface{}{1}, []bool{true}, 0},
		{"TestBinarySearchTree_Delete case3", comparable.IntComparator, []interface{}{6, 1, 9}, []interface{}{1}, []bool{true}, 2},
		{"TestBinarySearchTree_Delete case4", comparable.IntComparator, []interface{}{6, 1, 9}, []interface{}{9}, []bool{true}, 2},
		{"TestBinarySearchTree_Delete case5", comparable.IntComparator, []interface{}{6, 1, 9}, []interface{}{6}, []bool{true}, 2},
		{"TestBinarySearchTree_Delete case6", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{1}, []bool{true}, 2},
		{"TestBinarySearchTree_Delete case7", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{6}, []bool{true}, 2},
		{"TestBinarySearchTree_Delete case8", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{9}, []bool{true}, 2},
		{"TestBinarySearchTree_Delete case9", comparable.IntComparator, []interface{}{9, 6, 1}, []interface{}{1}, []bool{true}, 2},
		{"TestBinarySearchTree_Delete case10", comparable.IntComparator, []interface{}{9, 6, 1}, []interface{}{6}, []bool{true}, 2},
		{"TestBinarySearchTree_Delete case11", comparable.IntComparator, []interface{}{9, 6, 1}, []interface{}{9}, []bool{true}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.cmp)
			for i := 0; i < len(tt.insert); i++ {
				got.Insert(tt.insert[i])
			}
			for i := 0; i < len(tt.args); i++ {
				if err := got.Delete(tt.args[i]); (err != nil) == tt.wantErr[i] {
					t.Errorf("BinarySearchTree.Delete() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			if size := got.Size(); size != tt.size {
				t.Errorf("BinarySearchTree.Size() = %v, want %v", size, tt.size)
			}
		})
	}
}

func TestBinarySearchTree_DeleteMin(t *testing.T) {
	tests := []struct {
		name    string
		cmp     comparable.Comparator
		insert  []interface{}
		wantErr bool
		size    int
	}{
		{"TestBinarySearchTree_DeleteMin case1", comparable.IntComparator, []interface{}{}, false, 0},
		{"TestBinarySearchTree_DeleteMin case2", comparable.IntComparator, []interface{}{9, 6, 1}, true, 2},
		{"TestBinarySearchTree_DeleteMin case3", comparable.IntComparator, []interface{}{1, 6, 9}, true, 2},
		{"TestBinarySearchTree_DeleteMin case4", comparable.IntComparator, []interface{}{6, 1, 9}, true, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.cmp)
			for i := 0; i < len(tt.insert); i++ {
				got.Insert(tt.insert[i])
			}
			if err := got.DeleteMin(); (err != nil) == tt.wantErr {
				t.Errorf("BinarySearchTree.DeleteMin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBinarySearchTree_DeleteMax(t *testing.T) {
	tests := []struct {
		name    string
		cmp     comparable.Comparator
		insert  []interface{}
		wantErr bool
		size    int
	}{
		{"TestBinarySearchTree_DeleteMax case1", comparable.IntComparator, []interface{}{}, false, 0},
		{"TestBinarySearchTree_DeleteMax case2", comparable.IntComparator, []interface{}{9, 6, 1}, true, 2},
		{"TestBinarySearchTree_DeleteMax case3", comparable.IntComparator, []interface{}{1, 6, 9}, true, 2},
		{"TestBinarySearchTree_DeleteMax case4", comparable.IntComparator, []interface{}{6, 1, 9}, true, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.cmp)
			for i := 0; i < len(tt.insert); i++ {
				got.Insert(tt.insert[i])
			}
			if err := got.DeleteMax(); (err != nil) == tt.wantErr {
				t.Errorf("BinarySearchTree.DeleteMax() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBinarySearchTree_Contains(t *testing.T) {
	tests := []struct {
		name    string
		cmp     comparable.Comparator
		insert  []interface{}
		delMax  bool
		delMin  bool
		args    []interface{}
		wantErr []bool
	}{
		{"TestBinarySearchTree_Contains case1", comparable.IntComparator, []interface{}{}, false, false, []interface{}{1}, []bool{false}},
		{"TestBinarySearchTree_Contains case2", comparable.IntComparator, []interface{}{6, 1, 9}, false, false, []interface{}{10}, []bool{false}},
		{"TestBinarySearchTree_Contains case3", comparable.IntComparator, []interface{}{6, 1, 9}, false, false, []interface{}{1}, []bool{true}},
		{"TestBinarySearchTree_Contains case4", comparable.IntComparator, []interface{}{6, 1, 9}, false, false, []interface{}{6}, []bool{true}},
		{"TestBinarySearchTree_Contains case5", comparable.IntComparator, []interface{}{6, 1, 9}, false, false, []interface{}{9}, []bool{true}},
		{"TestBinarySearchTree_Contains case6", comparable.IntComparator, []interface{}{1, 6, 9}, false, false, []interface{}{1}, []bool{true}},
		{"TestBinarySearchTree_Contains case7", comparable.IntComparator, []interface{}{1, 6, 9}, false, false, []interface{}{6}, []bool{true}},
		{"TestBinarySearchTree_Contains case8", comparable.IntComparator, []interface{}{1, 6, 9}, false, false, []interface{}{9}, []bool{true}},
		{"TestBinarySearchTree_Contains case9", comparable.IntComparator, []interface{}{1, 6, 9}, true, false, []interface{}{9}, []bool{false}},
		{"TestBinarySearchTree_Contains case10", comparable.IntComparator, []interface{}{1, 6, 9}, false, true, []interface{}{1}, []bool{false}},
		{"TestBinarySearchTree_Contains case11", comparable.IntComparator, []interface{}{1, 6, 9}, true, true, []interface{}{1, 9}, []bool{false, false}},
		{"TestBinarySearchTree_Contains case12", comparable.IntComparator, []interface{}{9, 6, 1}, true, false, []interface{}{9}, []bool{false}},
		{"TestBinarySearchTree_Contains case13", comparable.IntComparator, []interface{}{9, 6, 1}, false, true, []interface{}{1}, []bool{false}},
		{"TestBinarySearchTree_Contains case14", comparable.IntComparator, []interface{}{9, 6, 1}, true, true, []interface{}{1, 9}, []bool{false, false}},
		{"TestBinarySearchTree_Contains case15", comparable.IntComparator, []interface{}{6, 1, 9}, true, false, []interface{}{9}, []bool{false}},
		{"TestBinarySearchTree_Contains case16", comparable.IntComparator, []interface{}{6, 1, 9}, false, true, []interface{}{1}, []bool{false}},
		{"TestBinarySearchTree_Contains case17", comparable.IntComparator, []interface{}{6, 1, 9}, true, true, []interface{}{1, 9}, []bool{false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.cmp)
			for i := 0; i < len(tt.insert); i++ {
				got.Insert(tt.insert[i])
			}
			if tt.delMax {
				got.DeleteMax()
			}
			if tt.delMin {
				got.DeleteMin()
			}
			for i := 0; i < len(tt.args); i++ {
				if isContain := got.Contains(tt.args[i]); isContain != tt.wantErr[i] {
					t.Errorf("BinarySearchTree.Contains() = %v, want %v", isContain, tt.wantErr[i])
				}
			}
		})
	}
}

func TestBinarySearchTree_FindMin(t *testing.T) {
	tests := []struct {
		name    string
		cmp     comparable.Comparator
		insert  []interface{}
		want    interface{}
		wantErr bool
	}{
		{"TestBinarySearchTree_FindMin case1", comparable.IntComparator, []interface{}{}, nil, false},
		{"TestBinarySearchTree_FindMin case2", comparable.IntComparator, []interface{}{1, 6, 9}, 1, true},
		{"TestBinarySearchTree_FindMin case3", comparable.IntComparator, []interface{}{9, 6, 1}, 1, true},
		{"TestBinarySearchTree_FindMin case4", comparable.IntComparator, []interface{}{6, 1, 9}, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.cmp)
			for i := 0; i < len(tt.insert); i++ {
				got.Insert(tt.insert[i])
			}
			element, err := got.FindMin()
			if (err != nil) == tt.wantErr {
				t.Errorf("BinarySearchTree.FindMin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(element, tt.want) {
				t.Errorf("BinarySearchTree.FindMin() = %v, want %v", element, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_FindMax(t *testing.T) {
	tests := []struct {
		name    string
		cmp     comparable.Comparator
		insert  []interface{}
		want    interface{}
		wantErr bool
	}{
		{"TestBinarySearchTree_FindMax case1", comparable.IntComparator, []interface{}{}, nil, false},
		{"TestBinarySearchTree_FindMax case2", comparable.IntComparator, []interface{}{1, 6, 9}, 9, true},
		{"TestBinarySearchTree_FindMax case3", comparable.IntComparator, []interface{}{9, 6, 1}, 9, true},
		{"TestBinarySearchTree_FindMax case4", comparable.IntComparator, []interface{}{6, 1, 9}, 9, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.cmp)
			for i := 0; i < len(tt.insert); i++ {
				got.Insert(tt.insert[i])
			}
			element, err := got.FindMax()
			if (err != nil) == tt.wantErr {
				t.Errorf("BinarySearchTree.FindMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(element, tt.want) {
				t.Errorf("BinarySearchTree.FindMax() = %v, want %v", element, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		cmp    comparable.Comparator
		insert []interface{}
		delete []interface{}
		want   bool
	}{
		{"TestBinarySearchTree_IsEmpty case1", comparable.IntComparator, []interface{}{}, []interface{}{}, true},
		{"TestBinarySearchTree_IsEmpty case2", comparable.IntComparator, []interface{}{8, 6, 1}, []interface{}{1, 6, 8}, true},
		{"TestBinarySearchTree_IsEmpty case3", comparable.IntComparator, []interface{}{1, 6, 8}, []interface{}{1, 6, 8}, true},
		{"TestBinarySearchTree_IsEmpty case4", comparable.IntComparator, []interface{}{6, 1, 8}, []interface{}{1, 6, 8}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.cmp)
			for i := 0; i < len(tt.insert); i++ {
				got.Insert(tt.insert[i])
			}
			for i := 0; i < len(tt.delete); i++ {
				got.Delete(tt.delete[i])
			}
			if isEmpty := got.IsEmpty(); isEmpty != tt.want {
				t.Errorf("BinarySearchTree.IsEmpty() = %v, want %v", isEmpty, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Size(t *testing.T) {
	tests := []struct {
		name   string
		cmp    comparable.Comparator
		insert []interface{}
		delete []interface{}
		want   int
	}{
		{"TestBinarySearchTree_Size case1", comparable.IntComparator, []interface{}{}, []interface{}{}, 0},
		{"TestBinarySearchTree_Size case2", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{}, 3},
		{"TestBinarySearchTree_Size case3", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{1}, 2},
		{"TestBinarySearchTree_Size case4", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{1, 6}, 1},
		{"TestBinarySearchTree_Size case5", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, 0},
		{"TestBinarySearchTree_Size case6", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{1, 6, 9, 10}, 0},
		{"TestBinarySearchTree_Size case7", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{10}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.cmp)
			for i := 0; i < len(tt.insert); i++ {
				got.Insert(tt.insert[i])
			}
			for i := 0; i < len(tt.delete); i++ {
				got.Delete(tt.delete[i])
			}
			if size := got.Size(); size != tt.want {
				t.Errorf("BinarySearchTree.Size() = %v, want %v", size, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Clear(t *testing.T) {
	tests := []struct {
		name   string
		cmp    comparable.Comparator
		insert []interface{}
		delete []interface{}
		want   int
	}{
		{"TestBinarySearchTree_Clear case1", comparable.IntComparator, []interface{}{}, []interface{}{}, 0},
		{"TestBinarySearchTree_Clear case2", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{}, 0},
		{"TestBinarySearchTree_Clear case3", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{1}, 0},
		{"TestBinarySearchTree_Clear case4", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{1, 6}, 0},
		{"TestBinarySearchTree_Clear case5", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, 0},
		{"TestBinarySearchTree_Clear case6", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{1, 6, 9, 10}, 0},
		{"TestBinarySearchTree_Clear case7", comparable.IntComparator, []interface{}{1, 6, 9}, []interface{}{10}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinarySearchTree(tt.cmp)
			for i := 0; i < len(tt.insert); i++ {
				got.Insert(tt.insert[i])
			}
			for i := 0; i < len(tt.delete); i++ {
				got.Delete(tt.delete[i])
			}
			got.Clear()
			if size := got.Size(); size != tt.want {
				t.Errorf("BinarySearchTree.Size() = %v, want %v", size, tt.want)
			}
		})
	}
}
