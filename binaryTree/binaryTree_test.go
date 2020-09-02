package binaryTree

import (
	"reflect"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestNewBinaryTree case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBinaryTree()
			if got == nil {
				t.Errorf("got nil, want binaryTree")
				return
			}
			if !reflect.DeepEqual(got.root, (*Node)(nil)) {
				t.Errorf("NewBinaryTree() = %v, want %v", got.root, (*Node)(nil))
			}
			if !reflect.DeepEqual(got.size, 0) {
				t.Errorf("NewBinaryTree() = %v, want %v", got.size, 0)
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
		{"TestNewNode case1", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNode(tt.args.newVal)
			if got == nil {
				t.Errorf("got nil, want binaryTree Node")
				return
			}
			if !reflect.DeepEqual(got.value, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.left, (*Node)(nil)) {
				t.Errorf("NewNode() = %v, want %v", got.left, (*Node)(nil))
			}
			if !reflect.DeepEqual(got.right, (*Node)(nil)) {
				t.Errorf("NewNode() = %v, want %v", got.right, (*Node)(nil))
			}
		})
	}
}

func TestBinaryTree_CreateWithLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		size int
	}{
		{"TestBinaryTree_CreateWithLevelOrder case1", []interface{}{}, 0},
		{"TestBinaryTree_CreateWithLevelOrder case2", []interface{}{1}, 1},
		{"TestBinaryTree_CreateWithLevelOrder case3", []interface{}{1, 6, 9}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateWithLevelOrder(tt.args)
			if got == nil {
				t.Errorf("got nil, want binaryTree")
				return
			}
			if !reflect.DeepEqual(got.Size(), tt.size) {
				t.Errorf("BinaryTree.Size() = %v, want %v", got.Size(), tt.size)
			}
		})
	}
}

func TestBinaryTree_PreOrder(t *testing.T) {
	tests := []struct {
		name   string
		create []interface{}
		size   int
		want   interface{}
	}{
		{"TestBinaryTree_PreOrder case1", []interface{}{}, 0, nil},
		{"TestBinaryTree_PreOrder case2", []interface{}{1, 6, 9}, 3, []interface{}{1, 6, 9}},
		{"TestBinaryTree_PreOrder case3", []interface{}{1}, 1, []interface{}{1}},
		{"TestBinaryTree_PreOrder case4", []interface{}{1, 6}, 2, []interface{}{1, 6}},
		{"TestBinaryTree_PreOrder case5", []interface{}{1, nil, 9}, 2, []interface{}{1, 9}},
		{"TestBinaryTree_PreOrder case6", []interface{}{1, 6, 9, 8, 11}, 5, []interface{}{1, 6, 8, 11, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateWithLevelOrder(tt.create)
			if got == nil {
				t.Errorf("got nil, want binaryTree")
				return
			}
			if !reflect.DeepEqual(got.Size(), tt.size) {
				t.Errorf("BinaryTree.Size() = %v, want %v", got.Size(), tt.size)
			}
			if element := got.PreOrder(); !reflect.DeepEqual(element, tt.want) {
				t.Errorf("BinaryTree.PreOrder() = %v, want %v", element, tt.want)
			}
		})
	}
}

func TestBinaryTree_InOrder(t *testing.T) {
	tests := []struct {
		name   string
		create []interface{}
		size   int
		want   interface{}
	}{
		{"TestBinaryTree_InOrder case1", []interface{}{}, 0, nil},
		{"TestBinaryTree_InOrder case2", []interface{}{1, 6, 9}, 3, []interface{}{6, 1, 9}},
		{"TestBinaryTree_InOrder case3", []interface{}{1}, 1, []interface{}{1}},
		{"TestBinaryTree_InOrder case4", []interface{}{1, 6}, 2, []interface{}{6, 1}},
		{"TestBinaryTree_InOrder case5", []interface{}{1, nil, 9}, 2, []interface{}{1, 9}},
		{"TestBinaryTree_InOrder case6", []interface{}{1, 6, 9, 8, 11}, 5, []interface{}{8, 6, 11, 1, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateWithLevelOrder(tt.create)
			if got == nil {
				t.Errorf("got nil, want binaryTree")
				return
			}
			if !reflect.DeepEqual(got.Size(), tt.size) {
				t.Errorf("BinaryTree.Size() = %v, want %v", got.Size(), tt.size)
			}
			if element := got.InOrder(); !reflect.DeepEqual(element, tt.want) {
				t.Errorf("BinaryTree.InOrder() = %v, want %v", element, tt.want)
			}
		})
	}
}

func TestBinaryTree_PostOrder(t *testing.T) {
	tests := []struct {
		name   string
		create []interface{}
		size   int
		want   interface{}
	}{
		{"TestBinaryTree_PostOrder case1", []interface{}{}, 0, nil},
		{"TestBinaryTree_PostOrder case2", []interface{}{1, 6, 9}, 3, []interface{}{6, 9, 1}},
		{"TestBinaryTree_PostOrder case3", []interface{}{1}, 1, []interface{}{1}},
		{"TestBinaryTree_PostOrder case4", []interface{}{1, 6}, 2, []interface{}{6, 1}},
		{"TestBinaryTree_PostOrder case5", []interface{}{1, nil, 9}, 2, []interface{}{9, 1}},
		{"TestBinaryTree_PostOrder case6", []interface{}{1, 6, 9, 8, 11}, 5, []interface{}{8, 11, 6, 9, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateWithLevelOrder(tt.create)
			if got == nil {
				t.Errorf("got nil, want binaryTree")
				return
			}
			if !reflect.DeepEqual(got.Size(), tt.size) {
				t.Errorf("BinaryTree.Size() = %v, want %v", got.Size(), tt.size)
			}
			if element := got.PostOrder(); !reflect.DeepEqual(element, tt.want) {
				t.Errorf("BinaryTree.PostOrder() = %v, want %v", element, tt.want)
			}
		})
	}
}

