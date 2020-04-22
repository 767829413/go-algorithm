package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func buildTestData() interface{} {
	return []map[string]interface{}{
		map[string]interface{}{
			"expected": true,
			"data":     []string{"a", "r", "s", "x", "s", "r", "a"},
		},
		map[string]interface{}{
			"expected": true,
			"data":     []string{"a"},
		},
		map[string]interface{}{
			"expected": true,
			"data":     []string{"C", "C"},
		},
		map[string]interface{}{
			"expected": true,
			"data":     []string{"a", "r", "s", "s", "r", "a"},
		},
		map[string]interface{}{
			"expected": false,
			"data":     []string{"a", "b", "c", "d", "e"},
		},
		map[string]interface{}{
			"expected": false,
			"data":     []string{},
		},
	}
}
func TestIsPalindromeByStack(t *testing.T) {
	datas := buildTestData().([]map[string]interface{})
	for _, v := range datas {
		l := NewLinkedList()
		for _, i := range v["data"].([]string) {
			l.InsertToTail(i)
		}
		assert.Equal(t, v["expected"], IsPalindromeByStack(l), l.Print())
	}
}

func TestIsPalindromeByFastLowPoint(t *testing.T) {
	datas := buildTestData().([]map[string]interface{})
	for _, v := range datas {
		l := NewLinkedList()
		for _, i := range v["data"].([]string) {
			l.InsertToTail(i)
		}
		assert.Equal(t, v["expected"], IsPalindromeByFastLowPoint(l), l.Print())
	}
}

func TestDFDF(t *testing.T) {
	arr := []string{"a", "r", "s", "x", "s", "r", "a"}
	//arr := []string{"a", "b", "c", "d", "e", "f", "g"}
	l := NewLinkedList()
	for _, v := range arr {
		l.InsertToTail(v)
	}
	t.Log(l.Print())
	t.Log(IsPalindromeByFastLowPoint(l))
	t.Log(l.Print())
}
