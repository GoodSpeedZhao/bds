package sortable

import (
	"strings"
	"testing"
	"time"

	"github.com/GoodSpeedZhao/bds/comparable"
)

func TestSortable_Len(t *testing.T) {
	tests := []struct {
		name   string
		values []interface{}
		cmp    comparable.Comparator
		want   int
	}{
		{"TestSortable_Len case1", []interface{}{}, comparable.IntComparator, 0},
		{"TestSortable_Len case2", []interface{}{1}, comparable.IntComparator, 1},
		{"TestSortable_Len case3", []interface{}{1, 6, 9}, comparable.IntComparator, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sortable{tt.values, tt.cmp}
			if got := s.Len(); got != tt.want {
				t.Errorf("Sortable.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortable_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		values []interface{}
		cmp    comparable.Comparator
		args   args
	}{
		{"TestSortable_Swap case1", []interface{}{1, 6, 9}, comparable.IntComparator, args{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sortable{tt.values, tt.cmp}
			s.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestSortable_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		values []interface{}
		cmp    comparable.Comparator
		args   args
		want   bool
	}{
		{"TestSortable_Less case1", []interface{}{1, 6, 9}, comparable.IntComparator, args{0, 1}, true},
		{"TestSortable_Less case2", []interface{}{1, 6, 9}, comparable.IntComparator, args{1, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sortable{tt.values, tt.cmp}
			if got := s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Sortable.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort_Int(t *testing.T) {
	type args struct {
		values []interface{}
		cmp    comparable.Comparator
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestSort_Int case1", args{[]interface{}{1, 6, 9}, comparable.IntComparator}},
		{"TestSort_Int case2", args{[]interface{}{9, 6, 1}, comparable.IntComparator}},
		{"TestSort_Int case3", args{[]interface{}{9, 6, 12, 1}, comparable.IntComparator}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.values, tt.args.cmp)
			for i := 1; i < len(tt.args.values); i++ {
				if tt.args.values[i-1].(int) > tt.args.values[i].(int) {
					t.Errorf("sorted error.")
				}
			}
		})
	}
}

func TestSort_Float64(t *testing.T) {
	type args struct {
		values []interface{}
		cmp    comparable.Comparator
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestSort_Float64 case1", args{[]interface{}{1.0, 6.0, 9.0}, comparable.Float64Comparator}},
		{"TestSort_Float64 case2", args{[]interface{}{9.0, 6.0, 1.0}, comparable.Float64Comparator}},
		{"TestSort_Float64 case3", args{[]interface{}{9.0, 6.0, 12.0, 1.0}, comparable.Float64Comparator}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.values, tt.args.cmp)
			for i := 1; i < len(tt.args.values); i++ {
				if tt.args.values[i-1].(float64) > tt.args.values[i].(float64) {
					t.Errorf("sorted error.")
				}
			}
		})
	}
}

func TestSort_String(t *testing.T) {
	type args struct {
		values []interface{}
		cmp    comparable.Comparator
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestSort_String case1", args{[]interface{}{"c", "b", "a"}, comparable.StringComparator}},
		{"TestSort_String case2", args{[]interface{}{"abc", "bcd", "efg"}, comparable.StringComparator}},
		{"TestSort_String case3", args{[]interface{}{"", "a", "A"}, comparable.StringComparator}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.values, tt.args.cmp)
			for i := 1; i < len(tt.args.values); i++ {
				if strings.Compare(tt.args.values[i-1].(string), tt.args.values[i].(string)) > 0 {
					t.Errorf("sorted error.")
				}
			}
		})
	}
}

func TestSort_Time(t *testing.T) {
	type args struct {
		values []interface{}
		cmp    comparable.Comparator
	}
	now := time.Now()
	tests := []struct {
		name string
		args args
	}{
		{"TestSort_String case1", args{[]interface{}{now.Add(3 * time.Hour), now, now.Add(1 * time.Hour)}, comparable.TimeComparator}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.values, tt.args.cmp)
			for i := 1; i < len(tt.args.values); i++ {
				if tt.args.values[i-1].(time.Time).After(tt.args.values[i].(time.Time)) {
					t.Errorf("sorted error.")
				}
			}
		})
	}
}

func TestSort_Custom(t *testing.T) {
	type Rank struct {
		id       int
		nickName string
	}
	rankCmp := func(a, b interface{}) int {
		aRank := a.(Rank)
		bRank := b.(Rank)
		switch {
		case aRank.id > bRank.id:
			return 1
		case aRank.id < bRank.id:
			return -1
		default:
			return 0
		}
	}

	type args struct {
		values []interface{}
		cmp    comparable.Comparator
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestSort_Custom case1",
			args{[]interface{}{
				Rank{100, "a"},
				Rank{1, "b"},
				Rank{50, "c"},
				Rank{20, "d"},
				Rank{75, "e"},
			}, rankCmp}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.values, tt.args.cmp)
			for i := 1; i < len(tt.args.values); i++ {
				if tt.args.values[i-1].(Rank).id > tt.args.values[i].(Rank).id {
					t.Errorf("sorted error.")
				}
			}
		})
	}
}
