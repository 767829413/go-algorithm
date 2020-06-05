package main

import "fmt"

func main() {
	var d int
	for i := 0; i < 8; i++ {
		d = d << 1
		fmt.Println(d)
	}
}
