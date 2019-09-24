package priorityqueue

import (
	"container/heap"
	"sync"
)

type item struct {
	Value    interface{}
	Priority int64
	Index    int
}

type priorityQueue []*item

func newPriorityQueue(capacity int64) priorityQueue {
	return make(priorityQueue, 0, capacity)
}

func (pq *priorityQueue) Len() int {
	return len(*pq)
}

func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i].Priority < (*pq)[j].Priority
}

func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].Index = i
	(*pq)[j].Index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	c := cap(*pq)
	if n+1 > c {
		npq := make(priorityQueue, n, c*2)
		copy(npq, *pq)
		*pq = npq
	}
	*pq = (*pq)[0 : n+1]
	item := x.(*item)
	item.Index = n
	(*pq)[n] = item
}

func (pq *priorityQueue) Pop() interface{} {
	n := len(*pq)
	c := cap(*pq)
	if n < (c/2) && c > 25 {
		npq := make(priorityQueue, n, c/2)
		copy(npq, *pq)
		*pq = npq
	}
	item := (*pq)[n-1]
	item.Index = -1
	*pq = (*pq)[0 : n-1]
	return item
}

func (pq *priorityQueue) PeekAndShift(max int64) (*item, int64) {
	if pq.Len() == 0 {
		return nil, 0
	}

	item := (*pq)[0]
	if item.Priority > max {
		return nil, item.Priority - max
	}
	heap.Remove(pq, 0)

	return item, 0
}

// PriorityQueue 优先级队列
type PriorityQueue struct {
	mu sync.Mutex
	pq priorityQueue
}

// New 创建一个容量为 capacity 的优先级队列
func New(capacity int64) *PriorityQueue {
	return &PriorityQueue{pq: newPriorityQueue(capacity)}
}

// Len 队列长度
func (p *PriorityQueue) Len() int {
	return p.pq.Len()
}

// Push 插入一个元素
func (p *PriorityQueue) Push(value interface{}, Priority int64) {
	p.mu.Lock()
	heap.Push(&p.pq, &item{Value: value, Priority: Priority})
	p.mu.Unlock()
}

// Pop 弹出一个优先级最小的元素
func (p *PriorityQueue) Pop() (interface{}, int64, bool) {
	if p.pq.Len() == 0 {
		return nil, 0, false
	}

	p.mu.Lock()
	e := heap.Pop(&p.pq)
	p.mu.Unlock()
	item := e.(*item)
	return item.Value, item.Priority, true
}

// Peek 窥视一个优先级最小的元素
func (p *PriorityQueue) Peek() (interface{}, int64, bool) {
	if p.pq.Len() == 0 {
		return nil, 0, false
	}

	p.mu.Lock()
	item := p.pq[0]
	p.mu.Unlock()
	return item.Value, item.Priority, true
}
