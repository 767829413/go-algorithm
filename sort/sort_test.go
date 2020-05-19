package sort

import (
	"algorithm/utilcom/test"
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := test.BuildTestArr(500, 1000, 0)
	test.PrintSortComparison(arr, BubbleSort)
}

func TestInsertSort(t *testing.T) {
	arr := test.BuildTestArr(500, 1000, 0)
	test.PrintSortComparison(arr, InsertSort)
}

func TestSelectSort(t *testing.T) {
	arr := test.BuildTestArr(500, 1000, 0)
	test.PrintSortComparison(arr, SelectSort)
}

func TestMergeSort(t *testing.T) {
	arr := test.BuildTestArr(500, 1000, 0)
	test.PrintSortComparison(arr, MergeSort)
	//MergeSort(arr)
}

func TestQuicksort(t *testing.T) {
	arr := test.BuildTestArr(500, 1000, 0)
	test.PrintSortComparison(arr, Quicksort)
	//Quicksort(arr)
}

func TestCountSort(t *testing.T) {
	arr := test.BuildTestArr(500, 1000, 0)
	test.PrintSortComparison(arr, CountSort)
}

func TestBucketSort(t *testing.T) {
	arr := test.BuildTestArr(500, 1000, 0)
	test.PrintSortComparison(arr, BucketSort)
}

func TestBucketSortSimple(t *testing.T) {
	arr := test.BuildTestArr(500, 1000, 0)
	test.PrintSortComparison(arr, BucketSortSimple)
}

func TestFindKMax(t *testing.T) {
	arr := test.BuildTestArr(5, 50, 0)
	fmt.Println(arr)
	fmt.Println(FindKMax(arr, 5))
}
