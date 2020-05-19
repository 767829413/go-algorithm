package test

import (
	"fmt"
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

func PrintSortComparison(arr []int, f func(arr []int)) {
	//排序前
	fmt.Println(arr)
	fmt.Println(len(arr))
	//排序
	f(arr)
	//排序后
	fmt.Println(arr)
	fmt.Println(len(arr))
}
