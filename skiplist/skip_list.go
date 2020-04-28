package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	maxLayer  = 16
	randomNum = 5
)

//构建跳表节点
type skipListNode struct {
	data     int
	layer    int
	forwards [maxLayer]*skipListNode
}

func NewSkipListNode(v int, layer int) *skipListNode {
	return &skipListNode{
		data:  v,
		layer: layer,
	}
}

type SkipList struct {
	head       *skipListNode
	layerCount int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:       NewSkipListNode(math.MinInt32, 0),
		layerCount: 1,
	}
}

func (sl *SkipList) Insert(v int) bool {
	layerRandom := sl.randomLayer()
	//构建新节点
	newNode := NewSkipListNode(v, layerRandom)
	//构建路径
	update := make([]*skipListNode, layerRandom)
	for i := 0; i < layerRandom; i++ {
		update[i] = sl.head
	}
	cur := sl.head
	//自顶之下索引查找
	for i := layerRandom - 1; i >= 0; i-- {
		for cur.forwards[i] != nil && cur.forwards[i].data < v {
			cur = cur.forwards[i]
		}
		update[i] = cur
	}
	for i := 0; i < layerRandom; i++ {
		newNode.forwards[i] = update[i].forwards[i]
		update[i].forwards[i] = newNode
	}
	if sl.layerCount < layerRandom {
		sl.layerCount = layerRandom
	}
	return true
}

func (sl *SkipList) Delete(v int) bool {
	//构建路径
	update := make([]*skipListNode, sl.layerCount)
	cur := sl.head
	//自顶向下索引查找
	start := sl.layerCount - 1
	for ; start >= 0; start-- {
		for cur.forwards[start] != nil && cur.forwards[start].data < v {
			cur = cur.forwards[start]
		}
		update[start] = cur
	}
	flag := false
	if update[0].forwards[0] != nil && update[0].forwards[0].data == v {
		i := sl.layerCount - 1
		for ; i >= 0; i-- {
			if update[i].forwards[i] != nil && update[i].forwards[i].data == v {
				update[i].forwards[i] = update[i].forwards[i].forwards[i]
			}
		}
		flag = true
	}
	for sl.layerCount > 1 && sl.head.forwards[sl.layerCount] == nil {
		sl.layerCount--
	}
	return flag
}

func (sl *SkipList) Find(v int) *skipListNode {
	cur := sl.head
	//自顶向下索引查找
	start := sl.layerCount - 1
	for ; start >= 0; start-- {
		for cur.forwards[start] != nil && cur.forwards[start].data < v {
			cur = cur.forwards[start]
		}
	}
	if cur.forwards[0] != nil && cur.forwards[0].data == v {
		return cur.forwards[0]
	}
	return nil
}

func (sl *SkipList) Print() string {
	cur := sl.head
	format := ""
	for cur.forwards[0] != nil {
		format += fmt.Sprintf("%+v", cur.forwards[0].data)
		format += "->"
		cur = cur.forwards[0]
	}
	format += "nil"
	return format
}

/**
1的概率50%
2的概率25%
3的概率12.5%
...
 */
func (sl *SkipList) randomLayer() int {
	layer := 1
	for (rand.Int31n(10) < randomNum) && (layer < maxLayer) {
		layer++
	}
	return layer
}
