package main

import "fmt"

type sss struct {
	name string
	age  int
}

func main() {
	fmt.Println(int(^uint(0) >> 1))
	fmt.Println(^int(^uint(0) >> 1))
	fmt.Println(^int(^uint(0) >> 1)-1)
}
