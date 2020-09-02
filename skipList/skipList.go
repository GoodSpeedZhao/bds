package skipList

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

const (
	MAXLEVEL = 16
	MAXROUND = 65536
)

type Node struct {
	key   int
	value interface{}
	next  [MAXLEVEL]*Node
}

type SkipList struct {
	head  *Node
	tail  *Node
	size  int
	level int
}

func randLevel() int {
	return MAXLEVEL - int(math.Log2(1.0+MAXROUND*rand.Float64()))
}

func NewNode(newKey int, newVal interface{}) *Node {
	return &Node{
		key:   newKey,
		value: newVal,
		next:  [MAXLEVEL]*Node{},
	}
}

func NewSkipList() *SkipList {
	s := new(SkipList)

	s.head = NewNode(math.MinInt64, nil)
	s.tail = NewNode(math.MaxInt64, nil)

	for i := 0; i < MAXLEVEL; i++ {
		s.head.next[i] = s.tail
	}

	s.size = 0
	s.level = 1

	return s
}

func (s *SkipList) Search(key int) (*Node, error) {
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i].key < key {
			cur = cur.next[i]
		}
		if cur.next[i].key == key {
			return cur.next[i], nil
		}
	}

	return nil, errors.New("Not Found.")
}

func (s *SkipList) Add(key int, newVal interface{}) {
	pre := make([]*Node, MAXLEVEL)
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i].key < key {
			cur = cur.next[i]
		}
		pre[i] = cur
	}

	level := randLevel()

	newNode := NewNode(key, newVal)
	for i := 0; i < level; i++ {
		if pre[i] == nil {
			pre[i] = s.head
		}

		pre[i].next[i], newNode.next[i] = newNode, pre[i].next[i]
	}

	if s.level < level {
		s.level = level
	}

	s.size++
}

func (s *SkipList) Erase(key int) (result bool) {
	result = false
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i].key < key {
			cur = cur.next[i]
		}
		if cur.next[i].key == key {
			tmp := cur.next[i]
			cur.next[i] = cur.next[i].next[i]
			tmp.next[i] = nil
			result = true
		}
	}

	if result {
		s.size--
	}

	return
}

func (s *SkipList) Get(key int) (*Node, error) {
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i].key < key {
			cur = cur.next[i]
		}
		if cur.next[i].key == key {
			return cur.next[i], nil
		}
	}

	return nil, errors.New("Not Found.")
}

func (s *SkipList) Print() {
	cur := s.head
	for cur != nil {
		fmt.Println(cur.value)
		cur = cur.next[0]
	}
}

func (s *SkipList) GetSize() int {
	return s.size
}
