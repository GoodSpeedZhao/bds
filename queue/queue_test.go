package queue

import (
	"math"
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TestNewQueue case1", args{-1}, 0},
		{"TestNewQueue case2", args{math.MaxUint16 + 100}, math.MaxUint16},
		{"TestNewQueue case3", args{100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewQueue(tt.args.capacity)
			if this == nil {
				t.Errorf("got nil, want queue")
				return
			}
			if this.dataStore == nil {
				t.Errorf("NewQueue() = %v, want queue dataStore", this.dataStore)
			}
			if this.front != 0 {
				t.Errorf("NewQueue() = %v, want front = 0", this.front)
			}
			if this.rear != -1 {
				t.Errorf("NewQueue() = %v, want rear = 0", this.rear)
			}
			if this.size != 0 {
				t.Errorf("NewQueue() = %v, want size = 0", this.size)
			}
			if !reflect.DeepEqual(this.capacity, tt.want) {
				t.Errorf("NewArrayList() = %v, want capacity = %v", this.capacity, tt.want)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	type args struct {
		newVal []interface{}
	}
	tests := []struct {
		name     string
		capacity int
		size     int
		args     args
		wantErr  []bool
	}{
		{"TestQueue_Push case1", -1, 0, args{[]interface{}{1, 6, 9}}, []bool{false, false, false}},
		{"TestQueue_Push case2", math.MaxUint8 + 100, 3, args{[]interface{}{1, 6, 9}}, []bool{true, true, true}},
		{"TestQueue_Push case3", 2, 2, args{[]interface{}{1, 6, 9}}, []bool{true, true, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewQueue(tt.capacity)
			for i := 0; i < len(tt.args.newVal); i++ {
				if err := this.Push(tt.args.newVal[i]); (err != nil) == tt.wantErr[i] {
					t.Errorf("Queue.Push() error = %v, wantErr %v", err, tt.wantErr[i])
				}
			}
			if this.size != tt.size {
				t.Errorf("Queue.Size() error got = %v, want size = %v", this.size, tt.size)
			}
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	type args struct {
		wantValue interface{}
		wantErr   bool
	}
	tests := []struct {
		name     string
		capacity int
		push     []args
		size     int
		args     []args
	}{
		{"TestQueue_Pop case1", -1, []args{}, 0, []args{{nil, false}}},
		{"TestQueue_Pop case2", math.MaxInt16 + 10, []args{}, 0, []args{{nil, false}}},
		{"TestQueue_Pop case3", 3, []args{}, 0, []args{{nil, false}}},
		{"TestQueue_Pop case4", 3, []args{{1, true}}, 0, []args{{1, true}}},
		{"TestQueue_Pop case5", 3, []args{{1, true}, {6, true}}, 0, []args{{1, true}, {6, true}}},
		{"TestQueue_Pop case6", 3, []args{{1, true}, {6, true}, {9, true}}, 0, []args{{1, true}, {6, true}, {9, true}}},
		{"TestQueue_Pop case7", 3, []args{{1, true}, {6, true}, {9, true}, {10, false}}, 0, []args{{1, true}, {6, true}, {9, true}}},
		{"TestQueue_Pop case8", 3, []args{{1, true}, {6, true}, {9, true}, {10, false}}, 0, []args{{1, true}, {6, true}, {9, true}, {nil, false}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewQueue(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				if err := this.Push(tt.push[i].wantValue); (err != nil) == tt.push[i].wantErr {
					t.Errorf("Queue.Push() error = %v, wantErr %v", err, tt.push[i].wantErr)
				}
			}
			for i := 0; i < len(tt.args); i++ {
				gotValue, err := this.Pop()
				if (err != nil) == tt.args[i].wantErr {
					t.Errorf("Queue.Pop() error = %v, wantErr %v", err, tt.args[i].wantErr)
					return
				}
				if !reflect.DeepEqual(gotValue, tt.args[i].wantValue) {
					t.Errorf("Queue.Pop() = %v, want %v", gotValue, tt.args[i].wantValue)
				}
			}
			if this.size != tt.size {
				t.Errorf("Queue.Size() error got = %v, want size = %v", this.size, tt.size)
			}
		})
	}
}

func TestQueue_Front(t *testing.T) {
	type args struct {
		wantValue interface{}
		wantErr   bool
	}
	tests := []struct {
		name      string
		capacity  int
		push      []args
		pop       []args
		size      int
		wantValue interface{}
		wantErr   bool
	}{
		{"TestQueue_Front case1", -1, []args{}, []args{}, 0, nil, false},
		{"TestQueue_Front case2", math.MaxUint16 + 100, []args{}, []args{}, 0, nil, false},
		{"TestQueue_Front case3", -1, []args{{1, false}}, []args{}, 0, nil, false},
		{"TestQueue_Front case4", math.MaxUint16 + 100, []args{{1, true}}, []args{}, 1, 1, true},
		{"TestQueue_Front case5", 3, []args{{1, true}}, []args{}, 1, 1, true},
		{"TestQueue_Front case6", 3, []args{{1, true}, {6, true}}, []args{}, 2, 1, true},
		{"TestQueue_Front case7", 3, []args{{1, true}, {6, true}, {9, true}}, []args{}, 3, 1, true},
		{"TestQueue_Front case8", 3, []args{{1, true}, {6, true}, {9, true}, {12, false}}, []args{}, 3, 1, true},
		{"TestQueue_Front case9", 3, []args{{1, true}, {6, true}, {9, true}}, []args{{1, true}}, 2, 6, true},
		{"TestQueue_Front case10", 3, []args{{1, true}, {6, true}, {9, true}}, []args{{1, true}, {6, true}}, 1, 9, true},
		{"TestQueue_Front case11", 3, []args{{1, true}, {6, true}, {9, true}}, []args{{1, true}, {6, true}, {9, true}}, 0, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewQueue(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				if err := this.Push(tt.push[i].wantValue); (err != nil) == tt.push[i].wantErr {
					t.Errorf("Queue.Push() error = %v, wantErr %v", err, tt.push[i].wantErr)
				}
			}
			for i := 0; i < len(tt.pop); i++ {
				gotValue, err := this.Pop()
				if (err != nil) == tt.pop[i].wantErr {
					t.Errorf("Queue.Pop() error = %v, wantErr %v", err, tt.pop[i].wantErr)
					return
				}
				if !reflect.DeepEqual(gotValue, tt.pop[i].wantValue) {
					t.Errorf("Queue.Pop() = %v, want %v", gotValue, tt.pop[i].wantValue)
				}
			}
			if this.size != tt.size {
				t.Errorf("Queue.Size() error got = %v, want size = %v", this.size, tt.size)
			}
			gotValue, err := this.Front()
			if (err != nil) == tt.wantErr {
				t.Errorf("Queue.Front() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Queue.Front() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestQueue_Rear(t *testing.T) {
	type args struct {
		wantValue interface{}
		wantErr   bool
	}
	tests := []struct {
		name      string
		capacity  int
		push      []args
		pop       []args
		size      int
		wantValue interface{}
		wantErr   bool
	}{
		{"TestQueue_Rear case1", -1, []args{}, []args{}, 0, nil, false},
		{"TestQueue_Rear case2", math.MaxUint16 + 100, []args{}, []args{}, 0, nil, false},
		{"TestQueue_Rear case3", -1, []args{{1, false}}, []args{}, 0, nil, false},
		{"TestQueue_Rear case4", math.MaxUint16 + 100, []args{{1, true}}, []args{}, 1, 1, true},
		{"TestQueue_Rear case5", 3, []args{{1, true}}, []args{}, 1, 1, true},
		{"TestQueue_Rear case6", 3, []args{{1, true}, {6, true}}, []args{}, 2, 6, true},
		{"TestQueue_Rear case7", 3, []args{{1, true}, {6, true}, {9, true}}, []args{}, 3, 9, true},
		{"TestQueue_Rear case8", 3, []args{{1, true}, {6, true}, {9, true}, {12, false}}, []args{}, 3, 9, true},
		{"TestQueue_Rear case9", 3, []args{{1, true}, {6, true}, {9, true}}, []args{{1, true}}, 2, 9, true},
		{"TestQueue_Rear case10", 3, []args{{1, true}, {6, true}, {9, true}}, []args{{1, true}, {6, true}}, 1, 9, true},
		{"TestQueue_Rear case11", 3, []args{{1, true}, {6, true}, {9, true}}, []args{{1, true}, {6, true}, {9, true}}, 0, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewQueue(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				if err := this.Push(tt.push[i].wantValue); (err != nil) == tt.push[i].wantErr {
					t.Errorf("Queue.Push() error = %v, wantErr %v", err, tt.push[i].wantErr)
				}
			}
			for i := 0; i < len(tt.pop); i++ {
				gotValue, err := this.Pop()
				if (err != nil) == tt.pop[i].wantErr {
					t.Errorf("Queue.Pop() error = %v, wantErr %v", err, tt.pop[i].wantErr)
					return
				}
				if !reflect.DeepEqual(gotValue, tt.pop[i].wantValue) {
					t.Errorf("Queue.Pop() = %v, want %v", gotValue, tt.pop[i].wantValue)
				}
			}
			if this.size != tt.size {
				t.Errorf("Queue.Size() error got = %v, want size = %v", this.size, tt.size)
			}
			gotValue, err := this.Rear()
			if (err != nil) == tt.wantErr {
				t.Errorf("Queue.Rear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Queue.Rear() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	type args struct {
		wantValue interface{}
		wantErr   bool
	}
	tests := []struct {
		name     string
		capacity int
		push     []args
		pop      []args
		size     int
		want     bool
	}{
		{"TestQueue_IsEmpty case1", -1, []args{}, []args{}, 0, true},
		{"TestQueue_IsEmpty case2", math.MaxUint16 + 100, []args{}, []args{}, 0, true},
		{"TestQueue_IsEmpty case3", -1, []args{{1, false}}, []args{}, 0, true},
		{"TestQueue_IsEmpty case4", math.MaxUint16 + 100, []args{{1, true}}, []args{}, 1, false},
		{"TestQueue_IsEmpty case5", 3, []args{{1, true}}, []args{}, 1, false},
		{"TestQueue_IsEmpty case6", 3, []args{{1, true}}, []args{{1, true}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewQueue(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				if err := this.Push(tt.push[i].wantValue); (err != nil) == tt.push[i].wantErr {
					t.Errorf("Queue.Push() error = %v, wantErr %v", err, tt.push[i].wantErr)
				}
			}
			for i := 0; i < len(tt.pop); i++ {
				gotValue, err := this.Pop()
				if (err != nil) == tt.pop[i].wantErr {
					t.Errorf("Queue.Pop() error = %v, wantErr %v", err, tt.pop[i].wantErr)
					return
				}
				if !reflect.DeepEqual(gotValue, tt.pop[i].wantValue) {
					t.Errorf("Queue.Pop() = %v, want %v", gotValue, tt.pop[i].wantValue)
				}
			}
			if this.size != tt.size {
				t.Errorf("Queue.Size() error got = %v, want size = %v", this.size, tt.size)
			}
			if got := this.IsEmpty(); got != tt.want {
				t.Errorf("Queue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_IsFull(t *testing.T) {
	type args struct {
		wantValue interface{}
		wantErr   bool
	}
	tests := []struct {
		name     string
		capacity int
		push     []args
		pop      []args
		size     int
		want     bool
	}{
		{"TestQueue_IsFull case1", -1, []args{}, []args{}, 0, true},
		{"TestQueue_IsFull case2", math.MaxUint16 + 100, []args{}, []args{}, 0, false},
		{"TestQueue_IsFull case3", -1, []args{{1, false}}, []args{}, 0, true},
		{"TestQueue_IsFull case4", math.MaxUint16 + 100, []args{{1, true}}, []args{}, 1, false},
		{"TestQueue_IsFull case5", 3, []args{{1, true}}, []args{}, 1, false},
		{"TestQueue_IsFull case6", 3, []args{{1, true}}, []args{{1, true}}, 0, false},
		{"TestQueue_IsFull case7", 3, []args{{1, true}, {6, true}, {9, true}}, []args{}, 3, true},
		{"TestQueue_IsFull case8", 3, []args{{1, true}, {6, true}, {9, true}, {10, false}}, []args{}, 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewQueue(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				if err := this.Push(tt.push[i].wantValue); (err != nil) == tt.push[i].wantErr {
					t.Errorf("Queue.Push() error = %v, wantErr %v", err, tt.push[i].wantErr)
				}
			}
			for i := 0; i < len(tt.pop); i++ {
				gotValue, err := this.Pop()
				if (err != nil) == tt.pop[i].wantErr {
					t.Errorf("Queue.Pop() error = %v, wantErr %v", err, tt.pop[i].wantErr)
					return
				}
				if !reflect.DeepEqual(gotValue, tt.pop[i].wantValue) {
					t.Errorf("Queue.Pop() = %v, want %v", gotValue, tt.pop[i].wantValue)
				}
			}
			if this.size != tt.size {
				t.Errorf("Queue.Size() error got = %v, want size = %v", this.size, tt.size)
			}
			if got := this.IsFull(); got != tt.want {
				t.Errorf("Queue.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Clear(t *testing.T) {
	type args struct {
		wantValue interface{}
		wantErr   bool
	}
	tests := []struct {
		name     string
		capacity int
		push     []args
		pop      []args
		size     int
	}{
		{"TestQueue_Clear case1", -1, []args{}, []args{}, 0},
		{"TestQueue_Clear case2", math.MaxUint16 + 100, []args{}, []args{}, 0},
		{"TestQueue_Clear case3", -1, []args{{1, false}}, []args{}, 0},
		{"TestQueue_Clear case4", math.MaxUint16 + 100, []args{{1, true}}, []args{}, 0},
		{"TestQueue_Clear case5", 3, []args{{1, true}}, []args{}, 0},
		{"TestQueue_Clear case6", 3, []args{{1, true}, {6, true}, {9, true}}, []args{}, 0},
		{"TestQueue_Clear case7", 3, []args{{1, true}, {6, true}, {9, true}, {10, false}}, []args{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewQueue(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				if err := this.Push(tt.push[i].wantValue); (err != nil) == tt.push[i].wantErr {
					t.Errorf("Queue.Push() error = %v, wantErr %v", err, tt.push[i].wantErr)
				}
			}
			for i := 0; i < len(tt.pop); i++ {
				gotValue, err := this.Pop()
				if (err != nil) == tt.pop[i].wantErr {
					t.Errorf("Queue.Pop() error = %v, wantErr %v", err, tt.pop[i].wantErr)
					return
				}
				if !reflect.DeepEqual(gotValue, tt.pop[i].wantValue) {
					t.Errorf("Queue.Pop() = %v, want %v", gotValue, tt.pop[i].wantValue)
				}
			}
			this.Clear()
			if this.size != tt.size {
				t.Errorf("Queue.Size() error got = %v, want size = %v", this.size, tt.size)
			}
		})
	}
}

func TestQueue_Size(t *testing.T) {
	type args struct {
		wantValue interface{}
		wantErr   bool
	}
	tests := []struct {
		name     string
		capacity int
		push     []args
		pop      []args
		want     int
	}{
		{"TestQueue_Size case1", -1, []args{}, []args{}, 0},
		{"TestQueue_Size case2", math.MaxUint16 + 100, []args{}, []args{}, 0},
		{"TestQueue_Size case3", -1, []args{{1, false}}, []args{}, 0},
		{"TestQueue_Size case4", math.MaxUint16 + 100, []args{{1, true}}, []args{}, 1},
		{"TestQueue_Size case5", 3, []args{{1, true}, {6, true}}, []args{}, 2},
		{"TestQueue_Size case6", 3, []args{{1, true}, {6, true}, {9, true}}, []args{}, 3},
		{"TestQueue_Size case7", 3, []args{{1, true}, {6, true}, {9, true}, {10, false}}, []args{}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewQueue(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				if err := this.Push(tt.push[i].wantValue); (err != nil) == tt.push[i].wantErr {
					t.Errorf("Queue.Push() error = %v, wantErr %v", err, tt.push[i].wantErr)
				}
			}
			for i := 0; i < len(tt.pop); i++ {
				gotValue, err := this.Pop()
				if (err != nil) == tt.pop[i].wantErr {
					t.Errorf("Queue.Pop() error = %v, wantErr %v", err, tt.pop[i].wantErr)
					return
				}
				if !reflect.DeepEqual(gotValue, tt.pop[i].wantValue) {
					t.Errorf("Queue.Pop() = %v, want %v", gotValue, tt.pop[i].wantValue)
				}
			}
			if got := this.Size(); got != tt.want {
				t.Errorf("Queue.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
