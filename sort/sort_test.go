package sort

import (
	"fmt"
	"testing"

	"github.com/767829413/go-algorithm/utilcom"
)

func TestBubbleSort(t *testing.T) {
	arr := utilcom.BuildTestArr(500, 1000, 0)
	utilcom.PrintSortComparison(arr, BubbleSort)
}

func TestInsertSort(t *testing.T) {
	arr := utilcom.BuildTestArr(500, 1000, 0)
	utilcom.PrintSortComparison(arr, InsertSort)
}

func TestSelectSort(t *testing.T) {
	arr := utilcom.BuildTestArr(500, 1000, 0)
	utilcom.PrintSortComparison(arr, SelectSort)
}

func TestMergeSort(t *testing.T) {
	arr := utilcom.BuildTestArr(500, 1000, 0)
	utilcom.PrintSortComparison(arr, MergeSort)
	//MergeSort(arr)
}

func TestQuicksort(t *testing.T) {
	arr := utilcom.BuildTestArr(500, 1000, 0)
	utilcom.PrintSortComparison(arr, Quicksort)
	//Quicksort(arr)
}

func TestCountSort(t *testing.T) {
	arr := utilcom.BuildTestArr(500, 1000, 0)
	utilcom.PrintSortComparison(arr, CountSort)
}

func TestBucketSort(t *testing.T) {
	arr := utilcom.BuildTestArr(500, 1000, 0)
	utilcom.PrintSortComparison(arr, BucketSort)
}

func TestBucketSortSimple(t *testing.T) {
	arr := utilcom.BuildTestArr(500, 1000, 0)
	utilcom.PrintSortComparison(arr, BucketSortSimple)
}

func TestHeapSort(t *testing.T) {
	arr := utilcom.BuildTestArr(500, 1000, 0)
	arr = append([]int{0}, arr...)
	utilcom.PrintSortComparison(arr, HeapSort)
}

func TestFindKMax(t *testing.T) {
	arr := utilcom.BuildTestArr(5, 50, 0)
	fmt.Println(arr)
	fmt.Println(FindKMax(arr, 5))
}
