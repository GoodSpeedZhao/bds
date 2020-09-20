package trieTree

import (
	"testing"
)

func TestNewTrieTree(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestNewTrieTree case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTrieTree()
			if got == nil {
				t.Errorf("got nil, want TrieTree")
				return
			}
			if got.root == nil {
				t.Errorf("got nil, want TrieTree root.")
			}
			if got.size != 0 {
				t.Errorf("got %v, want %v", got.size, 0)
			}
		})
	}
}

func TestNewTrieNode(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestNewTrieNode case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTrieNode()
			if got == nil {
				t.Errorf("got nil, want TrieNode")
				return
			}
			if got.children == nil {
				t.Errorf("got.children nil, want map.")
			}
			if got.count == nil {
				t.Errorf("got.count nil, want map.")
			}
			if got.isWord == true {
				t.Errorf("got isWord true, want false.")
			}
		})
	}
}

func TestTrieTree_Insert(t *testing.T) {
	tests := []struct {
		name string
		args []string
		size int
	}{
		{"TestTrieTree_Insert case1", []string{"a"}, 1},
		{"TestTrieTree_Insert case2", []string{"abc", "abc"}, 1},
		{"TestTrieTree_Insert case3", []string{"abc", "abcd"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTrieTree()
			got.Insert(tt.args...)
			if got.size != tt.size {
				t.Errorf("got %v, want %v", got.size, tt.size)
			}
		})
	}
}

func TestTrieTree_Delete(t *testing.T) {
	tests := []struct {
		name       string
		insert     []string
		args       []string
		wantResult []bool
		size       int
	}{
		{"TestTrieTree_Delete case1", []string{}, []string{"a"}, []bool{false}, 0},
		{"TestTrieTree_Delete case2", []string{"a"}, []string{"a"}, []bool{true}, 0},
		{"TestTrieTree_Delete case3", []string{"abc"}, []string{"abc"}, []bool{true}, 0},
		{"TestTrieTree_Delete case4", []string{"abc"}, []string{"abcd"}, []bool{false}, 1},
		{"TestTrieTree_Delete case5", []string{"abc", "abcd"}, []string{"abc"}, []bool{true}, 1},
		{"TestTrieTree_Delete case6", []string{"abc", "abcd"}, []string{"abcd"}, []bool{true}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTrieTree()
			got.Insert(tt.insert...)
			for i := 0; i < len(tt.args); i++ {
				if result := got.Delete(tt.args[i]); result != tt.wantResult[i] {
					t.Errorf("TrieTree.Delete() = %v, want %v", result, tt.wantResult[i])
				}
			}
			if got.size != tt.size {
				t.Errorf("got %v, want %v", got.size, tt.size)
			}
		})
	}
}

