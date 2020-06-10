package dynamicprogramming

import "fmt"

type Count struct {
	values []int    //商品价格
	num    int      //数量
	term   int      //条件
	states [][]bool //状态数据
}

func NewCount(values []int, num, term int) *Count {
	c := &Count{
		values: values,
		num:    num,
		term:   term,
	}
	states := make([][]bool, num)
	n := 3*term + 1
	for i := 0; i < num; i++ {
		states[i] = make([]bool, n)
	}
	states[0][0] = true
	if values[0] <= 3*term {
		states[0][values[0]] = true
	}
	c.states = states
	return c
}

func (c *Count) count() {
	for i := 1; i < c.num; i++ {
		//不买第i个商品
		for j := 0; j < 3*c.term; j++ {
			if c.states[i-1][j] {
				c.states[i][j] = c.states[i-1][j]
			}
		}
		//买第i个商品
		for j := 0; j <= 3*c.term-c.values[i]; j++ {
			if c.states[i-1][j] {
				c.states[i][c.values[i]+j] = true
			}
		}
	}
	j := 0
	for j = c.term; j < c.term*3+1; j++ {
		if c.states[c.num-1][j] {
			break
		}
	}
	if j == 3*c.term+1 {
		return
	}
	for i := c.num - 1; i >= 1; i-- {
		if j-c.values[i] >= 0 && c.states[i-1][j-c.values[i]] {
			fmt.Println(c.values[i], "")
			j = j - c.values[i]
		}
	}
	if j != 0 {
		fmt.Println(c.values[0])
	}
}