func TestBinaryTree_LevelOrder(t *testing.T) {
	tests := []struct {
		name   string
		create []interface{}
		size   int
		want   interface{}
	}{
		{"TestBinaryTree_LevelOrder case1", []interface{}{}, 0, nil},
		{"TestBinaryTree_LevelOrder case2", []interface{}{1, 6, 9}, 3, []interface{}{1, 6, 9}},
		{"TestBinaryTree_LevelOrder case3", []interface{}{1}, 1, []interface{}{1}},
		{"TestBinaryTree_LevelOrder case4", []interface{}{1, 6}, 2, []interface{}{1, 6}},
		{"TestBinaryTree_LevelOrder case5", []interface{}{1, nil, 9}, 2, []interface{}{1, 9}},
		{"TestBinaryTree_LevelOrder case6", []interface{}{1, 6, 9, 8, 11}, 5, []interface{}{1, 6, 9, 8, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateWithLevelOrder(tt.create)
			if got == nil {
				t.Errorf("got nil, want binaryTree")
				return
			}
			if !reflect.DeepEqual(got.Size(), tt.size) {
				t.Errorf("BinaryTree.Size() = %v, want %v", got.Size(), tt.size)
			}
			if element := got.LevelOrder(); !reflect.DeepEqual(element, tt.want) {
				t.Errorf("BinaryTree.LevelOrder() = %v, want %v", element, tt.want)
			}
		})
	}
}

func TestBinaryTree_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		create []interface{}
		want   bool
	}{
		{"TestBinaryTree_IsEmpty case1", []interface{}{}, true},
		{"TestBinaryTree_IsEmpty case2", []interface{}{1}, false},
		{"TestBinaryTree_IsEmpty case3", []interface{}{1, 6}, false},
		{"TestBinaryTree_IsEmpty case4", []interface{}{1, nil, 9}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateWithLevelOrder(tt.create)
			if got == nil {
				t.Errorf("got nil, want binaryTree")
				return
			}
			if isEmpty := got.IsEmpty(); isEmpty != tt.want {
				t.Errorf("BinaryTree.IsEmpty() = %v, want %v", isEmpty, tt.want)
			}
		})
	}
}

func TestBinaryTree_Size(t *testing.T) {
	tests := []struct {
		name   string
		create []interface{}
		want   int
	}{
		{"TestBinaryTree_Size case1", []interface{}{}, 0},
		{"TestBinaryTree_Size case2", []interface{}{1, 6}, 2},
		{"TestBinaryTree_Size case3", []interface{}{1, nil, 9}, 2},
		{"TestBinaryTree_Size case4", []interface{}{1, 6, 9}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateWithLevelOrder(tt.create)
			if got == nil {
				t.Errorf("got nil, want binaryTree")
				return
			}
			if !reflect.DeepEqual(got.Size(), tt.want) {
				t.Errorf("BinaryTree.Size() = %v, want %v", got.Size(), tt.want)
			}
		})
	}
}

func TestBinaryTree_Root(t *testing.T) {
	tests := []struct {
		name   string
		create []interface{}
		want   interface{}
	}{
		{"TestBinaryTree_Root case1", []interface{}{}, nil},
		{"TestBinaryTree_Root case2", []interface{}{1}, 1},
		{"TestBinaryTree_Root case3", []interface{}{1, 6}, 1},
		{"TestBinaryTree_Root case4", []interface{}{1, nil, 9}, 1},
		{"TestBinaryTree_Root case5", []interface{}{1, 6, 9}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateWithLevelOrder(tt.create)
			if got == nil {
				t.Errorf("got nil, want binaryTree")
				return
			}
			if got.Size() == 0 {
				if element := got.Root(); !reflect.DeepEqual(element, (*Node)(nil)) {
					t.Errorf("BinaryTree.Root() = %v, want %v", element, nil)
				}
			} else {
				if element := got.Root(); !reflect.DeepEqual(element.value, tt.want) {
					t.Errorf("BinaryTree.Root() = %v, want %v", element, tt.want)
				}
			}
		})
	}
}

func TestBinaryTree_Clear(t *testing.T) {
	tests := []struct {
		name   string
		create []interface{}
		want   int
	}{
		{"TestBinaryTree_Clear case1", []interface{}{}, 0},
		{"TestBinaryTree_Clear case2", []interface{}{1}, 0},
		{"TestBinaryTree_Clear case3", []interface{}{1, 6}, 0},
		{"TestBinaryTree_Clear case4", []interface{}{1, nil, 9}, 0},
		{"TestBinaryTree_Clear case5", []interface{}{1, 6, 9}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateWithLevelOrder(tt.create)
			if got == nil {
				t.Errorf("got nil, want binaryTree")
				return
			}
			got.Clear()
			if !reflect.DeepEqual(got.Size(), tt.want) {
				t.Errorf("BinaryTree.Size() = %v, want %v", got.Size(), tt.want)
			}
		})
	}
}
