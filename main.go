package main

import "fmt"

func main() {
	//print(3 % 2)
	//print("\n")
	//print(3 & (2 - 1))

	//k := randomstring.RandStringBytesMaskImprSrc(3)
	//k := "ssdf"
	//num := xxhash.ChecksumString32(k)
	//fmt.Println(k)
	//fmt.Println((num ^ (num >> 16)) & (uint32(10 - 1)))
	d1 := []int{0, 2, 4, 6, 8}
	d2 := make([]int, len(d1)*2)
	copy(d2, d1)
	fmt.Println(d2)
}
