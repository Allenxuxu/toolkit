package delayqueue

import (
	"fmt"
	"testing"
	"time"
)

type Entry struct {
	Key  int
	Time time.Time
}

func TestDelayQueue_Poll(t *testing.T) {
	var layout = "2006-01-02T15:04:05.999999"
	dq := New(10)

	noticC := make(chan interface{})
	go dq.Poll(noticC)

	var array = []Entry{
		{Key: 1},
		//{Key: 2},
		//{Key: 3},
		//{Key: 4},
		//{Key: 5},
		//{Key: 6},
		//{Key: 7},
	}
	tt := time.Now().Add(time.Second * 4)
	for i := 0; i < len(array); i++ {
		tmp := tt.Add(time.Duration(array[i].Key) * time.Millisecond)
		array[i].Time = tmp
		dq.Offer(array[i], tmp)
	}

	for i := 0; i < len(array); i++ {
		select {
		case run := <-noticC:
			if array[i].Key != run.(Entry).Key {
				t.Fatal(array[i].Key, i)
			}
			t.Log(array[i].Key, array[i].Time.Format(layout))
			t.Log(array[i].Key, time.Now().Format(layout))
		}
	}
	dq.Stop()
}

func TestDelayQueue_Stop(t *testing.T) {
	dq := New(10)

	noticC := make(chan interface{})
	go dq.Poll(noticC)

	time.Sleep(time.Millisecond * 100)

	dq.Stop()
	dq.Stop()
	dq.Stop()
}

func ExampleDelayQueue_Offer() {
	dq := New(1)

	noticC := make(chan interface{})
	go dq.Poll(noticC)
	var array = []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(array); i++ {
		dq.Offer(array[i], time.Now().Add(time.Duration(array[i])*time.Millisecond))
	}

	for i := 0; i < len(array); i++ {
		select {
		case run := <-noticC:
			fmt.Println(run.(int))
		}
	}
	dq.Stop()

	// Output:
	//1
	//2
	//3
	//4
	//5
	//6
	//7
}
