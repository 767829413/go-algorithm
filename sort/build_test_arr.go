package sort

import (
	"math/rand"
	"time"
)

func BuildTestArr(num, max, min int) []int {
	rand.Seed(time.Now().UnixNano())
	//构建随机[]int 最大值 63
	arr := make([]int, num, num)
	for i := 0; i < num; i++ {
		randNum := rand.Intn(max-min) + min
		arr[i] = randNum
	}
	return arr
}
