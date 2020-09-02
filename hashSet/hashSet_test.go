package hashSet

import (
	"testing"
)

func TestNewHashSet(t *testing.T) {
	type args struct {
		newVal []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestNewHashSet case1", args{[]interface{}{}}},
		{"TestNewHashSet case2", args{[]interface{}{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHashSet(tt.args.newVal...)
			if got == nil {
				t.Errorf("got nil, want HashSet")
				return
			}
			if got.dataStore == nil {
				t.Errorf("got nil, want HashSet dataStore")
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	type args struct {
		newVal []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TestSet_Add case1", args{[]interface{}{}}, 0},
		{"TestSet_Add case2", args{[]interface{}{1}}, 1},
		{"TestSet_Add case3", args{[]interface{}{1, 6, 9}}, 3},
		{"TestSet_Add case4", args{[]interface{}{1, 6, 9, 1, 6, 9}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHashSet(tt.args.newVal...)
			got.Add(tt.args.newVal...)
			if size := got.Size(); size != tt.want {
				t.Errorf("size = %v, want = %v", size, tt.want)
			}
		})
	}
}

func TestSet_Delete(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name   string
		newVal []interface{}
		add    []interface{}
		args   args
		want   int
	}{
		{"TestSet_Delete case1", []interface{}{}, []interface{}{}, args{[]interface{}{}}, 0},
		{"TestSet_Delete case2", []interface{}{1, 6, 9}, []interface{}{}, args{[]interface{}{}}, 3},
		{"TestSet_Delete case3", []interface{}{1, 6, 9}, []interface{}{}, args{[]interface{}{10}}, 3},
		{"TestSet_Delete case4", []interface{}{1, 6, 9}, []interface{}{}, args{[]interface{}{1}}, 2},
		{"TestSet_Delete case5", []interface{}{1, 6, 9}, []interface{}{}, args{[]interface{}{1, 6}}, 1},
		{"TestSet_Delete case6", []interface{}{1, 6, 9}, []interface{}{}, args{[]interface{}{1, 6, 9}}, 0},
		{"TestSet_Delete case7", []interface{}{1, 6, 9}, []interface{}{12, 16}, args{[]interface{}{12}}, 4},
		{"TestSet_Delete case8", []interface{}{1, 6, 9}, []interface{}{12, 16}, args{[]interface{}{16, 12}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHashSet(tt.newVal...)
			got.Add(tt.add...)
			got.Delete(tt.args.values...)
			if size := got.Size(); size != tt.want {
				t.Errorf("size = %v, want = %v", size, tt.want)
			}
		})
	}
}

func TestSet_Size(t *testing.T) {
	tests := []struct {
		name   string
		newVal []interface{}
		add    []interface{}
		del    []interface{}
		want   int
	}{
		{"TestSet_Size case1", []interface{}{}, []interface{}{}, []interface{}{}, 0},
		{"TestSet_Size case2", []interface{}{1, 6, 9}, []interface{}{}, []interface{}{}, 3},
		{"TestSet_Size case3", []interface{}{1, 6, 9}, []interface{}{12, 16}, []interface{}{}, 5},
		{"TestSet_Size case4", []interface{}{1, 6, 9}, []interface{}{}, []interface{}{1}, 2},
		{"TestSet_Size case5", []interface{}{1, 6, 9}, []interface{}{}, []interface{}{1, 6}, 1},
		{"TestSet_Size case6", []interface{}{1, 6, 9}, []interface{}{}, []interface{}{1, 6, 9}, 0},
		{"TestSet_Size case7", []interface{}{1, 6, 9}, []interface{}{12, 14}, []interface{}{12}, 4},
		{"TestSet_Size case8", []interface{}{1, 6, 9}, []interface{}{12, 14}, []interface{}{1, 6, 9, 12, 14}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHashSet(tt.newVal...)
			got.Add(tt.add...)
			got.Delete(tt.del...)
			if size := got.Size(); size != tt.want {
				t.Errorf("Set.Size() = %v, want %v", size, tt.want)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	tests := []struct {
		name   string
		newVal []interface{}
		add    []interface{}
		del    []interface{}
	}{
		{"TestSet_Clear case1", []interface{}{}, []interface{}{}, []interface{}{}},
		{"TestSet_Clear case2", []interface{}{1, 6, 9}, []interface{}{}, []interface{}{}},
		{"TestSet_Clear case3", []interface{}{}, []interface{}{1, 6, 9}, []interface{}{}},
		{"TestSet_Clear case4", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{}},
		{"TestSet_Clear case5", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{1, 6, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHashSet(tt.newVal...)
			got.Add(tt.add...)
			got.Delete(tt.del...)
			got.Clear()
			if size := got.Size(); size != 0 {
				t.Errorf("Set.Size() = %v, want %v", size, 0)
			}
		})
	}
}

func TestSet_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		newVal []interface{}
		add    []interface{}
		del    []interface{}
		want   bool
	}{
		{"TestSet_IsEmpty case1", []interface{}{}, []interface{}{}, []interface{}{}, true},
		{"TestSet_IsEmpty case2", []interface{}{1}, []interface{}{}, []interface{}{}, false},
		{"TestSet_IsEmpty case3", []interface{}{}, []interface{}{1}, []interface{}{}, false},
		{"TestSet_IsEmpty case4", []interface{}{}, []interface{}{1}, []interface{}{1}, true},
		{"TestSet_IsEmpty case5", []interface{}{1}, []interface{}{}, []interface{}{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHashSet(tt.newVal...)
			got.Add(tt.add...)
			got.Delete(tt.del...)
			if isEmpty := got.IsEmpty(); isEmpty != tt.want {
				t.Errorf("Set.IsEmpty() = %v, want %v", isEmpty, tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name   string
		newVal []interface{}
		add    []interface{}
		del    []interface{}
		args   args
		want   bool
	}{
		{"TestSet_Contains case1", []interface{}{}, []interface{}{}, []interface{}{}, args{[]interface{}{1}}, false},
		{"TestSet_Contains case2", []interface{}{1, 6, 9}, []interface{}{}, []interface{}{}, args{[]interface{}{6}}, true},
		{"TestSet_Contains case3", []interface{}{}, []interface{}{1, 6, 9}, []interface{}{}, args{[]interface{}{9}}, true},
		{"TestSet_Contains case4", []interface{}{1, 6, 9}, []interface{}{1, 6, 9}, []interface{}{}, args{[]interface{}{1}}, true},
		{"TestSet_Contains case5", []interface{}{1, 6, 9}, []interface{}{}, []interface{}{1}, args{[]interface{}{6}}, true},
		{"TestSet_Contains case6", []interface{}{1, 6, 9}, []interface{}{}, []interface{}{6}, args{[]interface{}{6}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHashSet(tt.newVal...)
			got.Add(tt.add...)
			got.Delete(tt.del...)
			if isContain := got.Contains(tt.args.values...); isContain != tt.want {
				t.Errorf("Set.Contains() = %v, want %v", isContain, tt.want)
			}
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	tests := []struct {
		name     string
		curSet   []interface{}
		otherSet []interface{}
		newSet   interface{}
		size     int
		contains []interface{}
		want     bool
	}{
		{"TestSet_Intersection case1",
			[]interface{}{},
			[]interface{}{},
			nil, 0, []interface{}{}, false},

		{"TestSet_Intersection case2",
			[]interface{}{1},
			[]interface{}{},
			nil, 0, []interface{}{}, false},

		{"TestSet_Intersection case3",
			[]interface{}{},
			[]interface{}{1},
			nil, 0, []interface{}{}, false},

		{"TestSet_Intersection case4",
			[]interface{}{1, 6, 9},
			[]interface{}{1},
			true, 1, []interface{}{1}, true},

		{"TestSet_Intersection case5",
			[]interface{}{1, 6, 9},
			[]interface{}{1, 6},
			true, 2, []interface{}{1, 6}, true},

		{"TestSet_Intersection case6",
			[]interface{}{1, 6, 9},
			[]interface{}{1, 6, 9},
			true, 3, []interface{}{1, 6, 9}, true},

		{"TestSet_Intersection case7",
			[]interface{}{1},
			[]interface{}{1, 6, 9},
			true, 1, []interface{}{1}, true},

		{"TestSet_Intersection case8",
			[]interface{}{1, 6},
			[]interface{}{1, 6, 9},
			true, 2, []interface{}{1, 6}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := NewHashSet(tt.curSet...)
			other := NewHashSet(tt.otherSet...)
			got := self.Intersection(other)
			if tt.newSet == nil {
				if got != nil {
					t.Errorf("got nil, want nil")
				}
				return
			}
			if got.Size() != tt.size {
				t.Errorf("Set.Size() = %v, want %v", got.Size(), tt.size)
			}
			if result := got.Contains(tt.contains...); result != tt.want {
				t.Errorf("Set.Contains() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	tests := []struct {
		name     string
		curSet   []interface{}
		otherSet []interface{}
		newSet   interface{}
		size     int
		contains []interface{}
		want     bool
	}{
		{"TestSet_Union case1",
			[]interface{}{},
			[]interface{}{},
			nil, 0, []interface{}{}, false},

		{"TestSet_Union case2",
			[]interface{}{1},
			[]interface{}{},
			true, 1, []interface{}{1}, true},

		{"TestSet_Union case3",
			[]interface{}{},
			[]interface{}{1},
			true, 1, []interface{}{1}, true},

		{"TestSet_Union case4",
			[]interface{}{1, 6, 9},
			[]interface{}{1},
			true, 3, []interface{}{1, 6, 9}, true},

		{"TestSet_Union case5",
			[]interface{}{1, 6, 9},
			[]interface{}{1, 6},
			true, 3, []interface{}{1, 6, 9}, true},

		{"TestSet_Union case6",
			[]interface{}{1, 6, 9},
			[]interface{}{1, 6, 9},
			true, 3, []interface{}{1, 6, 9}, true},

		{"TestSet_Union case7",
			[]interface{}{1},
			[]interface{}{1, 6, 9},
			true, 3, []interface{}{1, 6, 9}, true},

		{"TestSet_Union case8",
			[]interface{}{1, 6},
			[]interface{}{1, 6, 9},
			true, 3, []interface{}{1, 6, 9}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := NewHashSet(tt.curSet...)
			other := NewHashSet(tt.otherSet...)
			got := self.Union(other)
			if got == nil {
				t.Errorf("got nil, want HashSet")
				return
			}
			if got.Size() != tt.size {
				t.Errorf("Set.Size() = %v, want %v", got.Size(), tt.size)
			}
			if got.Size() != 0 {
				if result := got.Contains(tt.contains...); result != tt.want {
					t.Errorf("Set.Contains() = %v, want %v", result, tt.want)
				}
			}
		})
	}
}

func TestSet_Diff(t *testing.T) {
	tests := []struct {
		name     string
		curSet   []interface{}
		otherSet []interface{}
		newSet   interface{}
		size     int
		contains []interface{}
		want     bool
	}{
		{"TestSet_Diff case1",
			[]interface{}{},
			[]interface{}{},
			nil, 0, []interface{}{}, false},

		{"TestSet_Diff case2",
			[]interface{}{1},
			[]interface{}{},
			true, 1, []interface{}{1}, true},

		{"TestSet_Diff case3",
			[]interface{}{},
			[]interface{}{1},
			nil, 0, []interface{}{}, false},

		{"TestSet_Diff case4",
			[]interface{}{1, 6, 9},
			[]interface{}{1},
			true, 2, []interface{}{6, 9}, true},

		{"TestSet_Diff case5",
			[]interface{}{1, 6, 9},
			[]interface{}{1, 6},
			true, 1, []interface{}{9}, true},

		{"TestSet_Diff case6",
			[]interface{}{1, 6, 9},
			[]interface{}{1, 6, 9},
			nil, 0, []interface{}{}, false},

		{"TestSet_Diff case7",
			[]interface{}{1},
			[]interface{}{1, 6, 9},
			nil, 0, []interface{}{}, false},

		{"TestSet_Diff case8",
			[]interface{}{1, 6},
			[]interface{}{1, 6, 9},
			nil, 0, []interface{}{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := NewHashSet(tt.curSet...)
			other := NewHashSet(tt.otherSet...)
			got := self.Diff(other)
			if got == nil {
				t.Errorf("got nil, want nil")
				return
			}
			if got.Size() != tt.size {
				t.Errorf("Set.Size() = %v, want %v", got.Size(), tt.size)
			}
			if got.Size() != 0 {
				if result := got.Contains(tt.contains...); result != tt.want {
					t.Errorf("Set.Contains() = %v, want %v", result, tt.want)
				}
			}
		})
	}
}
