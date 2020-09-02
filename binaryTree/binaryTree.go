package binaryTree

import "github.com/GoodSpeedZhao/bds/deque"

type BinaryTree struct {
	root *Node
	size int
}

type Node struct {
	value interface{}
	left  *Node
	right *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		root: nil,
		size: 0,
	}
}

func NewNode(newVal interface{}) *Node {
	return &Node{
		value: newVal,
		left:  nil,
		right: nil,
	}
}

func CreateWithLevelOrder(newVal []interface{}) *BinaryTree {
	tree := NewBinaryTree()

	var root *Node
	if len(newVal) > 0 {
		curIdx := 1
		root = NewNode(newVal[0])
		que := deque.NewDequeWithCapacity(len(newVal))
		que.Append(root)
		for !que.IsEmpty() {
			levelCount := que.Size()
			tree.size += levelCount
			for i := 0; i < levelCount; i++ {
				tmp, err := que.Shift()
				if err != nil {
					panic(err)
				}

				if curIdx < len(newVal) {
					if newVal[curIdx] != nil {
						tmp.(*Node).left = NewNode(newVal[curIdx])
						que.Append(tmp.(*Node).left)
					}
				}
				curIdx++

				if curIdx < len(newVal) {
					if newVal[curIdx] != nil {
						tmp.(*Node).right = NewNode(newVal[curIdx])
						que.Append(tmp.(*Node).right)
					}
				}
				curIdx++
			}
		}
	}

	tree.root = root

	return tree
}

func (bt *BinaryTree) PreOrder() interface{} {
	if bt.IsEmpty() {
		return nil
	}

	values := make([]interface{}, 0, bt.size)
	bt.preOrder(bt.root, &values)
	return values
}

func (bt *BinaryTree) preOrder(root *Node, values *[]interface{}) {
	if root == nil {
		return
	}

	*values = append(*values, root.value)
	bt.preOrder(root.left, values)
	bt.preOrder(root.right, values)
}

func (bt *BinaryTree) InOrder() interface{} {
	if bt.IsEmpty() {
		return nil
	}

	values := make([]interface{}, 0, bt.size)
	bt.inOrder(bt.root, &values)
	return values
}

func (bt *BinaryTree) inOrder(root *Node, values *[]interface{}) {
	if root == nil {
		return
	}

	bt.inOrder(root.left, values)
	*values = append(*values, root.value)
	bt.inOrder(root.right, values)
}

func (bt *BinaryTree) PostOrder() interface{} {
	if bt.IsEmpty() {
		return nil
	}

	values := make([]interface{}, 0, bt.size)
	bt.postOrder(bt.root, &values)
	return values
}

func (bt *BinaryTree) postOrder(root *Node, values *[]interface{}) {
	if root == nil {
		return
	}

	bt.postOrder(root.left, values)
	bt.postOrder(root.right, values)
	*values = append(*values, root.value)
}

func (bt *BinaryTree) LevelOrder() interface{} {
	if bt.IsEmpty() {
		return nil
	}

	values := make([]interface{}, 0, bt.size)

	queue := deque.NewDeque()
	queue.Append(bt.root)
	for !queue.IsEmpty() {
		if node, err := queue.Shift(); err != nil {
			return nil
		} else {
			popNode := node.(*Node)
			values = append(values, popNode.value)
			if popNode.left != nil {
				queue.Append(popNode.left)
			}
			if popNode.right != nil {
				queue.Append(popNode.right)
			}
		}
	}

	return values
}

func (bt *BinaryTree) IsEmpty() bool {
	return bt.size == 0
}

func (bt *BinaryTree) Size() int {
	return bt.size
}

func (bt *BinaryTree) Root() *Node {
	return bt.root
}

func (bt *BinaryTree) Clear() {
	if bt.IsEmpty() {
		return
	}

	bt.clear(bt.root)
	bt.root = nil
	bt.size = 0
}

func (bt *BinaryTree) clear(root *Node) {
	if root == nil {
		return
	}

	bt.clear(root.left)
	bt.clear(root.right)
	root.left = nil
	root.right = nil
}
