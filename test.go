package main

import "fmt"

func main() {
	var data [][]string
	data = [][]string{

	}
	unique := make(map[string]bool)
	for _, v := range data {
		if unique[v[0]+v[1]] == false {
			unique[v[0]+v[1]] = true
		} else {
			fmt.Println("重复的appkey channel")
			fmt.Println(v[0])
			fmt.Println(v[1])
		}

	}
}
