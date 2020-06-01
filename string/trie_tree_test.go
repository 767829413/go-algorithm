package string

import (
	"testing"
)

func TestTrie(t *testing.T) {
	testStrs := []string{
		"人生如初见",
		"人就考很快乐回事",
		"xcfs;xflgksdkg;lk",
		"ewuieywruyrwhm学科划分SD卡",
	}
	tt := NewTrie()
	for _, v := range testStrs {
		tt.Insert([]rune(v))
	}
	for i := len(testStrs) - 1; i >= 0; i-- {
		t.Log(tt.Find([]rune(testStrs[i])))
	}
}
