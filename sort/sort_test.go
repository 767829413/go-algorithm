package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := BuildTestArr(6, 10, 0)
	//排序前
	fmt.Println(arr)
	//排序
	BubbleSort(arr)
	//排序后
	fmt.Println(arr)
}

func TestInsertSort(t *testing.T) {
	arr := BuildTestArr(6, 10, 0)
	//排序前
	fmt.Println(arr)
	//排序
	InsertSort(arr)
	//排序后
	fmt.Println(arr)
}

func TestSelectSort(t *testing.T) {
	arr := BuildTestArr(6, 10, 0)
	//排序前
	fmt.Println(arr)
	//排序
	SelectSort(arr)
	//排序后
	fmt.Println(arr)
}
