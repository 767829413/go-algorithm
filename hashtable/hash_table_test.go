package hashtable

import (
	"algorithm/utilcom/randomstring"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashTable(t *testing.T) {
	num := 10000
	ht := NewHashTable()
	testData := make(map[string]int, num)
	for i := 0; i < num; i++ {
		index := randomstring.RandStringBytesMaskImprSrc(3)
		testData[index] = i
		ht.Put(index, i)
	}
	for k, v := range testData {
		assert.Equal(t, v, ht.Find(k), "find")
		ht.Delete(k)
		assert.Equal(t, nil, ht.Find(k), "delete after")
	}
}
