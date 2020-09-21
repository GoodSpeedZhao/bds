package bitSet

type BitSet struct {
	dataStore []uint64
	size      int
}

const (
	cap            = ^uint64(0) // 所有位设置为 1
	log2ColumnSize = uint64(6)
	columnSize     = uint64(64)
)

func NewBitSet() *BitSet {
	return &BitSet{
		dataStore: []uint64{},
		size:      0,
	}
}

func NewBitSetWithLength(length int) (bitSet *BitSet) {
	defer func() {
		if r := recover(); r != nil {
			bitSet = &BitSet{
				dataStore: make([]uint64, 0),
				size:      0,
			}
		}
	}()

	bitSet = &BitSet{
		dataStore: make([]uint64, length),
		size:      length,
	}

	return
}

func (bs *BitSet) Has(newVal uint64) bool {
	row := bs.getRowIdx(newVal)
	return (row < len(bs.dataStore)) && ((bs.dataStore[row] & uint64(0x1) << (newVal % columnSize)) != 0)
}

func (bs *BitSet) Set(newVal uint64) {
	row := bs.getRowIdx(newVal)
	if row >= len(bs.dataStore) {
		newLen := row + 1
		newDataStore := make([]uint64, newLen, newLen)
		copy(newDataStore, bs.dataStore)
		bs.dataStore = newDataStore
	}
	bs.dataStore[row] |= uint64(0x1) << (newVal % columnSize)
}

func (bs *BitSet) Delete(val uint64) {
	row := bs.getRowIdx(val)
	if row < len(bs.dataStore) {
		bs.dataStore[row] &^= uint64(0x1) << (val % columnSize)
	}
}

func (bs *BitSet) Size() int {
	return bs.size
}

func (bs *BitSet) IsEmpty() bool {
	return bs.Size() == 0
}

func (bs *BitSet) Clear() {
	bs.dataStore = []uint64{}
	bs.size = 0
}

func (bs *BitSet) getRowIdx(target uint64) int {
	return int(target >> log2ColumnSize)
}

func (bs *BitSet) getLength() int {
	return len(bs.dataStore)
}

func (bs *BitSet) compareByLength(other *BitSet) (smaller, bigger *BitSet) {
	if bs.getLength() < other.getLength() {
		return bs, other
	} else {
		return other, bs
	}
}

func (bs *BitSet) Intersection(other *BitSet) *BitSet {
	smaller, bigger := bs.compareByLength(other)
	result := NewBitSetWithLength(smaller.getLength())
	for i := 0; i < smaller.getLength(); i++ {
		result.dataStore[i] = smaller.dataStore[i] & bigger.dataStore[i]
	}
	return result
}

func (bs *BitSet) Union(other *BitSet) *BitSet {
	smaller, bigger := bs.compareByLength(other)
	result := NewBitSetWithLength(smaller.getLength())
	for i := 0; i < smaller.getLength(); i++ {
		result.dataStore[i] = smaller.dataStore[i] | bigger.dataStore[i]
	}
	return result
}

func (bs *BitSet) Diff(other *BitSet) *BitSet {
	smaller, bigger := bs.compareByLength(other)
	result := NewBitSetWithLength(smaller.getLength())
	for i := 0; i < smaller.getLength(); i++ {
		result.dataStore[i] = smaller.dataStore[i] &^ bigger.dataStore[i]
	}
	return result
}

func popcount(x uint64) (n uint64) {
	x -= (x >> 1) & 0x5555555555555555
	x = (x>>2)&0x3333333333333333 + x&0x3333333333333333
	x += x >> 4
	x &= 0x0f0f0f0f0f0f0f0f
	x *= 0x0101010101010101
	return x >> 56
}

// // 整数的奇偶性判断
// func judge1(a int) {
// 	return a & 1
// }

// // 判断n是否是2的正整数冪
// func judge2(a int) {
// 	return a & (a - 1)
// }

// // 判断第n位是否为1
// func judge3(a int, n int) {
// 	return a & (1 << n)
// }

// // 将第n位设置为 1
// func judge4(a int, n int) {
// 	return a | (1 << n)
// }

// // 将第n位设置为 0
// func judge5(a int, n int) {
// 	return a & ~(1 << n)
// }

// // 将第n位设置取反
// func judge6(a int, n int) {
// 	return a ^ (1 << n)
// }

// // 将最右边的 1 设置为 0
// func judge7(a int) {
// 	return a & (a - 1)
// }
