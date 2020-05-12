package tree

import "fmt"

/**
左倾红黑树
 */

type llrbNode struct {
	value interface{}
	index int
	times int
	left  *llrbNode
	right *llrbNode
	color bool
}

func newLlrbNode(index int, v interface{}) *llrbNode {
	return &llrbNode{
		index: index,
		value: v,
	}
}

func (ln *llrbNode) IsRed() bool {
	return ln.color
}

func (ln *llrbNode) setRed() {
	ln.color = true
}

func (ln *llrbNode) setBlack() {
	ln.color = false
}

//左旋
func (ln *llrbNode) rotateLeft() (right *llrbNode) {
	right = ln.right
	if right == nil {
		return
	}
	ln.right = right.left
	right.left = ln
	right.color = ln.color
	ln.setRed()
	return
}

//右旋
func (ln *llrbNode) rotateRight() (left *llrbNode) {
	left = ln.left
	if left == nil {
		return
	}
	ln.left = left.right
	left.right = ln
	left.color = ln.color
	ln.setRed()
	return
}

//颜色反转(防止一个节点左右均为红节点,或者相邻节点均为红节点)
func (ln *llrbNode) flipColor() {
	ln.color = !ln.color
	if ln.left != nil {
		ln.left.color = !ln.left.color
	}
	if ln.right != nil {
		ln.right.color = !ln.right.color
	}
}

type LLRBTree struct {
	Root *llrbNode
}

func NewLLRBTree() *LLRBTree {
	return &LLRBTree{}
}

//插入节点
func (t *LLRBTree) Inset(index int, v interface{}) {
	newNode := newLlrbNode(index, v)
	newNode.setRed()
	t.Root = t.add(t.Root, newNode)
	t.Root.setBlack()
}

func (t *LLRBTree) add(fNode, node *llrbNode) *llrbNode {
	if fNode == nil {
		return node
	}

	if fNode.index == node.index {
		fNode.times++
		fNode.value = node.value
	} else if node.index > fNode.index {
		fNode.right = t.add(fNode.right, node)
	} else {
		fNode.left = t.add(fNode.left, node)
	}

	curNode := fNode
	if (curNode.right != nil && curNode.left != nil) && (curNode.right.IsRed() && !curNode.left.IsRed()) {
		curNode = curNode.rotateLeft()
	} else {
		if (curNode.left != nil && curNode.left.left != nil) && curNode.left.IsRed() && curNode.left.left.IsRed() {
			curNode = curNode.rotateRight()
		}
		if (curNode.left != nil && curNode.right != nil) && curNode.left.IsRed() && curNode.right.IsRed() {
			curNode.flipColor()
		}
	}
	return curNode
}

//删除节点
func (t *LLRBTree) Delete(index int) {
	//当找不到值时直接返回
	if t.Find(index) == nil {
		return
	}
	if t.Root.left != nil && t.Root.right != nil {
		if !t.Root.left.IsRed() && !t.Root.right.IsRed() {
			t.Root.setRed()
		}
	}
	t.Root = t.del(t.Root, index)
	if t.Root != nil {
		t.Root.setBlack()
	}
}

func (t *LLRBTree) Find(index int) *llrbNode {
	if t.Root == nil {
		return nil
	}
	curNode := t.Root
	for curNode != nil {
		if curNode.index == index {
			return curNode
		} else if curNode.index > index {
			curNode = curNode.left
		} else {
			curNode = curNode.right
		}
	}
	return curNode
}

//左移红节点
func (t *LLRBTree) moveLeft(node *llrbNode) *llrbNode {
	node.flipColor()
	if node.right != nil && node.right.IsRed() {
		node.right = node.right.rotateRight()
		node = node.rotateLeft()
		node.flipColor()
	}
	return node
}

//右移红节点
func (t *LLRBTree) moveRight(node *llrbNode) *llrbNode {
	node.flipColor()
	if (node.left != nil && node.left.left != nil) && node.left.left.IsRed() {
		node = node.rotateRight()
		node.flipColor()
	}
	return node
}

