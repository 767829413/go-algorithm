package binarySearch

import (
	"algorithm/sort"
	"algorithm/utilcom/test"
	"fmt"
	"testing"
)

type tf func(arr []int, find int, index *int)

func printTestRes(arr []int, find int, index int, f tf) {
	f(arr, find, &index)
	fmt.Println("有序数组集合: ", arr)
	fmt.Println("查找的数: ", find)
	fmt.Println("所在位置: ", index)
}

func TestEasyLoopBinarySearch(t *testing.T) {
	arr := test.BuildTestArr(10, 20, 0)
	find := arr[3]
	index := -1
	sort.Quicksort(arr)
	printTestRes(arr, find, index, EasyLoopBinarySearch)
}

func TestEasyRecursiveBinarySearch(t *testing.T) {
	arr := test.BuildTestArr(10, 20, 0)
	find := arr[3]
	index := -1
	sort.Quicksort(arr)
	printTestRes(arr, find, index, EasyRecursiveBinarySearch)
}

func TestFindFirstEqual(t *testing.T) {
	arr := []int{0, 0, 1, 2, 2, 3, 4, 5, 5, 5, 7, 15, 16, 19}
	find := 2
	index := -1
	printTestRes(arr, find, index, FindFirstEqual)
}

func TestFindEndEqual(t *testing.T) {
	arr := []int{0, 0, 1, 2, 2, 3, 4, 5, 5, 5, 5, 5, 7, 15, 16, 19}
	find := 5
	index := -1
	printTestRes(arr, find, index, FindEndEqual)
}

func TestFindFirstGt(t *testing.T) {
	arr := []int{4, 4, 5, 5, 6, 7, 8, 9, 9, 9, 9, 9, 12, 12, 14, 15, 16, 19}
	find := 13
	index := -1
	printTestRes(arr, find, index, FindFirstGe)
}

func TestFindEndGe(t *testing.T) {
	arr := []int{4, 4, 5, 5, 8, 9, 9, 9, 9, 9, 12, 12, 14, 15, 16, 19} //count : 16 max index 15
	find := 13
	index := -1
	printTestRes(arr, find, index, FindEndGe)
}

func TestFindFirstLe(t *testing.T) {
	arr := []int{4, 4, 6, 6, 8, 9, 9, 9, 9, 9, 12, 12, 14, 15, 16, 19} //count : 16 max index 15
	find := 4
	index := -1
	printTestRes(arr, find, index, FindFirstLe)
}

func TestFindEndLe(t *testing.T) {
	arr := []int{4, 4, 6, 6, 8, 9, 9, 9, 9, 9, 12, 12, 14, 15, 16, 19} //count : 16 max index 15
	find := 7
	index := -1
	printTestRes(arr, find, index, FindEndLe)
}
