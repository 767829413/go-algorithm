package main

import "fmt"

var (
	data = []int{5, 2, 4, 3, 1, 6}
)

func main() {
	for i := 1; i < len(data)-1; i++ {
		val := data[i]
		index := i - 1
		for index >= 0 && val < data[index] {
			data[index+1] = data[index]
			index--

		}
		data[index+1] = val
	}
	fmt.Println(data)
}
