package main

import "fmt"

type Demo struct {
	data interface{}
}

func main() {
	demo := Demo{}
	//demo.data = "sdsds"
	demo.data = 100
	if ff, ok := demo.data.(string); ok {
		fmt.Println("这是string")
		fmt.Println(ff)
	} else {
		fmt.Println("这不是string")
		fmt.Println(ff)
	}
}
