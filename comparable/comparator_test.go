package comparable

import (
	"testing"
	"time"
)

func TestIntComparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestIntComparator case1", args{1, 1, 0}},
		{"TestIntComparator case2", args{1, 2, -1}},
		{"TestIntComparator case3", args{2, 1, 1}},
		{"TestIntComparator case4", args{-1, -1, 0}},
		{"TestIntComparator case5", args{-1, 2, -1}},
		{"TestIntComparator case6", args{2, -1, 1}},
		{"TestIntComparator case7", args{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntComparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("IntComparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestInt8Comparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestInt8Comparator case1", args{int8(1), int8(1), 0}},
		{"TestInt8Comparator case2", args{int8(1), int8(2), -1}},
		{"TestInt8Comparator case3", args{int8(2), int8(1), 1}},
		{"TestInt8Comparator case4", args{int8(-1), int8(-1), 0}},
		{"TestInt8Comparator case5", args{int8(-1), int8(2), -1}},
		{"TestInt8Comparator case6", args{int8(2), int8(-1), 1}},
		{"TestInt8Comparator case7", args{int8(0), int8(0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8Comparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("Int8Comparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestInt16Comparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestInt16Comparator case1", args{int16(1), int16(1), 0}},
		{"TestInt16Comparator case2", args{int16(1), int16(2), -1}},
		{"TestInt16Comparator case3", args{int16(2), int16(1), 1}},
		{"TestInt16Comparator case4", args{int16(-1), int16(-1), 0}},
		{"TestInt16Comparator case5", args{int16(-1), int16(2), -1}},
		{"TestInt16Comparator case6", args{int16(2), int16(-1), 1}},
		{"TestInt16Comparator case7", args{int16(0), int16(0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16Comparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("Int16Comparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestInt32Comparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestInt32Comparator case1", args{int32(1), int32(1), 0}},
		{"TestInt32Comparator case2", args{int32(1), int32(2), -1}},
		{"TestInt32Comparator case3", args{int32(2), int32(1), 1}},
		{"TestInt32Comparator case4", args{int32(-1), int32(-1), 0}},
		{"TestInt32Comparator case5", args{int32(-1), int32(2), -1}},
		{"TestInt32Comparator case6", args{int32(2), int32(-1), 1}},
		{"TestInt32Comparator case7", args{int32(0), int32(0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Comparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("Int32Comparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestInt64Comparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestInt64Comparator case1", args{int64(1), int64(1), 0}},
		{"TestInt64Comparator case2", args{int64(1), int64(2), -1}},
		{"TestInt64Comparator case3", args{int64(2), int64(1), 1}},
		{"TestInt64Comparator case4", args{int64(-1), int64(-1), 0}},
		{"TestInt64Comparator case5", args{int64(-1), int64(2), -1}},
		{"TestInt64Comparator case6", args{int64(2), int64(-1), 1}},
		{"TestInt64Comparator case7", args{int64(0), int64(0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Comparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("Int64Comparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestUintComparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestUintComparator case1", args{uint(1), uint(1), 0}},
		{"TestUintComparator case2", args{uint(1), uint(2), -1}},
		{"TestUintComparator case3", args{uint(2), uint(1), 1}},
		{"TestUintComparator case4", args{uint(1), uint(0), 1}},
		{"TestUintComparator case5", args{uint(0), uint(1), -1}},
		{"TestUintComparator case6", args{uint(0), uint(0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintComparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("UintComparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestUint8Comparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestUint8Comparator case1", args{uint8(1), uint8(1), 0}},
		{"TestUint8Comparator case2", args{uint8(1), uint8(2), -1}},
		{"TestUint8Comparator case3", args{uint8(2), uint8(1), 1}},
		{"TestUint8Comparator case4", args{uint8(1), uint8(0), 1}},
		{"TestUint8Comparator case5", args{uint8(0), uint8(1), -1}},
		{"TestUint8Comparator case6", args{uint8(0), uint8(0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8Comparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("Uint8Comparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestUint16Comparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestUint16Comparator case1", args{uint16(1), uint16(1), 0}},
		{"TestUint16Comparator case2", args{uint16(1), uint16(2), -1}},
		{"TestUint16Comparator case3", args{uint16(2), uint16(1), 1}},
		{"TestUint16Comparator case4", args{uint16(1), uint16(0), 1}},
		{"TestUint16Comparator case5", args{uint16(0), uint16(1), -1}},
		{"TestUint16Comparator case6", args{uint16(0), uint16(0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16Comparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("Uint16Comparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestUint32Comparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestUint32Comparator case1", args{uint32(1), uint32(1), 0}},
		{"TestUint32Comparator case2", args{uint32(1), uint32(2), -1}},
		{"TestUint32Comparator case3", args{uint32(2), uint32(1), 1}},
		{"TestUint32Comparator case4", args{uint32(1), uint32(0), 1}},
		{"TestUint32Comparator case5", args{uint32(0), uint32(1), -1}},
		{"TestUint32Comparator case6", args{uint32(0), uint32(0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32Comparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("Uint32Comparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestUint64Comparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestUint64Comparator case1", args{uint64(1), uint64(1), 0}},
		{"TestUint64Comparator case2", args{uint64(1), uint64(2), -1}},
		{"TestUint64Comparator case3", args{uint64(2), uint64(1), 1}},
		{"TestUint64Comparator case4", args{uint64(1), uint64(0), 1}},
		{"TestUint64Comparator case5", args{uint64(0), uint64(1), -1}},
		{"TestUint64Comparator case6", args{uint64(0), uint64(0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64Comparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("Uint64Comparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestFloat32Comparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestFloat32Comparator case1", args{float32(1.0), float32(1.0), 0}},
		{"TestFloat32Comparator case2", args{float32(1.0), float32(2.0), -1}},
		{"TestFloat32Comparator case3", args{float32(2.0), float32(1.0), 1}},
		{"TestFloat32Comparator case4", args{float32(-1.0), float32(-1.0), 0}},
		{"TestFloat32Comparator case5", args{float32(-1.0), float32(2.0), -1}},
		{"TestFloat32Comparator case6", args{float32(2.0), float32(-1.0), 1}},
		{"TestFloat32Comparator case7", args{float32(0.0), float32(0.0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32Comparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("Float32Comparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestFloat64Comparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestFloat64Comparator case1", args{1.0, 1.0, 0}},
		{"TestFloat64Comparator case2", args{1.0, 2.0, -1}},
		{"TestFloat64Comparator case3", args{2.0, 1.0, 1}},
		{"TestFloat64Comparator case4", args{-1.0, -1.0, 0}},
		{"TestFloat64Comparator case5", args{-1.0, 2.0, -1}},
		{"TestFloat64Comparator case6", args{2.0, -1.0, 1}},
		{"TestFloat64Comparator case7", args{0.0, 0.0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Comparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("Float64Comparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestByteComparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestByteComparator case1", args{byte(1), byte(1), 0}},
		{"TestByteComparator case2", args{byte(1), byte(2), -1}},
		{"TestByteComparator case3", args{byte(2), byte(1), 1}},
		{"TestByteComparator case4", args{byte(1), byte(0), 1}},
		{"TestByteComparator case5", args{byte(0), byte(1), -1}},
		{"TestByteComparator case6", args{byte(0), byte(0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteComparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("ByteComparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestRuneComparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestRuneComparator case1", args{rune(1), rune(1), 0}},
		{"TestRuneComparator case2", args{rune(1), rune(2), -1}},
		{"TestRuneComparator case3", args{rune(2), rune(1), 1}},
		{"TestRuneComparator case4", args{rune(-1), rune(-1), 0}},
		{"TestRuneComparator case5", args{rune(-1), rune(2), -1}},
		{"TestRuneComparator case6", args{rune(2), rune(-1), 1}},
		{"TestRuneComparator case7", args{rune(0), rune(0), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RuneComparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("RuneComparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestTimeComparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	now := time.Now()
	tests := []struct {
		name string
		args args
	}{

		{"TestTimeComparator case1", args{now, now, 0}},
		{"TestTimeComparator case2", args{now.Add(1 * time.Hour), now, 1}},
		{"TestTimeComparator case3", args{now, now.Add(1 * time.Hour), -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeComparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("TimeComparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestStringComparator(t *testing.T) {
	type args struct {
		a    interface{}
		b    interface{}
		want int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestStringComparator case1", args{"a", "a", 0}},
		{"TestStringComparator case2", args{"a", "b", -1}},
		{"TestStringComparator case3", args{"b", "a", 1}},
		{"TestStringComparator case4", args{"aa", "aab", -1}},
		{"TestStringComparator case5", args{"", "", 0}},
		{"TestStringComparator case6", args{"a", "", 1}},
		{"TestStringComparator case7", args{"", "a", -1}},
		{"TestStringComparator case8", args{"a", "A", 1}},
		{"TestStringComparator case9", args{"bba", "bbA", 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringComparator(tt.args.a, tt.args.b); got != tt.args.want {
				t.Errorf("StringComparator() = %v, want %v", got, tt.args.want)
			}
		})
	}
}
