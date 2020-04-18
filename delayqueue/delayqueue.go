package delayqueue

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/Allenxuxu/toolkit/priorityqueue"
)

// DelayQueue is an unbounded blocking queue of *Delayed* elements, in which
// an element can only be taken when its delay has expired. The head of the
// queue is the *Delayed* element whose delay expired furthest in the past.
type DelayQueue struct {
	pq    *priorityqueue.PriorityQueue
	exitC chan interface{}
	// Similar to the sleeping state of runtime.timers.
	sleeping int32
	wakeupC  chan interface{}
}

// New creates an instance of delayQueue with the specified size.
func New(size int64) *DelayQueue {
	return &DelayQueue{
		pq:      priorityqueue.New(size),
		exitC:   make(chan interface{}),
		wakeupC: make(chan interface{}),
	}
}

// Offer inserts the element into the current queue.
func (dq *DelayQueue) Offer(value interface{}, expiration time.Time) {
	dq.pq.Push(value, timeToMs(expiration))
	if dq.pq.Len() == 1 {
		if atomic.CompareAndSwapInt32(&dq.sleeping, 1, 0) {
			dq.wakeupC <- struct{}{}
		}
	}
}

// Poll starts an infinite loop, in which it continually waits for an element to
// expire and then send the expired element to the timing wheel via the channel C.
func (dq *DelayQueue) Poll(noticC chan interface{}) {
	for {
		now := timeToMs(time.Now())

		value, p, _ := dq.pq.Pop()

		if value == nil {
			// No items left or at least one item is pending.
			atomic.StoreInt32(&dq.sleeping, 1)
			select {
			case <-dq.wakeupC:
				// Wait until a new item is added.
				continue
			case <-dq.exitC:
				goto exit
			}
		}

		if now <= p {
			fmt.Println("now < p ", now, p)
			// At least one item is pending.
			select {
			case <-dq.wakeupC:
				// A new item with an "earlier" expiration than the current "earliest" one is added.
				continue
			case <-time.After(time.Duration(p-now) * time.Millisecond):
				// The current "earliest" item expires.

				// Reset the sleeping state since there's no need to receive from wakeupC.
				if atomic.SwapInt32(&dq.sleeping, 0) == 0 {
					// A caller of Offer() is being blocked on sending to wakeupC,
					// drain wakeupC to unblock the caller.
					<-dq.wakeupC
				}
				continue
			case <-dq.exitC:
				goto exit
			}
		}

		select {
		case noticC <- value:
			// Send the expired element to the timing wheel.
		case <-dq.exitC:
			goto exit
		}
	}

exit:
	// Reset the states
	atomic.StoreInt32(&dq.sleeping, 0)
}

// Stop poll
func (dq *DelayQueue) Stop() {
	select {
	case <-dq.exitC:
	default:
		close(dq.exitC)
	}
}

func timeToMs(t time.Time) int64 {
	return int64(time.Duration(t.UnixNano()) / time.Millisecond)
}
