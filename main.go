package main

func main() {
	var _ Test = (*ImplT1)(nil)
}

type Test interface {
	A() string
}

type ImplT1 struct{}

// func (i *ImplT1) A() string { return "" }
