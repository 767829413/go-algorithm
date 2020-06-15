package dynamicprogramming

import "testing"

func TestLevenshtein_ModifyDest(t *testing.T) {
	aPart := []rune("mitcmu")
	bPart := []rune("mtacnu")
	l := NewLevenshtein(aPart, bPart, len(aPart), len(bPart))
	t.Log(l.GetMinDest())
}

func TestLevenshtein_ModifyDest1(t *testing.T) {
	aPart := []rune("mitcmu")
	bPart := []rune("mtacnu")
	l := NewLevenshtein(aPart, bPart, len(aPart), len(bPart))
	t.Log(l.GetMinDest1())
}