func TestTrieTree_Contains(t *testing.T) {
	tests := []struct {
		name   string
		insert []string
		delete []string
		args   string
		want   bool
	}{
		{"TestTrieTree_Contains case1", []string{}, []string{}, "a", false},
		{"TestTrieTree_Contains case2", []string{}, []string{"a"}, "a", false},
		{"TestTrieTree_Contains case3", []string{"a"}, []string{"a"}, "a", false},
		{"TestTrieTree_Contains case4", []string{"a"}, []string{}, "a", true},
		{"TestTrieTree_Contains case5", []string{"a", "abc"}, []string{"a"}, "a", false},
		{"TestTrieTree_Contains case6", []string{"a", "abc"}, []string{"a"}, "abc", true},

		{"TestTrieTree_Contains case7", []string{"a", "abc"}, []string{}, ".bc", true},
		{"TestTrieTree_Contains case8", []string{"a", "abc"}, []string{}, "a.c", true},
		{"TestTrieTree_Contains case9", []string{"a", "abc"}, []string{}, "ab.", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTrieTree()
			got.Insert(tt.insert...)
			for i := 0; i < len(tt.delete); i++ {
				got.Delete(tt.delete[i])
			}
			if got := got.Contains(tt.args); got != tt.want {
				t.Errorf("TrieTree.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrieTree_PrefixNumber(t *testing.T) {
	tests := []struct {
		name   string
		insert []string
		delete []string
		args   string
		want   int
	}{
		{"TestTrieTree_PrefixNumber case1", []string{}, []string{}, "", 0},
		{"TestTrieTree_PrefixNumber case2", []string{}, []string{}, "a", 0},
		{"TestTrieTree_PrefixNumber case3", []string{}, []string{}, "abc", 0},
		{"TestTrieTree_PrefixNumber case4", []string{"a"}, []string{}, "a", 1},
		{"TestTrieTree_PrefixNumber case5", []string{"a", "abc"}, []string{}, "a", 2},
		{"TestTrieTree_PrefixNumber case6", []string{"a", "abc"}, []string{}, "ab", 1},
		{"TestTrieTree_PrefixNumber case7", []string{"a", "abc"}, []string{}, "abc", 1},
		{"TestTrieTree_PrefixNumber case8", []string{"a", "abc", "abd"}, []string{}, "a", 3},
		{"TestTrieTree_PrefixNumber case9", []string{"a", "abc", "abd"}, []string{}, "ab", 2},
		{"TestTrieTree_PrefixNumber case10", []string{"a", "abc", "abd"}, []string{}, "abd", 1},
		{"TestTrieTree_PrefixNumber case11", []string{"a", "abc", "abd"}, []string{}, "abdc", 0},
		{"TestTrieTree_PrefixNumber case12", []string{"a", "abc", "abd"}, []string{}, "e", 0},
		{"TestTrieTree_PrefixNumber case13", []string{"a", "abc", "abd"}, []string{"abc"}, "ab", 1},
		{"TestTrieTree_PrefixNumber case14", []string{"a", "abc", "abd"}, []string{"abd"}, "ab", 1},
		{"TestTrieTree_PrefixNumber case15", []string{"a", "abc", "abd"}, []string{"abd"}, "a", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTrieTree()
			got.Insert(tt.insert...)
			for i := 0; i < len(tt.delete); i++ {
				got.Delete(tt.delete[i])
			}
			if sum := got.PrefixNumber(tt.args); sum != tt.want {
				t.Errorf("got %v, want %v", sum, tt.want)
			}
		})
	}
}

func TestTrieTree_Clear(t *testing.T) {
	tests := []struct {
		name   string
		insert []string
		delete []string
		size   int
	}{
		{"TestTrieTree_Clear case1", []string{}, []string{}, 0},
		{"TestTrieTree_Clear case2", []string{}, []string{"a"}, 0},
		{"TestTrieTree_Clear case3", []string{}, []string{"abc"}, 0},
		{"TestTrieTree_Clear case4", []string{"a"}, []string{}, 0},
		{"TestTrieTree_Clear case5", []string{"abc"}, []string{}, 0},
		{"TestTrieTree_Clear case6", []string{"abc"}, []string{"abcd"}, 0},
		{"TestTrieTree_Clear case7", []string{"abc"}, []string{"abc"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTrieTree()
			got.Insert(tt.insert...)
			for i := 0; i < len(tt.delete); i++ {
				got.Delete(tt.delete[i])
			}
			got.Clear()
			if got.size != tt.size {
				t.Errorf("got %v, want %v", got.size, tt.size)
			}
		})
	}
}

func TestTrieTree_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		insert []string
		delete []string
		want   bool
	}{
		{"TestTrieTree_IsEmpty case1", []string{}, []string{}, true},
		{"TestTrieTree_IsEmpty case2", []string{}, []string{"a"}, true},
		{"TestTrieTree_IsEmpty case3", []string{}, []string{"abc"}, true},
		{"TestTrieTree_IsEmpty case4", []string{"abcd"}, []string{"abc"}, false},
		{"TestTrieTree_IsEmpty case5", []string{"abcd"}, []string{"abcd"}, true},
		{"TestTrieTree_IsEmpty case6", []string{"abcd"}, []string{"abcd", "abc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTrieTree()
			got.Insert(tt.insert...)
			for i := 0; i < len(tt.delete); i++ {
				got.Delete(tt.delete[i])
			}
			if isEmpty := got.IsEmpty(); isEmpty != tt.want {
				t.Errorf("TrieTree.IsEmpty() = %v, want %v", isEmpty, tt.want)
			}
		})
	}
}

func TestTrieTree_Size(t *testing.T) {
	tests := []struct {
		name   string
		insert []string
		delete []string
		want   int
	}{
		{"TestTrieTree_Size case1", []string{}, []string{}, 0},
		{"TestTrieTree_Size case2", []string{"a"}, []string{}, 1},
		{"TestTrieTree_Size case3", []string{"a", "a"}, []string{}, 1},
		{"TestTrieTree_Size case4", []string{"abc"}, []string{}, 1},
		{"TestTrieTree_Size case5", []string{"abc", "abc"}, []string{}, 1},
		{"TestTrieTree_Size case6", []string{"abc", "abcd"}, []string{}, 2},
		{"TestTrieTree_Size case7", []string{"abc", "abcd"}, []string{"abc"}, 1},
		{"TestTrieTree_Size case8", []string{"abc", "abcd"}, []string{"abcd"}, 1},
		{"TestTrieTree_Size case9", []string{"abc", "abcd"}, []string{"abc", "abcd"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTrieTree()
			got.Insert(tt.insert...)
			for i := 0; i < len(tt.delete); i++ {
				got.Delete(tt.delete[i])
			}
			if size := got.Size(); size != tt.want {
				t.Errorf("TrieTree.Size() = %v, want %v", size, tt.want)
			}
		})
	}
}
