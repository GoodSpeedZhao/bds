package trieTree

type Node struct {
	children map[rune]*Node
	count    map[rune]int
	isWord   bool
}

type TrieTree struct {
	root *Node
	size int
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		root: NewTrieNode(),
		size: 0,
	}
}

func NewTrieNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
		count:    make(map[rune]int),
		isWord:   false,
	}
}

func (tt *TrieTree) Insert(words ...string) {
	for _, word := range words {
		r := []rune(word)
		result := tt.insert(tt.root, r)
		if !result {
			tt.size++
		}
	}
}

func (tt *TrieTree) insert(root *Node, r []rune) (result bool) {
	if len(r) == 1 {
		if child, ok := root.children[r[0]]; !ok {
			child = NewTrieNode()
			root.children[r[0]] = child
			root.count[r[0]]++
		}

		result = root.children[r[0]].isWord
		if !root.children[r[0]].isWord {
			root.children[r[0]].isWord = true
		}
		return
	}

	if child, ok := root.children[r[0]]; !ok {
		child = NewTrieNode()
		root.children[r[0]] = child
	}

	result = tt.insert(root.children[r[0]], r[1:])
	if !result {
		root.count[r[0]]++
	}
	return result
}

func (tt *TrieTree) Delete(word string) (result bool) {
	if tt.IsEmpty() {
		return false
	}

	r := []rune(word)
	result = tt.delete(tt.root, r)
	if result {
		tt.size--
	}
	return
}

func (tt *TrieTree) delete(root *Node, r []rune) (result bool) {
	if child, ok := root.children[r[0]]; !ok {
		return false
	} else if len(r) == 1 {
		root.count[r[0]]--
		if root.count[r[0]] == 0 {
			delete(root.count, r[0])
			delete(root.children, r[0])
		}

		result = child.isWord
		if child.isWord {
			child.isWord = false
		}
		return
	}

	result = tt.delete(root.children[r[0]], r[1:])
	if result {
		root.count[r[0]]--
		if root.count[r[0]] == 0 {
			delete(root.count, r[0])
			delete(root.children, r[0])
		}
	}
	return
}

func (tt *TrieTree) Contains(word string) bool {
	if tt.IsEmpty() {
		return false
	}

	return tt.match(tt.root, []rune(word), 0)
}

// . 表示替代任何一个字符
func (tt *TrieTree) match(root *Node, word []rune, index int) bool {
	if index == len(word) {
		return root.isWord
	}

	ch := word[index]

	if ch == '.' {
		for _, child := range root.children {
			if tt.match(child, word, index+1) {
				return true
			}
		}
		return false
	} else {
		if child, ok := root.children[ch]; !ok {
			return false
		} else {
			return tt.match(child, word, index+1)
		}
	}
}

func (tt *TrieTree) PrefixNumber(preStr string) int {
	if tt.IsEmpty() {
		return 0
	}

	curNode := tt.root
	for _, c := range []rune(preStr) {
		if child, ok := curNode.children[c]; ok {
			curNode = child
		} else {
			return 0
		}
	}

	sum := 0
	if curNode.isWord {
		sum++
	}
	for _, val := range curNode.count {
		sum += val
	}
	return sum
}

func (tt *TrieTree) Clear() {
	tt.clear(tt.root)
	tt.root = NewTrieNode()
	tt.size = 0
}

func (tt *TrieTree) clear(root *Node) {
	if root == nil {
		return
	}

	for _, child := range root.children {
		tt.clear(child)
	}
	root.children = nil
}

func (tt *TrieTree) IsEmpty() bool {
	return tt.Size() == 0
}

func (tt *TrieTree) Size() int {
	return tt.size
}
