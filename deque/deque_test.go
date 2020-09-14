package deque

import (
	"math"
	"reflect"
	"testing"
)

func TestNewDeque(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestNewDeque case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDeque()
			if got == nil {
				t.Errorf("got nil, want Deque")
				return
			}
			if got.dataStore == nil {
				t.Errorf("got nil, want Deque dataStore")
			}
			if got.capacity != math.MaxUint32 {
				t.Errorf("got %v, want %v", got.capacity, math.MaxUint32)
			}
		})
	}
}

func TestNewDequeWithCapacity(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TestNewDequeWithCapacity case1", args{-10}, 0},
		{"TestNewDequeWithCapacity case2", args{10}, 10},
		{"TestNewDequeWithCapacity case3", args{math.MaxUint32 + 10}, math.MaxUint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDequeWithCapacity(tt.args.capacity)
			if got == nil {
				t.Errorf("got nil, want Deque")
				return
			}
			if got.dataStore == nil {
				t.Errorf("got nil, want Deque dataStore")
			}
			if got.capacity != tt.want {
				t.Errorf("got %v, want %v", got.capacity, tt.want)
			}
		})
	}
}

func TestDeque_Append(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		size     int
		args     []interface{}
		wantErr  []bool
	}{
		{"TestDeque_Append case1", -10, 0, []interface{}{1, 6, 9}, []bool{false, false, false}},
		{"TestDeque_Append case2", math.MaxUint32 + 100, 3, []interface{}{1, 6, 9}, []bool{true, true, true}},
		{"TestDeque_Append case3", 2, 2, []interface{}{1, 6, 9}, []bool{true, true, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDequeWithCapacity(tt.capacity)
			for i := 0; i < len(tt.args); i++ {
				if err := got.Append(tt.args[i]); err != tt.wantErr[i] {
					t.Errorf("Deque.Append() error = %v, wantErr %v", err, tt.wantErr[i])
				}
			}
			if got.Size() != tt.size {
				t.Errorf("Deque.Size() error got = %v, want size = %v", got.Size(), tt.size)
			}
		})
	}
}

func TestDeque_Prepend(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		size     int
		args     []interface{}
		wantErr  []bool
	}{
		{"TestDeque_Prepend case1", -10, 0, []interface{}{1, 6, 9}, []bool{false, false, false}},
		{"TestDeque_Prepend case2", math.MaxUint32 + 100, 3, []interface{}{1, 6, 9}, []bool{true, true, true}},
		{"TestDeque_Prepend case3", 2, 2, []interface{}{1, 6, 9}, []bool{true, true, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDequeWithCapacity(tt.capacity)
			for i := 0; i < len(tt.args); i++ {
				if err := got.Prepend(tt.args[i]); err != tt.wantErr[i] {
					t.Errorf("Deque.Prepend() error = %v, wantErr %v", err, tt.wantErr[i])
				}
			}
			if got.Size() != tt.size {
				t.Errorf("Deque.Size() error got = %v, want size = %v", got.Size(), tt.size)
			}
		})
	}
}

func TestDeque_Shift(t *testing.T) {
	tests := []struct {
		name      string
		capacity  int
		append    []interface{}
		args      int
		wantErr   []bool
		wantValue []interface{}
		size      int
	}{
		{"TestDeque_Shift case1", -10, []interface{}{}, 1, []bool{false}, []interface{}{}, 0},
		{"TestDeque_Shift case2", math.MaxUint32 + 100, []interface{}{1, 6, 9}, 3, []bool{true, true, true}, []interface{}{1, 6, 9}, 0},
		{"TestDeque_Shift case3", 2, []interface{}{1, 6, 9}, 2, []bool{true, true}, []interface{}{1, 6}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDequeWithCapacity(tt.capacity)
			for i := 0; i < len(tt.append); i++ {
				got.Append(tt.append[i])
			}
			for i := 0; i < tt.args; i++ {
				value, err := got.Shift()
				if (err != nil) == tt.wantErr[i] {
					t.Errorf("Deque.Shift() error = %v, wantErr %v", err, tt.wantErr[i])
					return
				}
				if err == nil && !reflect.DeepEqual(value, tt.wantValue[i]) {
					t.Errorf("Deque.Shift() = %v, want %v", value, tt.wantValue[i])
				}
			}
			if got.Size() != tt.size {
				t.Errorf("Deque.Size() error got = %v, want size = %v", got.Size(), tt.size)
			}
		})
	}
}

