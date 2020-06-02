package string

import "testing"

func TestAcTrie(t *testing.T) {
	ac := NewAcTrie()
	str1 := "他妈的"
	str2 := "这他妈的"
	ac.Insert([]rune(str1))
	ac.Insert([]rune(str2))
	ac.BuildFailPoint()
	ac.match([]rune("你告诉我什么叫他妈的惊喜?还是这就叫这他妈的惊喜"))
}
