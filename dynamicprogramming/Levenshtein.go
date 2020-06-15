package dynamicprogramming

/**
Levenshtein distance (莱文斯坦距离)
 */
type Levenshtein struct {
	partA   []rune
	aNum    int
	partB   []rune
	bNum    int
	minDest int
}

func NewLevenshtein(partA, partB []rune, aNum, bNum int) *Levenshtein {
	return &Levenshtein{
		partA:   partA,
		partB:   partB,
		aNum:    aNum,
		bNum:    bNum,
		minDest: int(^uint(0) >> 1),
	}
}

//回溯实现
func (l *Levenshtein) GetMinDest() int {
	l.modifyDest(0, 0, 0)
	return l.minDest
}

func (l *Levenshtein) modifyDest(i, j, md int) {
	if i == l.aNum || j == l.bNum {
		if i < l.aNum {
			md += l.aNum - i
		}
		if j < l.bNum {
			md += l.bNum - j
		}
		if l.minDest > md {
			l.minDest = md
		}
		return
	}
	if l.partA[i] == l.partB[j] {
		l.modifyDest(i+1, j+1, md)
	} else {
		//删除A[i] 或 B[j]前面添加
		l.modifyDest(i+1, j, md+1)
		//删除B[i] 或 A[i]前面添加
		l.modifyDest(i, j+1, md+1)
		//A[i] B[j] 同时删除 或 同时替换
		l.modifyDest(i+1, j+1, md+1)
	}
}

//采用动态规划
func (l *Levenshtein) GetMinDest1() int {
	return l.modifyDest1()
}

func (l *Levenshtein) modifyDest1() int {
	//初始化状态数组
	states := l.initStatus()
	for i := 1; i < l.aNum; i++ {
		for j := 1; j < l.bNum; j++ {
			//从上一个状态跳转过来 [i-1][j] [i][j-1]都是做了添加或删除操作
			if l.partA[i] == l.partB[j] {
				states[i][j] = l.min(states[i-1][j]+1, states[i][j-1]+1, states[i-1][j-1])
			} else {
				states[i][j] = l.min(states[i-1][j]+1, states[i][j-1]+1, states[i-1][j-1]+1)
			}
		}
	}
	return states[l.aNum-1][l.bNum-1]
}

func (l *Levenshtein) initStatus() [][]int {
	states := make([][]int, l.aNum)
	for i := 0; i < l.aNum; i++ {
		arr := make([]int, l.bNum)
		states[i] = arr
	}
	for j := 0; j < l.bNum; j++ {
		if l.partA[0] == l.partB[j] {
			states[0][j] = j
		} else if j != 0 {
			states[0][j] = states[0][j-1] + 1
		} else {
			states[0][j] = 1
		}
	}
	for i := 0; i < l.aNum; i++ {
		if l.partB[0] == l.partA[i] {
			states[i][0] = i
		} else if i != 0 {
			states[i][0] = states[i-1][0] + 1
		} else {
			states[i][0] = 1
		}
	}
	return states
}

func (l *Levenshtein) min(a, b, c int) int {
	switch true {
	case a < b && a < c:
		return a
	case b < c:
		return b
	default:
		return c
	}
}
