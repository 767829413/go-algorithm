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
func (this *ArrayStack) IsEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

//压栈
func (this *ArrayStack) Push(v interface{}) {
	if this.top < 0 {
		this.top = 0
	} else {
		this.top++
	}
	if this.top > len(this.data)-1 {
		this.data = append(this.data, v)
	} else {
		this.data[this.top] = v
	}
}

//出栈
func (this *ArrayStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.data[this.top]
	this.top--
	return v
}

//获取栈顶数据
func (this *ArrayStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.top]
}

//清空栈
func (this *ArrayStack) Flush() {
	this.top = -1
}

//打印栈
func (this *ArrayStack) Print() (format string) {
	if this.IsEmpty() {
		format = "empty"
	} else {
		format = ""
		for i := this.top; i >= 0; i-- {
			format += fmt.Sprintf("%+v", this.data[i])
			format += "->"
		}
		format += "end"
	}
	return
}
