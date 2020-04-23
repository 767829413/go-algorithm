package recursive

/**
斐波那契数列
 */
type Fibonacci struct {
	data map[int]int
}

func NewFibonacci(n int) *Fibonacci {
	return &Fibonacci{
		data: make(map[int]int, n),
	}
}

func (f *Fibonacci) Recursive(n int) int {
	if v, ok := f.data[n]; ok {
		return v
	}
	if n <= 1 {
		f.data[1] = 1
		return 1
	} else if n == 2 {
		f.data[2] = 1
		return 1
	} else {
		res := f.Recursive(n-1) + f.Recursive(n-2)
		f.data[n] = res
		return res
	}

}

func (f *Fibonacci) TailRecursive(n int) int {
	return f.tail(1, 1, n)
}

func (f *Fibonacci) tail(left int, right int, n int) int {
	if n < 2 {
		return left
	}
	return f.tail(right, left+right, n-1)
}

func (f *Fibonacci) Loop(n int) int {
	var a, b = 1, 1
	for ; n > 1; n-- {
		a, b = b, a+b
	}
	return a
}