func TestDeque_Pop(t *testing.T) {
	tests := []struct {
		name      string
		capacity  int
		append    []interface{}
		args      int
		wantErr   []bool
		wantValue []interface{}
		size      int
	}{
		{"TestDeque_Pop case1", -10, []interface{}{}, 1, []bool{false}, []interface{}{}, 0},
		{"TestDeque_Pop case2", math.MaxUint32 + 100, []interface{}{1, 6, 9}, 3, []bool{true, true, true}, []interface{}{9, 6, 1}, 0},
		{"TestDeque_Pop case3", 2, []interface{}{1, 6, 9}, 2, []bool{true, true}, []interface{}{6, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDequeWithCapacity(tt.capacity)
			for i := 0; i < len(tt.append); i++ {
				got.Append(tt.append[i])
			}
			for i := 0; i < tt.args; i++ {
				value, err := got.Pop()
				if (err != nil) == tt.wantErr[i] {
					t.Errorf("Deque.Pop() error = %v, wantErr %v", err, tt.wantErr[i])
					return
				}
				if err == nil && !reflect.DeepEqual(value, tt.wantValue[i]) {
					t.Errorf("Deque.Pop() = %v, want %v", value, tt.wantValue[i])
				}
			}
			if got.Size() != tt.size {
				t.Errorf("Deque.Size() error got = %v, want size = %v", got.Size(), tt.size)
			}
		})
	}
}

func TestDeque_Front(t *testing.T) {
	tests := []struct {
		name      string
		capacity  int
		append    []interface{}
		prepend   []interface{}
		pop       int
		args      int
		wantErr   []bool
		wantValue []interface{}
	}{
		{"TestDeque_Front case1", -10, []interface{}{}, []interface{}{}, 0, 1, []bool{false}, nil},
		{"TestDeque_Front case2", math.MaxUint32 + 100, []interface{}{1, 6, 9}, []interface{}{}, 0, 1, []bool{true}, []interface{}{1}},
		{"TestDeque_Front case3", math.MaxUint32 + 100, []interface{}{}, []interface{}{1, 6, 9}, 0, 1, []bool{true}, []interface{}{9}},
		{"TestDeque_Front case4", 2, []interface{}{1, 6, 9}, []interface{}{}, 0, 1, []bool{true}, []interface{}{1}},
		{"TestDeque_Front case5", 2, []interface{}{}, []interface{}{1, 6, 9}, 0, 1, []bool{true}, []interface{}{6}},
		{"TestDeque_Front case6", 6, []interface{}{1, 6, 9}, []interface{}{}, 1, 1, []bool{true}, []interface{}{1}},
		{"TestDeque_Front case7", 6, []interface{}{1, 6, 9}, []interface{}{}, 2, 1, []bool{true}, []interface{}{1}},
		{"TestDeque_Front case8", 6, []interface{}{1, 6, 9}, []interface{}{}, 3, 1, []bool{false}, nil},
		{"TestDeque_Front case9", 6, []interface{}{}, []interface{}{1, 6, 9}, 1, 1, []bool{true}, []interface{}{9}},
		{"TestDeque_Front case10", 6, []interface{}{}, []interface{}{1, 6, 9}, 2, 1, []bool{true}, []interface{}{9}},
		{"TestDeque_Front case11", 6, []interface{}{}, []interface{}{1, 6, 9}, 3, 1, []bool{false}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDequeWithCapacity(tt.capacity)
			for i := 0; i < len(tt.append); i++ {
				got.Append(tt.append[i])
			}
			for i := 0; i < len(tt.prepend); i++ {
				got.Prepend(tt.prepend[i])
			}
			for i := 0; i < tt.pop; i++ {
				got.Pop()
			}
			for i := 0; i < tt.args; i++ {
				value, err := got.Front()
				if (err != nil) == tt.wantErr[i] {
					t.Errorf("Deque.Front() error = %v, wantErr %v", err, tt.wantErr[i])
					return
				}
				if err == nil && !reflect.DeepEqual(value, tt.wantValue[i]) {
					t.Errorf("Deque.Front() = %v, want %v", value, tt.wantValue[i])
				}
			}
		})
	}
}

