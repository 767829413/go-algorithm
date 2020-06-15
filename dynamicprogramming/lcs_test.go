package dynamicprogramming

import "testing"

func TestLCS(t *testing.T) {
	aPart := []rune("mitcmu")
	bPart := []rune("mtacnu")
	l := NewLCS(aPart, bPart, len(aPart), len(bPart))
	t.Log(l.GetMaxDest())
}

func TestLCS1(t *testing.T) {
	aPart := []rune("mitcmu")
	bPart := []rune("mtacnu")
	l := NewLCS(aPart, bPart, len(aPart), len(bPart))
	t.Log(l.GetMaxDest1())
}
