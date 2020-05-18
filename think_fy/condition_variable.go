package think_fy

import (
	"fmt"
	"sync"
)

type Channel struct {
	mut   sync.Mutex
	cond  *sync.Cond
	queue *ArrayQueue
}

func NewChannel(n int) *Channel {
	if n < 1 {
		panic("todo: support unbuffered channel")
	}
	c := new(Channel)
	c.cond = sync.NewCond(&c.mut)
	c.queue = NewArrayQueue(n)
	return c
}

func (c *Channel) TryPush(v interface{}) bool {
	c.mut.Lock()
	defer c.mut.Unlock()
	if c.queue.Full() {
		return false
	}
	if c.queue.Empty() {
		c.cond.Broadcast()
	}
	c.queue.EnQueue(v)
	return true
}

func (c *Channel) TryPop() (interface{}, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	if c.queue.Empty() {
		return nil, false
	}
	if c.queue.Full() {
		c.cond.Broadcast()
	}
	return c.queue.DeQueue(), true
}

func (c *Channel) Push(v interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for c.queue.Full() {
		c.cond.Wait()
	}
	if c.queue.Empty() {
		c.cond.Broadcast()
	}
	c.queue.EnQueue(v)
}

func (c *Channel) Pop() interface{} {
	c.mut.Lock()
	defer c.mut.Unlock()
	for c.queue.Empty() {
		c.cond.Wait()
	}
	if c.queue.Full() {
		c.cond.Broadcast()
	}
	return c.queue.DeQueue()
}

type ArrayQueue struct {
	data []interface{}
	cap  int //容量
	head int //头指针
	tail int //尾指针
}

func NewArrayQueue(n int) *ArrayQueue {
	if n == 0 {
		return nil
	}
	return &ArrayQueue{
		data: make([]interface{}, n),
		cap:  n,
		head: 0,
		tail: 0,
	}
}

func (aq *ArrayQueue) Full() bool {
	return aq.tail == aq.cap
}

func (aq *ArrayQueue) Empty() bool {
	return aq.head == aq.tail
}

func (aq *ArrayQueue) EnQueue(v interface{}) bool {
	if aq.tail == aq.cap {
		return false
	}
	aq.data[aq.tail] = v
	aq.tail++
	return true
}

func (aq *ArrayQueue) DeQueue() interface{} {
	if aq.head == aq.tail {
		return nil
	}
	v := aq.data[aq.head]
	aq.head++
	return v
}

func (aq *ArrayQueue) Print() (format string) {
	if aq.head == aq.tail {
		format = "empty"
	} else {
		format = "head"
		for i := aq.head; i < aq.tail; i++ {
			format += fmt.Sprintf("<-%+v", aq.data[i])
		}
		format += "<-tail"
	}
	return
}
