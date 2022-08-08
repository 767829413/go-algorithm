package handle

import (
	"fmt"
	"sync"
)

type MyChannel struct {
	mut   sync.Mutex
	cond  *sync.Cond
	queue *MyArrayQueue
}

func NewMyChannel(n int) *MyChannel {
	if n < 1 {
		panic("todo: support unbuffered channel")
	}
	c := new(MyChannel)
	c.cond = sync.NewCond(&c.mut)
	c.queue = NewMyArrayQueue(n)
	return c
}

func (c *MyChannel) TryPush(v interface{}) bool {
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

func (c *MyChannel) TryPop() (interface{}, bool) {
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

func (c *MyChannel) Push(v interface{}) {
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

func (c *MyChannel) Pop() interface{} {
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

type MyArrayQueue struct {
	data []interface{}
	cap  int //容量
	head int //头指针
	tail int //尾指针
}

func NewMyArrayQueue(n int) *MyArrayQueue {
	if n == 0 {
		return nil
	}
	return &MyArrayQueue{
		data: make([]interface{}, n),
		cap:  n,
		head: 0,
		tail: 0,
	}
}

func (aq *MyArrayQueue) Full() bool {
	return aq.tail == aq.cap
}

func (aq *MyArrayQueue) Empty() bool {
	return aq.head == aq.tail
}

func (aq *MyArrayQueue) EnQueue(v interface{}) bool {
	if aq.tail == aq.cap {
		return false
	}
	aq.data[aq.tail] = v
	aq.tail++
	return true
}

func (aq *MyArrayQueue) DeQueue() interface{} {
	if aq.head == aq.tail {
		return nil
	}
	v := aq.data[aq.head]
	aq.head++
	return v
}

func (aq *MyArrayQueue) Print() (format string) {
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
