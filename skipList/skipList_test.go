package skipList

import (
	"math"
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	type args struct {
		newKey int
		newVal interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantKey int
		wantVal interface{}
	}{
		{"TestNewNode case1", args{1, 1}, 1, 1},
		{"TestNewNode case2", args{10, 10}, 10, 10},
		{"TestNewNode case3", args{6, 60}, 6, 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNode(tt.args.newKey, tt.args.newVal)
			if got == nil {
				t.Errorf("got nil, want Node")
				return
			}
			if !reflect.DeepEqual(got.key, tt.wantKey) {
				t.Errorf("newKey = %v, wantKey %v", got.key, tt.wantKey)
			}
			if !reflect.DeepEqual(got.value, tt.wantVal) {
				t.Errorf("newVal = %v, wantVal %v", got.value, tt.wantVal)
			}
		})
	}
}

func TestNewSkipList(t *testing.T) {
	tests := []struct {
		name        string
		wantHeadKey int
		wantTailKey int
		wantSize    int
		wantLevel   int
	}{
		{"TestNewSkipList case1", math.MinInt64, math.MaxInt64, 0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSkipList()
			if got == nil {
				t.Errorf("got nil, want skipList")
				return
			}
			if !reflect.DeepEqual(got.head.key, tt.wantHeadKey) {
				t.Errorf("head.key = %v, wantHeadKey %v", got.head.key, tt.wantHeadKey)
			}
			if !reflect.DeepEqual(got.tail.key, tt.wantTailKey) {
				t.Errorf("tail.key = %v, wantTailKey %v", got.tail.key, tt.wantTailKey)
			}
			if !reflect.DeepEqual(got.size, tt.wantSize) {
				t.Errorf("got.size = %v, wantSize %v", got.size, tt.wantSize)
			}
			if !reflect.DeepEqual(got.level, tt.wantLevel) {
				t.Errorf("got.level = %v, wantLevel %v", got.level, tt.wantLevel)
			}
		})
	}
}

func TestSkipList_Search(t *testing.T) {
	type args struct {
		newKey int
		newVal interface{}
	}
	tests := []struct {
		name    string
		key     int
		add     []args
		erase   []int
		want    interface{}
		wantErr bool
	}{
		{"TestSkipList_Search case1", 1, []args{}, []int{}, 1, false},
		{"TestSkipList_Search case2", 1, []args{{1, 1}}, []int{}, 1, true},
		{"TestSkipList_Search case3", 9, []args{{1, 1}, {3, 3}, {6, 6}, {9, 9}}, []int{}, 9, true},
		{"TestSkipList_Search case4", 5, []args{{1, 1}, {3, 3}, {5, 5}, {6, 6}, {9, 9}}, []int{}, 5, true},
		{"TestSkipList_Search case5", 5, []args{{1, 1}, {3, 3}, {5, 5}, {6, 6}, {9, 9}}, []int{1}, 5, true},
		{"TestSkipList_Search case6", 5, []args{{1, 1}, {3, 3}, {5, 5}, {6, 6}, {9, 9}}, []int{1, 3}, 5, true},
		{"TestSkipList_Search case7", 5, []args{{1, 1}, {3, 3}, {5, 5}, {6, 6}, {9, 9}}, []int{6, 9}, 5, true},
		{"TestSkipList_Search case8", 5, []args{{1, 1}, {3, 3}, {5, 5}, {6, 6}, {9, 9}}, []int{5, 6, 9}, 5, false},
		{"TestSkipList_Search case9", 5, []args{{1, 1}, {3, 3}, {5, 5}, {6, 6}, {9, 9}}, []int{1, 3, 5}, 5, false},
		{"TestSkipList_Search case10", 5, []args{{1, 1}, {3, 3}, {5, 5}, {6, 6}, {9, 9}}, []int{1, 3, 5, 6, 9}, 5, false},
		{"TestSkipList_Search case11", 5, []args{{1, 1}, {3, 3}, {6, 6}, {9, 9}}, []int{}, 5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSkipList()
			for i := 0; i < len(tt.add); i++ {
				got.Add(tt.add[i].newKey, tt.add[i].newVal)
			}
			for i := 0; i < len(tt.erase); i++ {
				got.Erase(tt.erase[i])
			}
			element, err := got.Search(tt.key)
			if (err != nil) == tt.wantErr {
				t.Errorf("SkipList.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil) && !reflect.DeepEqual(element.value, tt.want) {
				t.Errorf("SkipList.Search() = %v, want %v", element.value, tt.want)
			}
		})
	}
}

func TestSkipList_Add(t *testing.T) {
	type args struct {
		key    int
		newVal interface{}
	}
	tests := []struct {
		name  string
		args  []args
		erase []int
		add   []args
		want  int
	}{
		{
			"TestSkipList_Add case1",
			[]args{{1, 1}, {3, 3}, {6, 6}},
			[]int{},
			[]args{},
			3,
		},
		{
			"TestSkipList_Add case2",
			[]args{{1, 1}, {3, 3}, {6, 6}, {9, 9}},
			[]int{1},
			[]args{},
			3,
		},
		{
			"TestSkipList_Add case3",
			[]args{{1, 1}, {3, 3}, {6, 6}, {9, 9}},
			[]int{9},
			[]args{},
			3,
		},
		{
			"TestSkipList_Add case4",
			[]args{{1, 1}, {3, 3}, {6, 6}, {9, 9}},
			[]int{3},
			[]args{},
			3,
		},
		{
			"TestSkipList_Add case5",
			[]args{{1, 1}, {3, 3}, {6, 6}, {9, 9}},
			[]int{1, 3, 6, 9},
			[]args{},
			0,
		},
		{
			"TestSkipList_Add case6",
			[]args{{1, 1}, {3, 3}, {6, 6}, {9, 9}},
			[]int{1, 3, 6, 9},
			[]args{{1, 1}, {9, 9}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSkipList()
			for i := 0; i < len(tt.args); i++ {
				got.Add(tt.args[i].key, tt.args[i].newVal)
			}
			for i := 0; i < len(tt.erase); i++ {
				got.Erase(tt.erase[i])
			}
			for i := 0; i < len(tt.add); i++ {
				got.Add(tt.add[i].key, tt.add[i].newVal)
			}
			if !reflect.DeepEqual(got.size, tt.want) {
				t.Errorf("SkipList.Add() = %v, want %v", got.size, tt.want)
			}
		})
	}
}

