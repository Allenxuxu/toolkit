package priorityqueue

import "testing"

func TestNew(t *testing.T) {
	pq := New(10)
	if pq.Len() != 0 {
		t.Fatal()
	}

	_, _, exist := pq.Peek()
	if exist != false {
		t.Fatal()
	}
	_, _, exist = pq.Pop()
	if exist != false {
		t.Fatal()
	}

	pq.Push("hello", 1)
	v, n, exist := pq.Peek()
	if exist != true || v.(string) != "hello" || n != 1 {
		t.Fatal()
	}
	v, n, exist = pq.Pop()
	if exist != true || v.(string) != "hello" || n != 1 {
		t.Fatal()
	}

	pq.Push("0", 0)
	pq.Push("1", 1)
	pq.Push("2", 2)
	pq.Push("-1", -1)
	v, n, exist = pq.Pop()
	if exist != true || v.(string) != "-1" || n != -1 {
		t.Fatal()
	}
	v, n, exist = pq.Pop()
	if exist != true || v.(string) != "0" || n != 0 {
		t.Fatal()
	}
	v, n, exist = pq.Pop()
	if exist != true || v.(string) != "1" || n != 1 {
		t.Fatal()
	}
	v, n, exist = pq.Pop()
	if exist != true || v.(string) != "2" || n != 2 {
		t.Fatal(v)
	}

	if pq.Len() != 0 {
		t.Fatal()
	}
}
