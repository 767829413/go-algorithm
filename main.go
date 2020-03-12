package main

import "fmt"

func main() {
	eg1()
}

func eg1() {
	// array表示一个长度为n的数组
	// 代码中的array.length就等于n
	arr := [2]int{}
	count := 0
	for {
		insert(&arr, &count, 100)
		if arr[0] > 500000{
			break
		}
	}
	fmt.Println(count)
	fmt.Println(arr)

}

func insert(arr *[2]int, count *int, val int) {
	if *count == len(arr) {
		sum := 0
		for i := 0; i < len(arr); i++ {
			sum = sum + arr[i]
		}
		arr[0] = sum
		*count = 1
	}
	arr[*count] = val
	*count++
}
