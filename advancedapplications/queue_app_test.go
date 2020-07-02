package advancedapplications

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestQueueApp(t *testing.T) {
	q := NewQueue(10)
	c := NewConsumer(q)
	p := NewProducer(q)
	wg := &sync.WaitGroup{}
	i := int64(0)
	for ; i < 20; i++ {
		wg.Add(2)
		go func(i int64) {
			defer wg.Done()
			p.produce(i)
		}(i)
		go func() {
			defer wg.Done()
			c.comsume()
		}()
	}
	wg.Wait()
}

func TestChannel(t *testing.T) {
	d := time.Duration(time.Second * 15)
	tt := time.NewTimer(d)
	for {
		<-tt.C
		break
	}
	fmt.Println("end")
}
