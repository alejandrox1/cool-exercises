/*
Here we include a regular queue and a priority queue.
Code taken from https://github.com/golang-collections/go-datastructures

The queue uses WaitGroups to pause listening threads on empty queues until a
message is received.
*/
package main

import (
	"errors"
	"runtime"
	"sync"
	"sync/atomic"
)

type waiters []*sema
type items []interface{}

type sema struct {
	wg       *sync.WaitGroup
	response *sync.WaitGroup
}

func newSema() *sema {
	return &sema{
		wg:       &sync.WaitGroup{},
		response: &sync.WaitGroup{},
	}
}

// get returns and removes the first waiter.
func (w *waiters) get() *sema {
	// Return if list of waiters is empty.
	if len(*w) == 0 {
		return nil
	}

	// Get the first waiter.
	sema := (*w)[0]
	// Copy everything from the second element to the begining of the list.
	copy((*w)[0:], (*w)[1:])
	(*w)[len(*w)-1] = nil
	*w = (*w)[:len(*w)-1]
	return sema
}

// put appends a new waiter.
func (w *waiters) put(sema *sema) {
	*w = append(*w, sema)
}

// get returns and removes the first 'number' items.
func (items *items) get(number int) []interface{} {
	returnItems := make([]interface{}, 0, number)
	index := 0

	for i := 0; i < number; i++ {
		if i >= len(*items) {
			break
		}

		returnItems = append(returnItems, (*items)[i])
		(*items)[i] = nil
		index++
	}

	*items = (*items)[index:]
	return returnItems
}

// getUntil accepts a checker function to decide how many items to return and
// remove.
func (items *items) getUntil(checker func(item interface{}) bool) []interface{} {
	length := len(*items)
	if length == 0 {
		return []interface{}{}
	}

	returnItems := make([]interface{}, 0, length)
	index := 0
	for i, item := range *items {
		if !checker(item) {
			break
		}

		returnItems = append(returnItems, item)
		index = i
	}

	*items = (*items)[index:]
	return returnItems
}

// Queue is a thread-safe queue.
type Queue struct {
	waiters  waiters
	items    items
	lock     sync.Mutex
	disposed bool
}

// New returs a Queue.
func New(hint int) *Queue {
	return &Queue{
		items: make([]interface{}, 0, hint),
	}
}

// Put will add specified items to Queue.
func (q *Queue) Put(items ...interface{}) error {
	if len(items) == 0 {
		return nil
	}

	q.lock.Lock()

	if q.disposed {
		q.lock.Unlock()
		return errors.New("Queue has been disposed")
	}

	q.items = append(q.items, items...)
	for {
		sema := q.waiters.get()
		if sema == nil {
			break
		}

		sema.response.Add(1)
		sema.wg.Done()
		sema.response.Wait()
		if len(q.items) == 0 {
			break
		}
	}

	q.lock.Unlock()
	return nil
}

// Get will add an item to the queue. If there ar items in the queue, Get will
// return up to the number passed is as parameter. If no items are in queue,
// the method will pause until items are added.
func (q *Queue) Get(number int) ([]interface{}, error) {
	if number < 1 {
		return []interface{}{}, nil
	}

	q.lock.Lock()

	if q.disposed {
		q.lock.Unlock()
		return nil, errors.New("Queue has been disposed")
	}

	var items []interface{}
	if len(q.items) == 0 {
		sema := newSema()
		q.waiters.put(sema)
		sema.wg.Add(1)
		q.lock.Unlock()

		sema.wg.Wait()
		// We are now inside put's lock.
		if q.disposed {
			return nil, errors.New("Queue has been disposed")
		}

		items = q.items.get(number)
		sema.response.Done()
		return items, nil
	}

	items = q.items.get(number)
	q.lock.Unlock()
	return items, nil
}

// TakeUntil takes a function to determine how many items to take from the
// queue. This method does not wait for elements if queue is empty.
func (q *Queue) TakeUntil(checker func(item interface{}) bool) ([]interface{}, error) {
	if checker == nil {
		return nil, nil
	}

	q.lock.Lock()

	if q.disposed {
		q.lock.Unlock()
		return nil, errors.New("Queue has been disposed")
	}

	result := q.items.getUntil(checker)
	q.lock.Unlock()
	return result, nil
}

// Empty returns true if queue is empty, false otherwise.
func (q *Queue) Empty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.items) == 0
}

// Len returns the number of items in the list.
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.items)
}

// Disposed returns a bool indicating whether the method Dispose has been
// called on this queue.
func (q *Queue) Disposed() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.disposed
}

// Dispose disposes of this queue. Any subsequent calls to to Get or Put will
// return an error.
func (q *Queue) Dispose() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.disposed = true
	for _, waiter := range q.waiters {
		waiter.response.Add(1)
		waiter.wg.Done()
	}

	q.items = nil
	q.waiters = nil
}

// ExecuteInParallel will call in parallel the provided function with each item
// in the queue until the queue is exhausted.
func ExecuteInParallel(q *Queue, fn func(interface{})) {
	if q == nil {
		return
	}

	q.lock.Lock()
	todo, done := int64(len(q.items)), int64(-1)
	if todo == 0 {
		return
	}

	numCPU := 1
	if runtime.NumCPU() > 1 {
		numCPU = runtime.NumCPU() - 1
	}

	var wg sync.WaitGroup
	wg.Add(numCPU)
	items := q.items

	for i := 0; i < numCPU; i++ {
		go func() {
			for {
				index := atomic.AddInt64(&done, 1)
				if index >= todo {
					wg.Done()
					break
				}

				fn(items[index])
				items[index] = 0
			}
		}()
	}

	wg.Wait()
	q.lock.Unlock()
	q.Dispose()
}
