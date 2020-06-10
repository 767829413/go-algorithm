package dynamicprogramming

import "fmt"

/**
示例1
通过记录状态,防止重复操作
限制条件 背包重量 期望目标 装载物品最重
 */
type bageg1 struct {
	max     int      //最大承重
	weight  []int    //物品重量
	itemNum int      //物品数量
	lw      int      //背包限制重量
	status  [][]bool //状态集合
}

func Newbageg1(lw, itemNum int, weight []int) *bageg1 {
	bag := &bageg1{
		weight:  weight,
		itemNum: itemNum,
		lw:      lw,
	}
	status := make([][]bool, itemNum)
	for i := 0; i < itemNum; i++ {
		status[i] = make([]bool, lw)
	}
	bag.status = status
	return bag
}

func (b *bageg1) GetMax(index, cw int) int {
	b.count(index, cw)
	return b.max
}

func (b *bageg1) count(index, cw int) {
	if index == b.itemNum || cw == b.lw {
		if cw > b.max {
			b.max = cw
		}
		return
	}
	if b.status[index][cw] {
		return
	}
	//记录状态
	b.status[index][cw] = true
	//递归分支1,不装
	b.count(index+1, cw)
	//递归分支2 装入(剪枝,大于限制不装)
	w := cw + b.weight[index]
	if w <= b.lw {
		b.count(index+1, w)
	}
}

/**
示例2
通过动态规划的思想,构建流程,消耗空间 itemNum * lw 二维数组
限制条件 背包重量 期望目标 装载物品最重
 */
type bageg2 struct {
	weight  []int    //物品重量
	itemNum int      //物品数量
	lw      int      //背包限制重量
	states  [][]bool //状态集合
}

func Newbageg2(itemNum, lw int, weight []int) *bageg2 {
	bag := &bageg2{
		weight:  weight,
		itemNum: itemNum,
		lw:      lw,
	}
	states := make([][]bool, itemNum)
	for i := 0; i < itemNum; i++ {
		states[i] = make([]bool, lw+1)
	}
	states[0][0] = true
	if weight[0] <= lw {
		states[0][weight[0]] = true
	}
	bag.states = states
	return bag
}

func (b *bageg2) GetMax() int {
	//动态规划状态转移
	for i := 1; i < b.itemNum; i++ {
		//分支1 第i个物品不放入背包
		for j := 0; j < b.lw; j++ {
			if b.states[i-1][j] {
				b.states[i][j] = b.states[i-1][j]
			}
		}
		//分支2 第i个物品放入背包
		for j := 0; j < b.lw-b.weight[i]; j++ {
			if b.states[i-1][j] {
				b.states[i][j+b.weight[i]] = true
			}
		}
	}
	for i := b.lw; i >= 0; i-- {
		if b.states[b.itemNum-1][i] {
			return i
		}
	}
	return -1
}

/**
示例3
通过动态规划的思想,构建流程,消耗空间 itemNum 一维数组
限制条件 背包重量 期望目标 装载物品最重
 */
type bageg3 struct {
	weight  []int  //物品重量
	itemNum int    //物品数量
	lw      int    //背包限制重量
	states  []bool //状态集合
}

func Newbageg3(itemNum, lw int, weight []int) *bageg3 {
	bag := &bageg3{
		weight:  weight,
		itemNum: itemNum,
		lw:      lw,
	}
	num := lw + 1
	states := make([]bool, num)
	states[0] = true
	if weight[0] <= lw {
		states[weight[0]] = true
	}
	bag.states = states
	return bag
}

func (b *bageg3) GetMax() int {
	for i := 1; i < b.itemNum; i++ {
		//第i个物品放入包中
		for j := b.lw - b.weight[i]; j >= 0; j-- {
			if b.states[j] {
				b.states[b.weight[i]+j] = true
			}
		}
	}
	fmt.Println(b.states)
	for i := b.lw; i >= 0; i-- {
		if b.states[i] {
			return i
		}
	}
	return -1
}

/**
示例4
通过动态规划的思想,构建流程,但是消耗空间 itemNum 一维数组
限制条件 背包重量 期望目标 装载物品总价值最高
 */
type bageg4 struct {
	weight  []int   //物品重量
	value   []int   //物品价值
	itemNum int     //物品数量
	lw      int     //背包限制重量
	states  [][]int //状态集合
}

func Newbageg4(itemNum, lw int, weight, value []int) *bageg4 {
	bag := &bageg4{
		weight:  weight,
		value:   value,
		itemNum: itemNum,
		lw:      lw,
	}
	states := make([][]int, itemNum)
	for i := 0; i < itemNum; i++ {
		c := make([]int, lw+1)
		for j := 0; j < lw+1; j++ {
			c[j] = -1
		}
		states[i] = c
	}
	states[0][0] = 0
	if weight[0] <= lw {
		states[0][weight[0]] = value[0]
	}
	bag.states = states
	return bag
}

func (b *bageg4) GetMax() int {
	for i := 1; i < b.itemNum; i++ {
		//不选择第i个物品
		for j := 0; j <= b.lw; j++ {
			if b.states[i-1][j] >= 0 {
				b.states[i][j] = b.states[i-1][j]
			}
		}
		//选择第i个物品
		for j := 0; j <= b.lw-b.weight[i]; j++ {
			if b.states[i-1][j] >= 0 {
				cv := b.value[i] + b.states[i-1][j]
				if cv > b.states[i][j+b.weight[i]] {
					b.states[i][j+b.weight[i]] = cv
				}
			}
		}
	}
	maxValue := -1
	for j := 0; j < b.lw; j++ {
		if b.states[b.itemNum-1][j] > maxValue {
			maxValue = b.states[b.itemNum-1][j]
		}
	}
	return maxValue
}
