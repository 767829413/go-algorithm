package divideconquer

import (
	"testing"

	"github.com/767829413/go-algorithm/utilcom"
)

func TestCount(t *testing.T) {
	arr := utilcom.BuildTestArr(5, 1000, 0)
	t.Log(arr)
	t.Log(Count(arr, len(arr)))
	t.Log(arr)
}
