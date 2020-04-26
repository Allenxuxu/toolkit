package queue

import "container/list"

type Queue interface {
	Push(v interface{})
	Pop() interface{}
	Len() int
}

type queue struct {
	list *list.List
}

func New() Queue {
	return &queue{
		list: list.New(),
	}
}

func (q *queue) Push(v interface{}) {
	q.list.PushBack(v)
}

func (q *queue) Pop() interface{} {
	v := q.list.Front()
	if v == nil {
		return nil
	}

	return q.list.Remove(v)
}

func (q *queue) Peek() interface{} {
	v := q.list.Front()
	if v == nil {
		return nil
	}

	return v.Value
}

func (q *queue) Len() int {
	return q.list.Len()
}
