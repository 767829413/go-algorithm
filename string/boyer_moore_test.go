package string

import (
	"fmt"
	"testing"
)

func TestBoyerMooreMatch(t *testing.T) {
	str1 := "agggdfhdgghdeertgvc xfsfggafafd"
	str2 := "g"
	res := BoyerMooreMatch([]rune(str1), []rune(str2), len(str1), len(str2))
	t.Log(res)
}

func TestBruteForceDebug(t *testing.T) {
	str := "cabcab"
	fmt.Println(str)
	t.Log(generatePre([]rune(str), len(str)))
}
