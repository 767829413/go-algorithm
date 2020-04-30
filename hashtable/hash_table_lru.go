package hashtable

import "fmt"

/**
基于散列表的LRU算法
 */

//链表选择双向链表
type doubleNode struct {
	k    interface{}
	v    interface{}
	next *doubleNode
	pre  *doubleNode
}

func NewDoubleNode(k interface{}, v interface{}) *doubleNode {
	return &doubleNode{
		k: k,
		v: v,
	}
}

type LRUHashTable struct {
	table    map[interface{}]*doubleNode //散列表来缓存记录
	head     *doubleNode
	tail     *doubleNode
	length   int
	capacity int
}

func NewLRUHashTable(capacity int) *LRUHashTable {
	head := NewDoubleNode("", nil)
	tail := NewDoubleNode("", nil)
	head.next = tail
	tail.pre = head
	return &LRUHashTable{
		table:    make(map[interface{}]*doubleNode),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

/**
新增
 */
func (lru *LRUHashTable) Add(k interface{}, v interface{}) {
	if existNode, ok := lru.table[k]; ok {
		existNode.v = v
		lru.moveToHead(existNode)
	} else {
		newNode := NewDoubleNode(k, v)
		lru.table[k] = newNode
		lru.addNode(newNode)
		lru.length++
		if lru.length > lru.capacity {
			tailNode := lru.popTail()
			delete(lru.table, tailNode.k)
			lru.length--
		}
	}
}

/**
删除
 */
func (lru *LRUHashTable) Delete(k interface{}) {
	if exitNode, ok := lru.table[k]; ok {
		lru.delNode(exitNode)
		delete(lru.table, k)
		lru.length--
	}
	return
}

/**
获取
*/
func (lru *LRUHashTable) Get(k interface{}) interface{} {
	if exitNode, ok := lru.table[k]; ok {
		lru.moveToHead(exitNode)
		return exitNode.v
	}
	return nil
}

/**
打印
 */
func (lru *LRUHashTable) Print() (format string) {
	format = "head -> "
	cur := lru.head.next
	for cur != nil && cur.next != nil {
		format += "(k=" + fmt.Sprintf("%+v", cur.k) + " | v=" + fmt.Sprintf("%+v", cur.v) + ")"
		format += " -> "
		cur = cur.next
	}
	format += "tail"
	return format
}

//新节点朝前放
func (lru *LRUHashTable) addNode(newNode *doubleNode) {
	newNode.pre = lru.head
	newNode.next = lru.head.next
	lru.head.next.pre = newNode
	lru.head.next = newNode
}

//容量满了需要移除尾部节点
func (lru *LRUHashTable) popTail() *doubleNode {
	tailNode := lru.tail.pre
	lru.delNode(tailNode)
	return tailNode
}

//移除节点
func (lru *LRUHashTable) delNode(node *doubleNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

//将节点朝前
func (lru *LRUHashTable) moveToHead(node *doubleNode) {
	lru.delNode(node)
	lru.addNode(node)
}