func TestDeque_Rear(t *testing.T) {
	tests := []struct {
		name      string
		capacity  int
		append    []interface{}
		prepend   []interface{}
		pop       int
		args      int
		wantErr   []bool
		wantValue []interface{}
	}{
		{"TestDeque_Rear case1", -10, []interface{}{}, []interface{}{}, 0, 1, []bool{false}, nil},
		{"TestDeque_Rear case2", math.MaxUint32 + 100, []interface{}{1, 6, 9}, []interface{}{}, 0, 1, []bool{true}, []interface{}{9}},
		{"TestDeque_Rear case3", math.MaxUint32 + 100, []interface{}{}, []interface{}{1, 6, 9}, 0, 1, []bool{true}, []interface{}{1}},
		{"TestDeque_Rear case4", 2, []interface{}{1, 6, 9}, []interface{}{}, 0, 1, []bool{true}, []interface{}{6}},
		{"TestDeque_Rear case5", 2, []interface{}{}, []interface{}{1, 6, 9}, 0, 1, []bool{true}, []interface{}{1}},
		{"TestDeque_Rear case6", 6, []interface{}{1, 6, 9}, []interface{}{}, 1, 1, []bool{true}, []interface{}{6}},
		{"TestDeque_Rear case7", 6, []interface{}{1, 6, 9}, []interface{}{}, 2, 1, []bool{true}, []interface{}{1}},
		{"TestDeque_Rear case8", 6, []interface{}{1, 6, 9}, []interface{}{}, 3, 1, []bool{false}, nil},
		{"TestDeque_Rear case9", 6, []interface{}{}, []interface{}{1, 6, 9}, 1, 1, []bool{true}, []interface{}{6}},
		{"TestDeque_Rear case10", 6, []interface{}{}, []interface{}{1, 6, 9}, 2, 1, []bool{true}, []interface{}{9}},
		{"TestDeque_Rear case11", 6, []interface{}{}, []interface{}{1, 6, 9}, 3, 1, []bool{false}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDequeWithCapacity(tt.capacity)
			for i := 0; i < len(tt.append); i++ {
				got.Append(tt.append[i])
			}
			for i := 0; i < len(tt.prepend); i++ {
				got.Prepend(tt.prepend[i])
			}
			for i := 0; i < tt.pop; i++ {
				got.Pop()
			}
			for i := 0; i < tt.args; i++ {
				value, err := got.Rear()
				if (err != nil) == tt.wantErr[i] {
					t.Errorf("Deque.Front() error = %v, wantErr %v", err, tt.wantErr[i])
					return
				}
				if err == nil && !reflect.DeepEqual(value, tt.wantValue[i]) {
					t.Errorf("Deque.Front() = %v, want %v", value, tt.wantValue[i])
				}
			}
		})
	}
}

func TestDeque_Size(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		append   []interface{}
		prepend  []interface{}
		pop      int
		want     int
	}{
		{"TestDeque_Size case1", -10, []interface{}{}, []interface{}{}, 0, 0},
		{"TestDeque_Size case2", math.MaxUint32 + 100, []interface{}{}, []interface{}{}, 0, 0},
		{"TestDeque_Size case3", 2, []interface{}{}, []interface{}{}, 0, 0},
		{"TestDeque_Size case4", -10, []interface{}{1}, []interface{}{1}, 0, 0},
		{"TestDeque_Size case5", math.MaxUint32 + 100, []interface{}{1}, []interface{}{2}, 0, 2},
		{"TestDeque_Size case6", math.MaxUint32 + 100, []interface{}{1}, []interface{}{2}, 1, 1},
		{"TestDeque_Size case7", math.MaxUint32 + 100, []interface{}{1}, []interface{}{2}, 2, 0},
		{"TestDeque_Size case8", 2, []interface{}{1}, []interface{}{2}, 0, 2},
		{"TestDeque_Size case9", 2, []interface{}{1}, []interface{}{2}, 1, 1},
		{"TestDeque_Size case10", 2, []interface{}{1}, []interface{}{2}, 2, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDequeWithCapacity(tt.capacity)
			for i := 0; i < len(tt.append); i++ {
				got.Append(tt.append[i])
			}
			for i := 0; i < len(tt.prepend); i++ {
				got.Prepend(tt.prepend[i])
			}
			for i := 0; i < tt.pop; i++ {
				got.Pop()
			}
			if got.Size() != tt.want {
				t.Errorf("Deque.Size() error got = %v, want size = %v", got.Size(), tt.want)
			}
		})
	}
}

