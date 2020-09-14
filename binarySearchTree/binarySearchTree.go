package binarySearchTree

import (
	"errors"

	"github.com/GoodSpeedZhao/bds/comparable"
)

type Node struct {
	value interface{}
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
	size int
	cmp  comparable.Comparator
}

func NewBinarySearchTree(comparator comparable.Comparator) *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
		size: 0,
		cmp:  comparator,
	}
}

func NewNode(newVal interface{}) *Node {
	return &Node{
		value: newVal,
		left:  nil,
		right: nil,
	}
}

func (bst *BinarySearchTree) Insert(newVal interface{}) (err error) {
	if bst.root == nil {
		bst.root = NewNode(newVal)
		bst.size++
		err = nil
		return
	}

	_, err = bst.insert(bst.root, newVal)
	if err == nil {
		bst.size++
	}
	return
}

func (bst *BinarySearchTree) insert(root *Node, newVal interface{}) (*Node, error) {
	if root == nil {
		return NewNode(newVal), nil
	} else if root.value == newVal {
		return nil, errors.New("Insert Failed. Already has val.")
	} else if bst.cmp(root.value, newVal) == 1 {
		if node, err := bst.insert(root.left, newVal); err != nil {
			return nil, err
		} else {
			root.left = node
		}
	} else {
		if node, err := bst.insert(root.right, newVal); err != nil {
			return nil, err
		} else {
			root.right = node
		}
	}

	return root, nil
}

func (bst *BinarySearchTree) Delete(val interface{}) (err error) {
	if bst.IsEmpty() {
		err = errors.New("BinarySearchTree is Empty")
		return
	}

	root, err := bst.delete(bst.root, val)
	if err == nil {
		bst.root = root
		bst.size--
	}
	return
}

func (bst *BinarySearchTree) delete(root *Node, val interface{}) (*Node, error) {
	if root == nil {
		return nil, errors.New("Delete Failed. Not has val.")
	} else if bst.cmp(root.value, val) == 1 {
		if node, err := bst.delete(root.left, val); err != nil {
			return nil, err
		} else {
			root.left = node
		}
	} else if bst.cmp(root.value, val) == -1 {
		if node, err := bst.delete(root.right, val); err != nil {
			return nil, err
		} else {
			root.right = node
		}
	} else {
		if root.left == nil {
			right := root.right
			root.right = nil
			return right, nil
		} else if root.right == nil {
			left := root.left
			root.left = nil
			return left, nil
		} else {
			newRoot := bst.findMin(root.right)
			newRoot.right = bst.deleteMin(root.right)
			newRoot.left = root.left
			root.left = nil
			root.right = nil
			return newRoot, nil
		}
	}

	return root, nil
}

func (bst *BinarySearchTree) DeleteMin() error {
	if bst.IsEmpty() {
		return errors.New("BinarySearchTree is Empty")
	}

	bst.root = bst.deleteMin(bst.root)
	return nil
}

func (bst *BinarySearchTree) deleteMin(root *Node) *Node {
	if root.left == nil {
		right := root.right
		root.right = nil
		return right
	}

	root.left = bst.deleteMin(root.left)
	return root
}

func (bst *BinarySearchTree) DeleteMax() error {
	if bst.IsEmpty() {
		return errors.New("BinarySearchTree is Empty")
	}

	bst.root = bst.deleteMax(bst.root)
	bst.size--
	return nil
}

func (bst *BinarySearchTree) deleteMax(root *Node) *Node {
	if root.right == nil {
		left := root.left
		root.left = nil
		return left
	}

	root.right = bst.deleteMax(root.right)
	return root
}

func (bst *BinarySearchTree) Contains(val interface{}) bool {
	if bst.IsEmpty() {
		return false
	}

	return bst.contains(bst.root, val)
}

func (bst *BinarySearchTree) contains(root *Node, val interface{}) bool {
	if root == nil {
		return false
	} else if root.value == val {
		return true
	} else if bst.cmp(root.value, val) == 1 {
		return bst.contains(root.left, val)
	} else {
		return bst.contains(root.right, val)
	}
}

func (bst *BinarySearchTree) FindMin() (interface{}, error) {
	if bst.IsEmpty() {
		return nil, errors.New("BinarySearchTree is Empty")
	}

	node := bst.findMin(bst.root)
	return node.value, nil
}

func (bst *BinarySearchTree) findMin(root *Node) *Node {
	for root.left != nil {
		root = root.left
	}
	return root
}

func (bst *BinarySearchTree) FindMax() (interface{}, error) {
	if bst.IsEmpty() {
		return nil, errors.New("BinarySearchTree is Empty")
	}
	node := bst.findMax(bst.root)
	return node.value, nil
}

func (bst *BinarySearchTree) findMax(root *Node) *Node {
	for root.right != nil {
		root = root.right
	}
	return root
}

func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.size == 0
}

func (bst *BinarySearchTree) Size() int {
	return bst.size
}

func (bst *BinarySearchTree) Clear() {
	if bst.IsEmpty() {
		return
	}

	bst.clear(bst.root)
	bst.root = nil
	bst.size = 0
}

func (bst *BinarySearchTree) clear(root *Node) {
	if root == nil {
		return
	}

	bst.clear(root.left)
	bst.clear(root.right)
	root.left = nil
	root.right = nil
}
