package main

import "fmt"

func main() {
	a := make([]int, 3)

	del(a)

	fmt.Println("final", a)
}

func add(a []int) {
	a = append(a, 999)
	fmt.Println("add", a)
}

func del(a []int) {
	a = append(a[:1], a[2:]...)
	fmt.Println("del", a)
}