func TestDeque_Capacity(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		want     int
	}{
		{"TestDeque_Capacity case1", -10, 0},
		{"TestDeque_Capacity case2", math.MaxUint32 + 100, math.MaxUint32},
		{"TestDeque_Capacity case3", 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDequeWithCapacity(tt.capacity)
			if cap := got.Capacity(); cap != tt.want {
				t.Errorf("Deque.Capacity() = %v, want %v", cap, tt.want)
			}
		})
	}
}

func TestDeque_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		append   []interface{}
		pop      int
		want     bool
	}{
		{"TestDeque_IsEmpty case1", -10, []interface{}{}, 0, true},
		{"TestDeque_IsEmpty case2", math.MaxUint32 + 100, []interface{}{}, 0, true},
		{"TestDeque_IsEmpty case3", 2, []interface{}{}, 0, true},
		{"TestDeque_IsEmpty case4", -10, []interface{}{1}, 0, true},
		{"TestDeque_IsEmpty case5", math.MaxUint32 + 100, []interface{}{1, 2, 3}, 0, false},
		{"TestDeque_IsEmpty case6", math.MaxUint32 + 100, []interface{}{1, 2, 3}, 2, false},
		{"TestDeque_IsEmpty case7", math.MaxUint32 + 100, []interface{}{1, 2, 3}, 3, true},
		{"TestDeque_IsEmpty case8", 2, []interface{}{}, 0, true},
		{"TestDeque_IsEmpty case9", 2, []interface{}{1, 2, 3}, 0, false},
		{"TestDeque_IsEmpty case10", 2, []interface{}{1, 2, 3}, 2, true},
		{"TestDeque_IsEmpty case11", 2, []interface{}{1, 2, 3}, 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDequeWithCapacity(tt.capacity)
			for i := 0; i < len(tt.append); i++ {
				got.Append(tt.append[i])
			}
			for i := 0; i < tt.pop; i++ {
				got.Pop()
			}
			if isEmpty := got.IsEmpty(); isEmpty != tt.want {
				t.Errorf("Deque.IsEmpty() = %v, want %v", isEmpty, tt.want)
			}
		})
	}
}

func TestDeque_IsFull(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		append   []interface{}
		pop      int
		want     bool
	}{
		{"TestDeque_IsFull case1", -10, []interface{}{}, 0, true},
		{"TestDeque_IsFull case2", math.MaxUint32 + 100, []interface{}{}, 0, false},
		{"TestDeque_IsFull case3", 2, []interface{}{}, 0, false},
		{"TestDeque_IsFull case4", -10, []interface{}{1}, 0, true},
		{"TestDeque_IsFull case5", math.MaxUint32 + 100, []interface{}{1, 2, 3}, 0, false},
		{"TestDeque_IsFull case6", math.MaxUint32 + 100, []interface{}{1, 2, 3}, 2, false},
		{"TestDeque_IsFull case7", math.MaxUint32 + 100, []interface{}{1, 2, 3}, 3, false},
		{"TestDeque_IsFull case8", 2, []interface{}{}, 0, false},
		{"TestDeque_IsFull case9", 2, []interface{}{1, 2, 3}, 0, true},
		{"TestDeque_IsFull case10", 2, []interface{}{1, 2, 3}, 2, false},
		{"TestDeque_IsFull case11", 2, []interface{}{1, 2, 3}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDequeWithCapacity(tt.capacity)
			for i := 0; i < len(tt.append); i++ {
				got.Append(tt.append[i])
			}
			for i := 0; i < tt.pop; i++ {
				got.Pop()
			}
			if isFull := got.IsFull(); isFull != tt.want {
				t.Errorf("Deque.IsFull() = %v, want %v", isFull, tt.want)
			}
		})
	}
}
