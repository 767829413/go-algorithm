package string

import "fmt"

type acNode struct {
	v         rune             //字符值
	children  map[rune]*acNode //构建Trie树的hash表
	isEndChar bool             //true标识结尾字符
	length    int              //isEndChar=true记录模式串长度
	fail      *acNode          //失败指针
}

func newAcNode(s rune) *acNode {
	return &acNode{
		v:        s,
		children: make(map[rune]*acNode),
		length:   -1,
	}
}

type AcTrie struct {
	Root *acNode
}

func NewAcTrie() *AcTrie {
	return &AcTrie{
		Root: newAcNode('/'),
	}
}

//匹配
func (at *AcTrie) match(text []rune) {
	num := len(text)
	cur := at.Root
	for i := 0; i < num; i++ {
		for _, ok := cur.children[text[i]]; !ok && cur != at.Root; {
			//失败指针替换,加快匹配
			cur = cur.fail
		}
		if v, ok := cur.children[text[i]]; !ok {
			//如果没有匹配的，从root开始重新匹配
			cur = at.Root
		} else {
			cur = v
			tmp := cur
			for tmp != at.Root {
				if tmp.isEndChar {
					pos := i - tmp.length + 1
					//打印出可以匹配的模式串
					fmt.Println("匹配起始下标: ", pos, "长度: ", tmp.length)
				}
				tmp = tmp.fail
			}
		}
	}

}

//构建模式串Trie树
func (at *AcTrie) Insert(s []rune) {
	cur := at.Root
	num := len(s)
	for i := 0; i < num; i++ {
		if v, ok := cur.children[s[i]]; ok {
			cur = v
		} else {
			cur.children[s[i]] = newAcNode(s[i])
			cur = cur.children[s[i]]
		}
	}
	cur.isEndChar = true
	cur.length = num
}

//构建Trie树上节点的失败指针
func (at *AcTrie) BuildFailPoint() {
	q := newQueueArr()
	q.enQueue(at.Root)
	for !q.isEmpty() {
		p := q.deQueue()
		for _, pc := range p.children {
			if p == at.Root {
				pc.fail = at.Root
			} else {
				f := p.fail
				for f != nil {
					if fc, ok := f.children[pc.v]; ok {
						pc.fail = fc
						break
					}
					f = f.fail
				}
				if f == nil {
					pc.fail = at.Root
				}
			}
			q.enQueue(pc)
		}
	}
}

type arr struct {
	Items []*acNode
}

func newQueueArr() *arr {
	return &arr{
		Items: []*acNode{},
	}
}

//入队
func (q *arr) enQueue(item *acNode) bool {
	q.Items = append(q.Items, item)
	return true
}

//出队
func (q *arr) deQueue() (item *acNode) {
	if q.isEmpty() {
		return nil
	}
	item = q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return
}

//队列是否为空
func (q *arr) isEmpty() bool {
	return len(q.Items) == 0
}
