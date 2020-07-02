package advancedapplications

import (
	"errors"
	"fmt"
	"time"
)

type Queue struct {
	data []int64
	size int
	head int
	tail int
}

func NewQueue(size int) *Queue {
	return &Queue{
		size: size,
		data: make([]int64, size),
	}
}

func (q *Queue) Add(v int64) error {
	if ((q.tail + 1) % q.size) == q.head {
		return errors.New("full")
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.size
	return nil
}
func (q *Queue) Poll() (res int64) {
	res = q.data[q.head]
	q.head = (q.head + 1) % q.size
	return res
}
func (q *Queue) isEmpty() bool {
	return q.head == q.tail
}

type Producer struct {
	queue *Queue
}

func NewProducer(queue *Queue) *Producer {
	return &Producer{
		queue: queue,
	}
}

func (p *Producer) produce(v int64) {
	for err := p.queue.Add(v); err != nil; err = p.queue.Add(v) {
		time.Sleep(1 * time.Second)
	}
}

type Consumer struct {
	queue *Queue
}

func NewConsumer(queue *Queue) *Consumer {
	return &Consumer{
		queue: queue,
	}
}
func (c *Consumer) comsume() {
	for {
		if c.queue.isEmpty() {
			break
		}
		fmt.Println(c.queue.Poll())
	}
}
