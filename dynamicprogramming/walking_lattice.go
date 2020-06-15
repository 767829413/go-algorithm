package dynamicprogramming

var maxInt = int(^uint(0) >> 1)
//获取两数最小
func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

/**
状态转移表法
走正方形格子
 */
type DPStatusTable struct {
	w      [][]int
	n      int
	states [][]int
}

func NewDPStatusTable(w [][]int, n int) *DPStatusTable {
	d := &DPStatusTable{
		w: w,
		n: n,
	}
	d.states = d.initStates()
	return d
}

func (d *DPStatusTable) GetMinDest() int {
	return d.minDistDP()
}
func (d *DPStatusTable) minDistDP() int {
	for i := 1; i < d.n; i++ {
		for j := 1; j < d.n; j++ {
			d.states[i][j] = d.w[i][j] + minInt(d.states[i][j-1], d.states[i-1][j])
		}
	}
	return d.states[d.n-1][d.n-1]
}

func (d *DPStatusTable) initStates() [][]int {
	states := make([][]int, d.n)
	for i := 0; i < d.n; i++ {
		states[i] = make([]int, d.n)
	}
	//第一行初始化
	sumRow := 0
	for i := 0; i < d.n; i++ {
		sumRow += d.w[0][i]
		states[0][i] = sumRow
	}
	//第一列初始化
	sumColumn := 0
	for i := 0; i < d.n; i++ {
		sumColumn += d.w[i][0]
		states[i][0] = sumColumn
	}
	return states
}

/**
状态转移方程法
走正方形格子
 */

type DPStatusFunc struct {
	w      [][]int
	n      int
	remark [][]int
}

func NewDPStatusFunc(w [][]int, n int) *DPStatusFunc {
	d := &DPStatusFunc{
		w: w,
		n: n,
	}
	remark := make([][]int, n)
	for i := 0; i < n; i++ {
		remark[i] = make([]int, n)
	}
	d.remark = remark
	return d
}

func (d *DPStatusFunc) GetMinDest() int {
	return d.minDest(d.n-1, d.n-1)
}

func (d *DPStatusFunc) minDest(i, j int) int {
	if i == j && j == 0 {
		return d.w[0][0]
	}
	if d.remark[i][j] > 0 {
		return d.remark[i][j]
	}
	leftMin := maxInt
	if j-1 >= 0 {
		leftMin = d.minDest(i, j-1)
	}
	upMin := maxInt
	if i-1 >= 0 {
		upMin = d.minDest(i-1, j)
	}
	currMinDist := d.w[i][j] + minInt(leftMin, upMin)
	d.remark[i][j] = currMinDist
	return currMinDist
}
