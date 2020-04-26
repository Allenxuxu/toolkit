package stack

import "container/list"

type Stack interface {
	Push(v interface{})
	Pop() interface{}
	Len() int
}

type stack struct {
	list *list.List
}

func New() Stack {
	return &stack{
		list: list.New(),
	}
}
func (s *stack) Push(v interface{}) {
	s.list.PushBack(v)
}

func (s *stack) Pop() interface{} {
	v := s.list.Back()
	if v == nil {
		return nil
	}

	return s.list.Remove(v)
}

func (s *stack) Peek() interface{} {
	v := s.list.Back()
	if v == nil {
		return nil
	}

	return v.Value
}

func (s *stack) Len() int {
	return s.list.Len()
}
