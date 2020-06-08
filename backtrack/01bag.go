package backtrack

var max int
/**
cw表示当前已经装进去的物品的重量和；index表示考察到哪个物品了；
bw背包重量；items表示每个物品的重量；n表示物品个数
假设背包可承受重量100，物品个数10，物品重量存储在数组a中，那可以这样调用函数：
f(0, 0, []int, 10, 100)
 */
func BagQuestion(index, cw int, items []int, n int, bw int) {
	if cw == bw || index == n { //背包满了或者物品放完
		if max < cw {
			max = cw
		}
		return
	}
	BagQuestion(index+1, cw, items, n, bw) //当前物品不装
	if cw+items[index] <= bw {
		BagQuestion(index+1, cw+items[index], items, n, bw) //当前物品装
	}
}

func GetMax() int {
	return max
}
