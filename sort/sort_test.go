package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := BuildTestArr(5, 50, 0)
	PrintSortComparison(arr, BubbleSort)
}

func TestInsertSort(t *testing.T) {
	arr := BuildTestArr(5, 50, 0)
	PrintSortComparison(arr, InsertSort)
}

func TestSelectSort(t *testing.T) {
	arr := BuildTestArr(5, 50, 0)
	PrintSortComparison(arr, SelectSort)
}

func TestMergeSort(t *testing.T) {
	arr := BuildTestArr(5, 50, 0)
	PrintSortComparison(arr, MergeSort)
}

func TestQuicksort(t *testing.T) {
	arr := BuildTestArr(5, 50, 0)
	PrintSortComparison(arr, Quicksort)
}

func TestFindKMax(t *testing.T) {
	arr := BuildTestArr(5, 50, 0)
	fmt.Println(arr)
	fmt.Println(FindKMax(arr, 5))
}

func TestCountSort(t *testing.T) {
	arr := BuildTestArr(5, 50, 0)
	PrintSortComparison(arr, CountSort)
}
