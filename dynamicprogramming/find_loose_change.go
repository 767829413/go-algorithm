package dynamicprogramming

type FindLooseChange struct {
	changeType []int
	n          int
	need       int
	min        int
}

func NewFindLooseChange(changeType []int, n, need int) *FindLooseChange {
	return &FindLooseChange{
		changeType: changeType,
		n:          n,
		need:       need,
		min:        int(^uint(0) >> 1),
	}
}

//回溯解决
func (f *FindLooseChange) GetMin() int {
	f.getMinItem1(f.need, 0)
	return f.min
}

func (f *FindLooseChange) getMinItem1(cp, num int) {
	if cp == 0 {
		if num < f.min {
			f.min = num
		}
	}
	for i := 0; i < f.n; i++ {
		lastSum := cp - f.changeType[i]
		if lastSum >= 0 {
			f.getMinItem1(lastSum, num+1)
		}
	}
}

//动态规划之状态转移表
func (f *FindLooseChange) GetMin2() int {
	return f.getMinItem2()
}

func (f *FindLooseChange) getMinItem2() int {
	status := make([]int, f.need+1)
	status[0] = 0
	for i := 0; i < f.n; i++ {
		status[f.changeType[i]] = 1
	}

	for i := 1; i <= f.need; i++ {
		if status[i] == 0 {
			min := int(^uint(0) >> 1)
			for j := 0; j < f.n; j++ {
				if i-f.changeType[j] > 0 {
					tempNum := status[i-f.changeType[j]] + 1
					if tempNum < min {
						min = tempNum
					}
				}
			}
			status[i] = min
		}
	}
	return status[f.need]
}