func TestSkipList_Erase(t *testing.T) {
	type args struct {
		newKey int
		newVal interface{}
	}
	tests := []struct {
		name       string
		add        []args
		erase      []int
		wantResult []bool
	}{
		{"TestSkipList_Erase case1", []args{}, []int{1}, []bool{false}},
		{"TestSkipList_Erase case2", []args{{3, 3}}, []int{1}, []bool{false}},
		{"TestSkipList_Erase case3", []args{{3, 3}, {6, 6}, {9, 9}}, []int{1}, []bool{false}},
		{"TestSkipList_Erase case4", []args{{3, 3}, {6, 6}, {9, 9}}, []int{3}, []bool{true}},
		{"TestSkipList_Erase case5", []args{{3, 3}, {6, 6}, {9, 9}}, []int{3, 6}, []bool{true, true}},
		{"TestSkipList_Erase case6", []args{{3, 3}, {6, 6}, {9, 9}}, []int{3, 6, 9}, []bool{true, true, true}},
		{"TestSkipList_Erase case7", []args{{3, 3}, {6, 6}, {9, 9}}, []int{3, 6, 9, 10}, []bool{true, true, true, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSkipList()
			for i := 0; i < len(tt.add); i++ {
				got.Add(tt.add[i].newKey, tt.add[i].newVal)
			}
			for i := 0; i < len(tt.erase); i++ {
				if gotResult := got.Erase(tt.erase[i]); gotResult != tt.wantResult[i] {
					t.Errorf("SkipList.Erase() = %v, want %v", gotResult, tt.wantResult[i])
				}
			}
		})
	}
}

func TestSkipList_Get(t *testing.T) {
	type args struct {
		newKey int
		newVal interface{}
	}
	tests := []struct {
		name    string
		add     []args
		erase   []int
		getKey  int
		want    interface{}
		wantErr bool
	}{
		{"TestSkipList_Get case1", []args{}, []int{}, 1, (*Node)(nil), false},
		{"TestSkipList_Get case2", []args{{3, 3}, {6, 6}}, []int{}, 1, (*Node)(nil), false},
		{"TestSkipList_Get case3", []args{{1, 1}, {3, 3}, {6, 6}}, []int{}, 1, interface{}(1), true},
		{"TestSkipList_Get case4", []args{{1, 1}, {3, 3}, {6, 6}}, []int{6}, 1, interface{}(1), true},
		{"TestSkipList_Get case5", []args{{1, 1}, {3, 3}, {6, 6}}, []int{3}, 1, interface{}(1), true},
		{"TestSkipList_Get case6", []args{{1, 1}, {3, 3}, {6, 6}}, []int{1}, 1, (*Node)(nil), false},
		{"TestSkipList_Get case7", []args{{1, 1}, {3, 3}, {6, 6}}, []int{1, 3, 6}, 1, (*Node)(nil), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSkipList()
			for i := 0; i < len(tt.add); i++ {
				got.Add(tt.add[i].newKey, tt.add[i].newVal)
			}
			for i := 0; i < len(tt.erase); i++ {
				got.Erase(tt.erase[i])
			}
			element, err := got.Get(tt.getKey)
			if (err != nil) == tt.wantErr {
				t.Errorf("SkipList.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (element != nil) && !reflect.DeepEqual(element.value, tt.want) {
				t.Errorf("SkipList.Get() = %v, want %v", element.value, tt.want)
			}
		})
	}
}

func TestSkipList_GetSize(t *testing.T) {
	type args struct {
		newKey int
		newVal interface{}
	}
	tests := []struct {
		name  string
		add   []args
		erase []int
		want  int
	}{
		{"TestSkipList_GetSize case1", []args{}, []int{}, 0},
		{"TestSkipList_GetSize case2", []args{{1, 1}}, []int{}, 1},
		{"TestSkipList_GetSize case3", []args{{1, 1}, {6, 6}, {9, 9}}, []int{}, 3},
		{"TestSkipList_GetSize case4", []args{{1, 1}, {6, 6}, {9, 9}}, []int{1}, 2},
		{"TestSkipList_GetSize case5", []args{{1, 1}, {6, 6}, {9, 9}}, []int{1, 6}, 1},
		{"TestSkipList_GetSize case6", []args{{1, 1}, {6, 6}, {9, 9}}, []int{1, 6, 9}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSkipList()
			for i := 0; i < len(tt.add); i++ {
				got.Add(tt.add[i].newKey, tt.add[i].newVal)
			}
			for i := 0; i < len(tt.erase); i++ {
				got.Erase(tt.erase[i])
			}
			if size := got.GetSize(); size != tt.want {
				t.Errorf("SkipList.GetSize() = %v, want %v", size, tt.want)
			}
		})
	}
}