//对该节点所在的子树删除元素
func (t *LLRBTree) del(fNode *llrbNode, index int) *llrbNode {
	curNode := fNode
	//删除的元素比子树根节点小，需要从左子树删除
	if index < curNode.index {
		//因为从左子树删除，所以要判断是否需要红色左移
		if curNode.left != nil {
			if !curNode.left.IsRed() && curNode.left.left != nil && !curNode.left.left.IsRed() {
				curNode = t.moveLeft(curNode)
			}
			curNode.left = t.del(curNode.left, index)
		}
	} else {
		//删除的元素等于或大于树根节点
		if curNode.left != nil && curNode.left.IsRed() {
			curNode = curNode.rotateRight()
		}
		//值相等，且没有右孩子节点，那么该节点一定是要被删除的叶子节点，直接删除
		if index == curNode.index && curNode.right == nil {
			return nil
		}
		//因为从右子树删除，所以要判断是否需要红色右移
		if curNode.right != nil {
			if !curNode.right.IsRed() && curNode.right.left != nil && !curNode.right.left.IsRed() {
				//右儿子和右儿子的左儿子都不是红色节点，那么没法递归下去，先红色右移
				curNode = t.moveRight(curNode)
			}
			//删除的节点找到了，它是中间节点，需要用最小后驱节点来替换它，然后删除最小后驱节点
			if index == curNode.index {
				minNode := t.FindMin(curNode.right)
				curNode.value = minNode.value
				curNode.index = minNode.index
				curNode.times = minNode.times
				//删除其最小后驱节点
				curNode.right = t.DeleteMin(curNode.right)
			} else {
				//删除的元素比子树根节点大，需要从右子树删除
				curNode.right = t.del(curNode.right, index)
			}
		}
	}
	//修复左倾红黑树特征
	return t.Recover(curNode)
}

func (t *LLRBTree) FindMin(node *llrbNode) *llrbNode {
	nowNode := node
	preNode := node
	//左子树为空，表面已经是最左的节点了，该值就是最小值,一直左子树查找
	for nowNode.left != nil {
		preNode = nowNode
		nowNode = nowNode.left
	}
	return preNode
}

func (t *LLRBTree) DeleteMin(node *llrbNode) *llrbNode {
	nowNode := node
	if nowNode.left == nil {
		return nil
	}
	//判断是否需要红色左移，因为最小元素在左子树中
	if !nowNode.left.IsRed() && nowNode.left.left != nil && !nowNode.left.left.IsRed() {
		nowNode = t.moveLeft(nowNode)
	}
	//递归从左子树删除
	nowNode.left = t.DeleteMin(nowNode.left)
	//修复左倾红黑树特征
	return t.Recover(nowNode)
}

func (t *LLRBTree) Recover(node *llrbNode) *llrbNode {
	nowNode := node
	//红链接在右边，左旋恢复，让红链接只出现在左边
	if nowNode.right != nil && nowNode.right.IsRed() {
		nowNode = nowNode.rotateLeft()
	}
	//连续两个左链接为红色，那么进行右旋
	if nowNode.left != nil {
		if nowNode.left.IsRed() && nowNode.left.left != nil && nowNode.left.left.IsRed() {
			nowNode = nowNode.rotateRight()
		}
		if nowNode.right != nil && nowNode.right.IsRed() && nowNode.left.IsRed() {
			nowNode.flipColor()
		}
	}
	return nowNode
}

func (t *LLRBTree) Print() {
	t.midTraverse(t.Root)
}

func (t *LLRBTree) midTraverse(node *llrbNode) {
	if node.left != nil {
		t.midTraverse(node.left)
	}
	if node != nil {
		fmt.Println(node.index, node.value)
	}
	if node.right != nil {
		t.midTraverse(node.right)
	}

}
