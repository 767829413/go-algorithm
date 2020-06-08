package backtrack

type Pattern struct {
	matched bool
	pattern []rune //正则表达式
	plen    int    //正则表达式长度
}

func NewPattern(pattern []rune, plen int) *Pattern {
	return &Pattern{
		pattern: pattern,
		plen:    plen,
	}
}

func (p *Pattern) Match(text []rune, tlen int) bool {
	p.rematch(0, 0, text, tlen)
	return p.matched
}

func (p *Pattern) rematch(pIndex, tIndex int, text []rune, tlen int) {
	if p.matched { //标识已经匹配
		return
	}
	if pIndex == p.plen { //正则串已经到结尾
		if tIndex == tlen { //文本串已经匹配到结尾
			p.matched = true
		}
		return
	}
	if p.pattern[pIndex] == '*' { //匹配任意个字符(0~n)
		for i := 0; i < tlen-tIndex; i++ {
			p.rematch(pIndex+1, tIndex+i, text, tlen)
		}
	} else if p.pattern[pIndex] == '?' { //匹配0个或1个字符
		p.rematch(pIndex+1, tIndex, text, tlen)
		p.rematch(pIndex+1, tIndex+1, text, tlen)
	} else if tIndex < tlen && p.pattern[pIndex] == text[tIndex] { //纯字符匹配
		p.rematch(pIndex+1, tIndex+1, text, tlen)
	}
}
