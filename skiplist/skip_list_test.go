package skiplist

import "testing"

func TestSkipList(t *testing.T) {
	sl := NewSkipList()
	for i := 1; i < 7; i++ {
		sl.Insert(i)
	}
	t.Log(sl.Print())
	t.Log(sl.Delete(5))
	t.Log(sl.Print())
	t.Log(sl.Delete(3))
	t.Log(sl.Print())
	t.Log(sl.Delete(3))
	t.Log(sl.Print())
}
