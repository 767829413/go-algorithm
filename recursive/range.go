package recursive

import "fmt"

type RangeType struct {
	data []interface{}
}

func NewRangeType(n int) *RangeType {
	return &RangeType{
		data: make([]interface{}, n),
	}
}

func (slice *RangeType) RangeALL(start int) {
	length := len(slice.data)
	if start == length-1 {
		// 如果已经是最后位置，直接将数组数据合并输出
		fmt.Println(slice.data)
	}

	for i := start; i < length; i++ {
		// i = start 时输出自己
		// 如果i和start的值相同就没有必要交换
		if i == start || slice.data[i] != slice.data[start] {
			//交换当前这个与后面的位置
			slice.data[i], slice.data[start] = slice.data[start], slice.data[i]
			//递归处理索引+1
			slice.RangeALL(start + 1)
			//换回来，因为是递归，如果不换回来会影响后面的操作，并且出现重复
			slice.data[i], slice.data[start] = slice.data[start], slice.data[i]
		}
	}
}
