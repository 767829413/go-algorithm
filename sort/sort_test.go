package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := BuildTestArr(10, 63, 0)
	PrintSortComparison(arr, BubbleSort)
}

func TestInsertSort(t *testing.T) {
	arr := BuildTestArr(10, 63, 0)
	PrintSortComparison(arr, InsertSort)
}

func TestSelectSort(t *testing.T) {
	arr := BuildTestArr(10, 63, 0)
	PrintSortComparison(arr, SelectSort)
}

func TestMergeSort(t *testing.T) {
	arr := BuildTestArr(10, 63, 0)
	PrintSortComparison(arr, MergeSort)
}
