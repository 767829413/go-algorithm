package tree

import "fmt"

type data struct {
	v     interface{}
	index int
}

func NewData(v interface{}, index int) *data {
	return &data{
		v:     v,
		index: index,
	}
}

type node struct {
	V     *data
	Left  *node
	Right *node
}

func NewNode(v interface{}, index int) *node {
	return &node{
		V: NewData(v, index),
	}
}

type BinarySearchTree struct {
	tree *node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Insert(v interface{}, index int) {
	if bst.tree == nil {
		bst.tree = NewNode(v, index)
		return
	}
	cur := bst.tree
	for cur != nil {
		if cur.V.index < index {
			if cur.Right == nil {
				cur.Right = NewNode(v, index)
				return
			}
			cur = cur.Right
		} else if cur.V.index > index {
			if cur.Left == nil {
				cur.Left = NewNode(v, index)
				return
			}
			cur = cur.Left
		} else {
			cur.V.v = v
			return
		}
	}
	return
}

/**
删除节点要注意该节点左右均存在的情况,
对于一个二叉查找树来说,替换该删除的节点应该是它右树的最小节点
 */
func (bst *BinarySearchTree) Delete(index int) {
	var fNode *node = nil
	var child *node = nil
	node := bst.tree
	//找到删除的节点
	for node != nil && node.V.index != index {
		fNode = node
		if index > node.V.index {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	if node == nil {
		return
	}
	//如果要删除的节点下左右均在,则和该节点右树的最小叶子节点交换位置
	if node.Left != nil && node.Right != nil {
		//查找最小右树节点
		minNode := node.Right
		minFNode := node
		for minNode.Left != nil {
			minFNode = minNode
			minNode = minNode.Left
		}
		//将该节点替换道删除节点位置
		node.V = minNode.V
		node = minNode
		fNode = minFNode
	}
	//删除叶子节点或只有左节点或只有右节点
	if node.Left != nil {
		child = node.Left
	} else if node.Right != nil {
		child = node.Right
	}
	if fNode == nil {
		bst.tree = child
	} else if fNode.Left == node {
		fNode.Left = child
	} else if fNode.Right == node {
		fNode.Right = child
	}
}

func (bst *BinarySearchTree) Find(index int) interface{} {
	if bst.tree == nil {
		return nil
	}
	node := bst.tree
	for node != nil {
		if index < node.V.index {
			node = node.Left
		} else if index > node.V.index {
			node = node.Right
		} else {
			return node.V.v
		}
	}
	return nil
}

func (bst *BinarySearchTree) FindMax() interface{} {
	if bst.tree == nil {
		return nil
	}
	node := bst.tree
	for node.Right != nil {
		node = node.Right
	}
	return node.V.v
}

func (bst *BinarySearchTree) FindMin() interface{} {
	if bst.tree == nil {
		return nil
	}
	node := bst.tree
	for node.Left != nil {
		node = node.Left
	}
	return node.V.v
}

func (bst *BinarySearchTree) Print() {
	bst.midOrderTraverse(bst.tree)
}

func (bst *BinarySearchTree) midOrderTraverse(node *node) {
	if node == nil {
		return
	}
	bst.midOrderTraverse(node.Left)
	fmt.Println(node.V)
	bst.midOrderTraverse(node.Right)
}
