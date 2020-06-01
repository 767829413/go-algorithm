package string

type trieNode struct {
	s         rune
	children  map[rune]*trieNode
	isEndChar bool
}

type Trie struct {
	Root *trieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: &trieNode{
			s:        rune('/'),
			children: make(map[rune]*trieNode),
		},
	}
}

func (t *Trie) Insert(text []rune) {
	cur := t.Root
	num := len(text)
	for i := 0; i < num; i++ {
		if v, ok := cur.children[text[i]]; ok {
			cur = v
		} else {
			cur.children[text[i]] = &trieNode{
				s:        text[i],
				children: make(map[rune]*trieNode),
			}
			cur = cur.children[text[i]]
		}
	}
	cur.isEndChar = true
}

func (t *Trie) Find(text []rune) bool {
	cur := t.Root
	num := len(text)
	for i := 0; i < num; i++ {
		if v, ok := cur.children[text[i]]; ok {
			cur = v
		}
	}
	if !cur.isEndChar {
		return false
	}
	return true
}
