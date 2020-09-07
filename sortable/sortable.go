package sortable

import (
	"sort"

	"github.com/GoodSpeedZhao/bds/comparable"
)

type Sortable struct {
	values []interface{}
	cmp    comparable.Comparator
}

func (s Sortable) Len() int {
	return len(s.values)
}
func (s Sortable) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
func (s Sortable) Less(i, j int) bool {
	return s.cmp(s.values[i], s.values[j]) < 0
}

func Sort(values []interface{}, cmp comparable.Comparator) {
	sort.Sort(Sortable{values, cmp})
}
