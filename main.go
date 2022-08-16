package main

func main() {
	var _ Test = (*ImplT1)(nil)
	var res [][]int
	for i := 0; i < 10; i++ {
		if res == nil {
			res = make([][]int, 0)
		}
	}
}

type Test interface {
	A() string
}

type ImplT1 struct{}

func (i *ImplT1) A() string { return "" }
