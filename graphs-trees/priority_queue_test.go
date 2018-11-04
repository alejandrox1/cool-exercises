package main

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert new item and modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	actual := map[string]int{}
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		actual[item.value] = item.priority
	}

	expected := map[string]int{
		"banana": 3, "apple": 2, "pear": 4, "orange": 5,
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v. Got: %v\n", expected, actual)
	}
}
