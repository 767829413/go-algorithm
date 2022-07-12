package main

import (
	"fmt"
	"math"
)

func main() {
	num := int(math.Pow10(0))
	v := 123
	for v > 0 {
		kk := v % 10
		k := v / 10
		fmt.Println(num, v, kk, k)
		v = k
	}

}
