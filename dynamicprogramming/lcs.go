package dynamicprogramming

/**
Longest common substring(最长公共子串)
*/
type LCS struct {
	partA   []rune
	aNum    int
	partB   []rune
	bNum    int
	maxDest int
}

func NewLCS(partA, partB []rune, aNum, bNum int) *LCS {
	return &LCS{
		partA:   partA,
		partB:   partB,
		aNum:    aNum,
		bNum:    bNum,
		maxDest: ^int(^uint(0) >> 1),
	}
}

//回溯处理
func (l *LCS) GetMaxDest() int {
	l.modifyDest(0, 0, 0)
	return l.maxDest
}

func (l *LCS) modifyDest(i, j, md int) {
	if i == l.aNum || j == l.bNum {
		if l.maxDest < md {
			l.maxDest = md
		}
		return
	}
	if l.partA[i] == l.partB[j] {
		l.modifyDest(i+1, j+1, md+1)
	} else {
		//删除A[i] 或 B[j]前面添加A[i]
		l.modifyDest(i+1, j, md)
		//删除B[i] 或 A[i]前面添加B[i]
		l.modifyDest(i, j+1, md)
		//A[i] B[j] 同时删除 或 同时替换
		l.modifyDest(i+1, j+1, md)
	}
}

//动态规划处理
func (l *LCS) GetMaxDest1() int {
	return l.modifyDest1()
}

func (l *LCS) modifyDest1() int {
	//初始化状态数组
	states := l.initStatus()
	for i := 1; i < l.aNum; i++ {
		for j := 1; j < l.bNum; j++ {
			//从上一个状态跳转过来 [i-1][j] [i][j-1]都是做了添加或删除操作
			if l.partA[i] == l.partB[j] {
				states[i][j] = l.max(states[i-1][j], states[i][j-1], states[i-1][j-1]+1)
			} else {
				states[i][j] = l.max(states[i-1][j], states[i][j-1], states[i-1][j-1])
			}
		}
	}
	return states[l.aNum-1][l.bNum-1]
}

func (l *LCS) initStatus() [][]int {
	status := make([][]int, l.aNum)
	for i := 0; i < l.aNum; i++ {
		arr := make([]int, l.bNum)
		status[i] = arr
	}
	for j := 0; j < l.bNum; j++ {
		if l.partA[0] == l.partB[j] {
			status[0][j] = 1
		} else if j != 0 {
			status[0][j] = status[0][j-1]
		} else {
			status[0][j] = 0
		}
	}
	for i := 0; i < l.aNum; i++ {
		if l.partB[0] == l.partA[i] {
			status[i][0] = 1
		} else if i != 0 {
			status[i][0] = status[i-1][0]
		} else {
			status[i][0] = 0
		}
	}
	return status
}

func (l *LCS) max(i int, i2 int, i3 int) int {
	max := ^int(^uint(0) >> 1)
	if i > max {
		max = i
	}
	if i2 > max {
		max = i2
	}
	if i3 > max {
		max = i3
	}
	return max
}
