package queue

type QueueLoop struct {
	Items []string
	Count int
	Head  int
	Tail  int
}

func NewQueueLoop(n int) *QueueLoop {
	items := []string{}
	for i := 0; i < n; i++ {
		items = append(items, "")
	}
	return &QueueLoop{
		Items: items,
		Count: n,
		Head:  0,
		Tail:  0,
	}
}

//入队
func (q *QueueLoop) EnQueue(item string) bool {
	if (q.Tail+1)%q.Count == q.Head {
		return false
	}
	q.Items[q.Tail] = item
	q.Tail = (q.Tail + 1) % q.Count
	return true
}

//出队
func (q *QueueLoop) DeQueue() (item string) {
	if q.Head == q.Tail {
		return
	}
	item = q.Items[q.Head]
	q.Head = (q.Head + 1) % q.Count
	return
}
