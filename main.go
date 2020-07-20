package main

import (
	"fmt"
	"time"
)

type sss struct {
	name string
	age  int
}

func main() {

}

func testa() {
	start1 := time.Now()
	for i := 0; i < 100; i++ {
		for i := 0; i < 1000; i++ {
			for i := 0; i < 10000; i++ {
			}
		}
	}
	fmt.Println(time.Since(start1))
	start := time.Now()
	for i := 0; i < 10000; i++ {
		for i := 0; i < 1000; i++ {
			for i := 0; i < 100; i++ {
			}
		}
	}
	fmt.Println(time.Since(start))
}

func testb() {
	for i := 0; i < 10000; i++ {
		for i := 0; i < 1000; i++ {
			for i := 0; i < 100; i++ {
			}
		}
	}
}

func test1() {
	var a float32
	var b float32
	a = 0
	b = 1
	for i := 0; i < 20000000; i++ {
		a += b
	}
	fmt.Println(a)
}

func test2() {
	var a float32
	var b float32
	var c float32
	var t float32
	a = 0
	b = 1
	c = 0
	for i := 0; i < 20000000; i++ {
		y := b - c
		t = a + y
		c = (t - a) - y
		a = t
		fmt.Println(c)
	}
	fmt.Println(t)
}
