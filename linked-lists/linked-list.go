package main

// DoublyLinkedList is a doubly linked list.
type DoublyLinkedList struct {
	head  *node
	tail  *node
	count int
}

type node struct {
	value int
	next  *node
	prev  *node
}

// GetLinkedList initializes an empty doubly linked list.
func GetLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil, 0}
}

// GetLinkedListFromValues serves as constructor of a doubly linked list.
func GetLinkedListFromValues(vals []int) *DoublyLinkedList {
	ll := GetLinkedList()
	if len(vals) == 0 {
		return ll
	}

	for _, val := range vals {
		ll.Insert(val)
	}
	return ll
}

// Insert creates a new node given an integer value and adds it to the end of
// the linked list.
func (ll *DoublyLinkedList) Insert(val int) {
	newNode := &node{value: val}
	ll.insertNode(newNode)
}

func (ll *DoublyLinkedList) insertNode(newNode *node) {
	ll.count++

	if ll.tail == nil {
		ll.head, ll.tail = newNode, newNode
	} else {
		ll.tail.next = newNode
		newNode.prev = ll.tail
		ll.tail = newNode
	}
}

func (ll *DoublyLinkedList) getNode(index int) *node {
	node := ll.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

// Slice returns a slice of all integer values in the linked list.
func (ll *DoublyLinkedList) Slice() []int {
	slice := make([]int, ll.count)

	node := ll.head
	for i := 0; i < ll.count; i++ {
		slice[i] = node.value
		node = node.next
	}

	return slice
}

// Get returns thevalue at a given position within the list.
func (ll *DoublyLinkedList) Get(index int) int {
	return ll.getNode(index).value
}

// Len returns the total number of nodes in the list.
func (ll *DoublyLinkedList) Len() int {
	return ll.count
}
