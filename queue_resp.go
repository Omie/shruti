package main

import "sync"

type TTSResponse struct {
	Text     string
	Audio    []byte
	Provider string
	KeepFile bool
}

type TTSResponseQueueNode struct {
	*TTSResponse
	Next *TTSResponseQueueNode
}

//	A go-routine safe FIFO (first in first out) data stucture.
type TTSResponseQueue struct {
	Head  *TTSResponseQueueNode
	Tail  *TTSResponseQueueNode
	Count int
	Lock  *sync.Mutex
}

//	Creates a new pointer to a new queue.
func NewTTSResponseQueue() *TTSResponseQueue {
	q := &TTSResponseQueue{}
	q.Lock = &sync.Mutex{}
	return q
}

//	Returns the number of elements in the queue (i.e. size/length)
//	go-routine safe.
func (q *TTSResponseQueue) Len() int {
	q.Lock.Lock()
	defer q.Lock.Unlock()
	return q.Count
}

//	Pushes/inserts a value at the end/Tail of the queue.
//	Note: this function does mutate the queue.
//	go-routine safe.
func (q *TTSResponseQueue) Push(item *TTSResponse) {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	n := &TTSResponseQueueNode{TTSResponse: item}

	if q.Tail == nil {
		q.Tail = n
		q.Head = n
	} else {
		q.Tail.Next = n
		q.Tail = n
	}
	q.Count++
}

//	Returns the value at the front of the queue.
//	i.e. the oldest value in the queue.
//	Note: this function does mutate the queue.
//	go-routine safe.
func (q *TTSResponseQueue) Poll() *TTSResponse {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	if q.Head == nil {
		return nil
	}

	n := q.Head
	q.Head = n.Next

	if q.Head == nil {
		q.Tail = nil
	}
	q.Count--

	return n.TTSResponse
}

//	Returns a read value at the front of the queue.
//	i.e. the oldest value in the queue.
//	Note: this function does NOT mutate the queue.
//	go-routine safe.
func (q *TTSResponseQueue) Peek() *TTSResponse {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	n := q.Head
	if n == nil || n.TTSResponse == nil {
		return nil
	}

	return n.TTSResponse
}
