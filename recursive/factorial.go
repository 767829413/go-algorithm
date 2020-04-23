package recursive

/**
阶乘
 */
type Factorial struct {
	data map[int]int
}

func NewFactorial(n int) *Factorial {
	return &Factorial{
		data: make(map[int]int, n),
	}
}

func (f *Factorial) Recursive(n int) int {
	if n < 1 {
		f.data[n] = 1
		return 1
	} else {
		res := n * f.Recursive(n-1)
		f.data[n] = res
		return res
	}
}

func (f *Factorial) TailRecursive(n int) int {
	return f.tail(1, n)
}

func (f *Factorial) tail(a, n int) int {
	if n < 1 {
		return a
	}
	return f.tail(n*a, n-1)
}

func (f *Factorial) Loop(n int) int {
	var a = 1
	for ; n > 1; n-- {
		a = a * n
	}
	return a
}
