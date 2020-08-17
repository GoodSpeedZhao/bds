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

func (this *BinaryTree) CreateWithLevelOrder(newVal []interface{}) *BinaryTree {
	tree := NewBinaryTree()

	var root *Node
	if len(newVal) > 0 {
		curIdx := 1
		root = NewNode(newVal[0])
		que := deque.NewDequeWithCapacity(len(newVal))
		que.Append(root)
		for !que.IsEmpty() {
			levelCount := que.Size()
			for i := 0; i < levelCount; i++ {
				tmp, err := que.Shift()
				if err != nil {
					panic(err)
				}

				if curIdx < len(newVal) {
					if newVal[curIdx] != nil {
						tmp.(*Node).left = NewNode(newVal[curIdx])
						que.Append(tmp.(*Node).left)
						curIdx++
					}
				}

				if curIdx < len(newVal) {
					if newVal[curIdx] != nil {
						tmp.(*Node).right = NewNode(newVal[curIdx])
						que.Append(tmp.(*Node).right)
						curIdx++
					}
				}
			}
		}
	}

	tree.root = root
	tree.size = len(newVal)

	return tree
}

func (this *BinaryTree) PreOrder() []interface{} {
	if this.IsEmpty() {
		return nil
	}

	values := make([]interface{}, this.size)
	this.preOrder(this.root, values)
	return values
}

func (this *BinaryTree) preOrder(root *Node, values []interface{}) {
	if root == nil {
		return
	}

	values = append(values, root.value)
	this.preOrder(root.left, values)
	this.preOrder(root.right, values)
}

func (this *BinaryTree) InOrder() []interface{} {
	if this.IsEmpty() {
		return nil
	}

	values := make([]interface{}, this.size)
	this.inOrder(this.root, values)
	return values
}

func (this *BinaryTree) inOrder(root *Node, values []interface{}) {
	if root == nil {
		return
	}

	this.inOrder(root.left, values)
	values = append(values, root.value)
	this.inOrder(root.right, values)
}

func (this *BinaryTree) PostOrder() []interface{} {
	if this.IsEmpty() {
		return nil
	}

	values := make([]interface{}, this.size)
	this.postOrder(this.root, values)
	return values
}

func (this *BinaryTree) postOrder(root *Node, values []interface{}) {
	if root == nil {
		return
	}

	this.postOrder(root.left, values)
	this.postOrder(root.right, values)
	values = append(values, root.value)
}

func (this *BinaryTree) LevelOrder() []interface{} {
	if this.IsEmpty() {
		return nil
	}

	values := make([]interface{}, this.size)

	dq := deque.NewDeque()
	dq.Append(this.root)
	for !dq.IsEmpty() {
		if node, err := dq.Shift(); err != nil {
			return nil
		} else {
			popNode := node.(*Node)
			values = append(values, popNode.value)
			if popNode.left != nil {
				dq.Append(popNode.left)
			}
			if popNode.right != nil {
				dq.Append(popNode.right)
			}
		}
	}

	return values
}

func (this *BinaryTree) IsEmpty() bool {
	return this.size == 0
}

func (this *BinaryTree) Size() int {
	return this.size
}

func (this *BinaryTree) Root() *Node {
	return this.root
}

func (this *BinaryTree) Clear() {
	if this.IsEmpty() {
		return
	}

	this.clear(this.root)
	this.root = nil
	this.size = 0
}

func (this *BinaryTree) clear(root *Node) {
	if root == nil {
		return
	}

	this.clear(root.left)
	this.clear(root.right)
	root.left = nil
	root.right = nil
}
