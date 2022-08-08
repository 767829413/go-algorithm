package struct

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewChannel(t *testing.T) {
	num := 100000
	c := NewChannel(num)
	wg := sync.WaitGroup{}
	for i := 0; i < num; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			c.Push(i)
		}(i)

		go func() {
			defer wg.Done()
			fmt.Println(c.Pop())
		}()
	}
	wg.Wait()
}
