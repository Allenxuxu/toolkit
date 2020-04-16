package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New()

	if 0 != q.Len() {
		t.Fatal()
	}

	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	if 10 != q.Len() {
		t.Fatal()
	}
	for i := 0; i < 10; i++ {
		v := q.Pop()
		if v.(int) != i {
			t.Fatal()
		}
	}
}
