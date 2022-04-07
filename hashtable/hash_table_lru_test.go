package hashtable

import (
	"testing"

	"github.com/767829413/go-algorithm/utilcom/randomstring"
	"github.com/stretchr/testify/assert"
)

func TestLRUHashTableOp(t *testing.T) {
	capacity := 10
	testData := make(map[interface{}]interface{}, capacity)
	lru := NewLRUHashTable(capacity)
	t.Log(lru.Print())
	for i := 0; i < capacity; i++ {
		v := randomstring.RandStringBytesMaskImprSrc(5)
		testData[v] = i
		lru.Add(v, i)
	}
	t.Log(lru.Print())
	for k, v := range testData {
		assert.Equal(t, v, lru.Get(k))
		lru.Delete(k)
		assert.Equal(t, nil, lru.Get(k))
	}
	t.Log(lru.Print())
}

func TestLRUHashTableGet(t *testing.T) {
	capacity := 10
	lru := NewLRUHashTable(capacity)
	lru.Add("sdsd", "1")
	lru.Add("rtrtrt", "2")
	lru.Add("bnnbnbn", "3")
	lru.Add("xcxcxs", "4")
	t.Log(lru.Print())
	t.Log(lru.Get("sdsd"))
	t.Log(lru.Print())
	t.Log(lru.Get("bnnbnbn"))
	t.Log(lru.Print())
}
