package applyleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(Values []int) (*ListNode, map[int]*ListNode) {
	l := len(Values)
	if l == 0 {
		panic("Wrong input")
	}
	m := make(map[int]*ListNode, len(Values))
	var recurrence func(start int) *ListNode
	recurrence = func(start int) *ListNode {
		if start == l {
			return nil
		}
		node := &ListNode{Val: Values[start]}
		m[start] = node
		node.Next = recurrence(start + 1)
		return node
	}
	return recurrence(0), m
}

func (ln *ListNode) GetValueArray() []int {
	res := []int{}
	cur := ln
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

type LRUCache struct {
	capacity int
	length   int
	data     map[int]*InteractiveList
	head     *InteractiveList // 哨兵节点
	tail     *InteractiveList // 哨兵节点
}

func Constructor(capacity int) LRUCache {
	if capacity == 0 {
		panic("Wrong input")
	}

	return LRUCache{
		capacity: capacity,
		data:     make(map[int]*InteractiveList, capacity),
		head:     &InteractiveList{},
		tail:     &InteractiveList{},
	}
}

func (lc *LRUCache) Get(key int) int {
	if lc.length == 0 {
		return -1
	}
	if v, ok := lc.data[key]; ok {
		lc.replaceNew(v)
		return v.Val
	}
	return -1
}

func (lc *LRUCache) Put(key int, value int) {
	if v, ok := lc.data[key]; ok {
		v.Val = value
		lc.replaceNew(v)
		return
	}
	if lc.length == lc.capacity {
		lc.refresh()
	}
	node := &InteractiveList{Key: key, Val: value}
	lc.data[key] = node
	if lc.length == 0 {
		node.Pre, node.Next = lc.head, lc.tail
		lc.head.Next = node
		lc.tail.Pre = node
		lc.length++
		return
	}
	preTmp := lc.tail.Pre
	node.Pre = preTmp
	node.Next = lc.tail
	preTmp.Next = node
	lc.tail.Pre = node
	lc.length++
}

// 存储空间满了,开始移除最久未使用的
func (lc *LRUCache) refresh() {
	delete(lc.data, lc.head.Next.Key)
	nextTmp := lc.head.Next.Next
	if nextTmp != nil {
		nextTmp.Pre = lc.head
	}
	lc.head.Next = nextTmp
	lc.length--
}

// 刷新数据,保持使用的在前面
func (lc *LRUCache) replaceNew(v *InteractiveList) {
	v.Pre.Next = v.Next
	v.Next.Pre = v.Pre
	lc.tail.Pre.Next = v
	v.Pre = lc.tail.Pre
	v.Next = lc.tail
	lc.tail.Pre = v
}

type InteractiveList struct {
	Key  int
	Val  int
	Next *InteractiveList
	Pre  *InteractiveList
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(arr []int) (*TreeNode, map[int]*TreeNode) {
	l, m := len(arr), make(map[int]*TreeNode, len(arr))
	var execBuild func(start int, arr []int) *TreeNode
	execBuild = func(start int, arr []int) *TreeNode {
		if start >= l || arr[start] == -1 {
			return nil
		}
		node := &TreeNode{}
		node.Val = arr[start]
		node.Left = execBuild(start*2+1, arr)
		node.Right = execBuild(start*2+2, arr)
		m[start] = node
		return node
	}
	return execBuild(0, arr), m
}
