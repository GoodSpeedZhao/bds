package stack

import (
	"math"
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestNewStack case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewStack()
			if this == nil {
				t.Errorf("got nil, want Stack")
				return
			}
			if this.dataStore == nil {
				t.Errorf("got nil, want Stack dataStore")
			}
			if this.capacity != math.MaxUint32 {
				t.Errorf("got %v, want %v", this.capacity, math.MaxUint32)
			}
		})
	}
}

func TestNewStackWithCapacity(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TestNewStackWithCapacity case1", args{-10}, 0},
		{"TestNewStackWithCapacity case2", args{10}, 10},
		{"TestNewStackWithCapacity case3", args{math.MaxUint32 + 10}, math.MaxUint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewStackWithCapacity(tt.args.capacity)
			if this == nil {
				t.Errorf("got nil, want Stack")
				return
			}
			if this.dataStore == nil {
				t.Errorf("got nil, want Stack dataStore")
			}
			if this.capacity != tt.want {
				t.Errorf("got %v, want %v", this.capacity, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name     string
		capacity int
		args     args
		wantErr  bool
		want1    interface{}
		wantErr1 bool
	}{
		{"TestStack_Push case1", 0, args{1}, false, nil, false},
		{"TestStack_Push case2", 3, args{1}, true, 1, true},
		{"TestStack_Push case3", 1, args{6}, true, 6, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewStackWithCapacity(tt.capacity)
			if err := this.Push(tt.args.value); (err != nil) == tt.wantErr {
				t.Errorf("Stack.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, err := this.Peek()
			if (err != nil) == tt.wantErr1 {
				t.Errorf("Stack.Peek() error = %v, wantErr %v", err, tt.wantErr1)
				return
			}
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Stack.Peek() = %v, want %v", got, tt.want1)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name      string
		capacity  int
		push      []interface{}
		pop       int
		wantValue interface{}
		wantErr   bool
	}{
		{"TestStack_Pop case1", 0, []interface{}{}, 0, nil, false},
		{"TestStack_Pop case2", 6, []interface{}{}, 0, nil, false},
		{"TestStack_Pop case3", 6, []interface{}{1, 3, 6}, 0, 6, true},
		{"TestStack_Pop case4", 6, []interface{}{1, 3, 6}, 1, 3, true},
		{"TestStack_Pop case5", 6, []interface{}{1, 3, 6}, 2, 1, true},
		{"TestStack_Pop case6", 6, []interface{}{1, 3, 6}, 3, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewStackWithCapacity(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				this.Push(tt.push[i])
			}
			for i := 0; i < tt.pop; i++ {
				this.Pop()
			}
			gotValue, err := this.Pop()
			if (err != nil) == tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Stack.Pop() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		push     []interface{}
		pop      int
		want     interface{}
		wantErr  bool
	}{
		{"TestStack_Peek case1", 0, []interface{}{}, 0, nil, false},
		{"TestStack_Peek case2", 3, []interface{}{}, 0, nil, false},
		{"TestStack_Peek case3", 3, []interface{}{1, 2, 3}, 0, 3, true},
		{"TestStack_Peek case4", 3, []interface{}{1, 2, 3}, 1, 2, true},
		{"TestStack_Peek case5", 3, []interface{}{1, 2, 3}, 2, 1, true},
		{"TestStack_Peek case6", 3, []interface{}{1, 2, 3}, 3, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewStackWithCapacity(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				this.Push(tt.push[i])
			}
			for i := 0; i < tt.pop; i++ {
				this.Pop()
			}
			got, err := this.Peek()
			if (err != nil) == tt.wantErr {
				t.Errorf("Stack.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		push     []interface{}
		pop      int
		want     bool
	}{
		{"TestStack_IsEmpty case1", 0, []interface{}{}, 0, true},
		{"TestStack_IsEmpty case2", 3, []interface{}{}, 0, true},
		{"TestStack_IsEmpty case3", 3, []interface{}{1, 2, 3}, 0, false},
		{"TestStack_IsEmpty case4", 3, []interface{}{1, 2, 3}, 1, false},
		{"TestStack_IsEmpty case5", 3, []interface{}{1, 2, 3}, 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewStackWithCapacity(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				this.Push(tt.push[i])
			}
			for i := 0; i < tt.pop; i++ {
				this.Pop()
			}
			if got := this.IsEmpty(); got != tt.want {
				t.Errorf("Stack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsFull(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		push     []interface{}
		pop      int
		want     bool
	}{
		{"TestStack_IsFull case1", 0, []interface{}{}, 0, true},
		{"TestStack_IsFull case2", 3, []interface{}{}, 0, false},
		{"TestStack_IsFull case3", 3, []interface{}{1, 2, 3}, 0, true},
		{"TestStack_IsFull case4", 3, []interface{}{1, 2, 3}, 1, false},
		{"TestStack_IsFull case5", 3, []interface{}{1, 2, 3}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewStackWithCapacity(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				this.Push(tt.push[i])
			}
			for i := 0; i < tt.pop; i++ {
				this.Pop()
			}
			if got := this.IsFull(); got != tt.want {
				t.Errorf("Stack.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Clear(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		push     []interface{}
		pop      int
		want     int
	}{
		{"TestStack_Clear case1", 0, []interface{}{}, 0, 0},
		{"TestStack_Clear case2", 3, []interface{}{1, 2, 3}, 0, 0},
		{"TestStack_Clear case3", 3, []interface{}{1, 2, 3}, 1, 0},
		{"TestStack_Clear case4", 3, []interface{}{1, 2, 3}, 2, 0},
		{"TestStack_Clear case5", 3, []interface{}{1, 2, 3}, 3, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewStackWithCapacity(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				this.Push(tt.push[i])
			}
			for i := 0; i < tt.pop; i++ {
				this.Pop()
			}
			this.Clear()
			if got := this.Size(); got != tt.want {
				t.Errorf("Stack.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		push     []interface{}
		pop      int
		want     int
	}{
		{"TestStack_Size case1", 0, []interface{}{}, 0, 0},
		{"TestStack_Size case2", 3, []interface{}{1}, 0, 1},
		{"TestStack_Size case3", 3, []interface{}{1, 2, 3}, 0, 3},
		{"TestStack_Size case4", 3, []interface{}{1, 2, 3}, 1, 2},
		{"TestStack_Size case5", 3, []interface{}{1, 2, 3}, 2, 1},
		{"TestStack_Size case6", 3, []interface{}{1, 2, 3}, 3, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewStackWithCapacity(tt.capacity)
			for i := 0; i < len(tt.push); i++ {
				this.Push(tt.push[i])
			}
			for i := 0; i < tt.pop; i++ {
				this.Pop()
			}
			if got := this.Size(); got != tt.want {
				t.Errorf("Stack.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
