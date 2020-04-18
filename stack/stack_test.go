package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()

	if 0 != s.Len() {
		t.Fatal()
	}

	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	if 10 != s.Len() {
		t.Fatal()
	}
	for i := 9; i >= 0; i-- {
		v := s.Pop()
		if v.(int) != i {
			t.Fatal()
		}
	}
}

func TestStack_Pop(t *testing.T) {
	q := New()

	if q.Pop() != nil {
		t.Fatal()
	}
}
