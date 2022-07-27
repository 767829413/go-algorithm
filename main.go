package main

import "fmt"

func main() {
	a := make([]int, 3)

	a[0] = '1'
	a[1] = '2'
	a[2] = '3'

	fmt.Println("final", int(a[0]+a[1]))
}
