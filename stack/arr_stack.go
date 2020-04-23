package stack

import "fmt"

/**
基于数组实现
 */

type ArrayStack struct {
	data []interface{}
	top  int
}

//初始化一个栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

//是否为空栈
func (as *ArrayStack) IsEmpty() bool {
	if as.top < 0 {
		return true
	}
	return false
}

//压栈
func (as *ArrayStack) Push(v interface{}) {
	if as.top < 0 {
		as.top = 0
	} else {
		as.top++
	}
	if as.top > len(as.data)-1 {
		as.data = append(as.data, v)
	} else {
		as.data[as.top] = v
	}
}

//出栈
func (as *ArrayStack) Pop() interface{} {
	if as.IsEmpty() {
		return nil
	}
	v := as.data[as.top]
	as.top--
	return v
}

//获取栈顶数据
func (as *ArrayStack) Top() interface{} {
	if as.IsEmpty() {
		return nil
	}
	return as.data[as.top]
}

//清空栈
func (as *ArrayStack) Flush() {
	as.top = -1
}

//打印栈
func (as *ArrayStack) Print() (format string) {
	if as.IsEmpty() {
		format = "empty"
	} else {
		format = ""
		for i := as.top; i >= 0; i-- {
			format += fmt.Sprintf("%+v", as.data[i])
			format += "->"
		}
		format += "end"
	}
	return
}
