package main

import "fmt"

func main() {
	str := "Assfsdgg"
	dd := []rune(str)
	for _, v := range dd {
		fmt.Println(v)
	}
}
