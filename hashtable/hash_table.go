package hashtable

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
)

const (
	defaultInitalCapacity = 8
	loadFactor            = 0.75
)

type entry struct {
	k    string
	v    interface{}
	next *entry
}

func NewEntry(k string, v interface{}, next *entry) *entry {
	return &entry{
		k:    k,
		v:    v,
		next: next,
	}
}

type HashTable struct {
	table []*entry
	size  int
	use   float64
}

func NewHashTable() *HashTable {
	return &HashTable{
		table: make([]*entry, defaultInitalCapacity),
	}
}

func (ht *HashTable) Put(k string, v interface{}) {
	index := ht.hash(k)
	//位置未被引用时,创建哨兵节点
	if ht.table[index] == nil {
		ht.table[index] = NewEntry("", nil, nil)
	}
	tmp := ht.table[index]
	//新增
	if tmp.next == nil {
		tmp.next = NewEntry(k, v, nil)
		ht.size++
		ht.use++
		//判断是否需要动态扩容
		if ht.use >= float64(len(ht.table))*loadFactor {
			ht.resize()
		}
	} else {
		//出现散列冲突,利用链表法解决(线性探测,二次探测,多重hash,链表法各有千秋)
		for tmp.next != nil {
			tmp = tmp.next
			//k相同则覆盖,结束插入
			if tmp.k == k {
				tmp.v = v
				return
			}
		}
		temp := ht.table[index].next
		ht.table[index].next = NewEntry(k, v, temp)
		ht.size++
	}
}

func (ht *HashTable) Delete(k string) interface{} {
	index := ht.hash(k)
	head := ht.table[index]
	if head == nil || head.next == nil {
		return nil
	}
	pre := head
	cur := pre
	for cur.next != nil {
		cur = cur.next
		if cur.k == k {
			v := cur.v
			pre.next = cur.next
			ht.size--
			return v
		}
		pre = pre.next
	}
	return nil
}

func (ht *HashTable) Find(k string) interface{} {
	index := ht.hash(k)
	tmp := ht.table[index]
	if tmp == nil || tmp.next == nil {
		return nil
	}
	for tmp.next != nil {
		tmp = tmp.next
		if tmp.k == k {
			return tmp.v
		}
	}
	return nil
}

func (ht *HashTable) Print(k string) (format string) {
	index := ht.hash(k)
	cur := ht.table[index]
	format = "head->"
	for nil != cur {
		format += "k=" + cur.k + "|" + "v=" + fmt.Sprintf("%+v", cur.v)
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	return format
}

func (ht *HashTable) hash(k string) int {
	if k == "" {
		return 0
	}
	num := xxhash.ChecksumString32(k)
	index := (num ^ (num >> 16)) & (uint32(len(ht.table) - 1))
	return int(index)
}

//扩容
func (ht *HashTable) resize() {
	count := len(ht.table) * 2
	newTable := make([]*entry, count)
	copy(newTable, ht.table)
	ht.table = newTable
	ht.use = 0
	for i := 0; i < count; i++ {
		//没有占用位置或没有插入数据只有哨兵
		if newTable[i] == nil || newTable[i].next == nil {
			continue
		}
		tmp := newTable[i]
		for tmp.next != nil {
			tmp = tmp.next
			newIndex := ht.hash(tmp.k)
			//如果计算后hash一致,则无需搬移
			if newIndex == i {
				continue
			}
			if newTable[newIndex] == nil {
				ht.use++
				//创建哨兵
				newTable[newIndex] = NewEntry("", nil, nil)
			}
			newNode := NewEntry(tmp.k, tmp.v, newTable[newIndex].next)
			newTable[newIndex].next = newNode
		}
	}
}
