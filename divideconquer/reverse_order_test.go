package divideconquer

import (
	"algorithm/utilcom/test"
	"testing"
)

func TestCount(t *testing.T) {
	arr := test.BuildTestArr(5, 1000, 0)
	t.Log(arr)
	t.Log(Count(arr, len(arr)))
	t.Log(arr)
}
