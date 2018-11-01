package main

// RemoveDuplicates removes duplicate entries using a map as a memory buffer to
// keep track of repeated elements.
func (ll *DoublyLinkedList) RemoveDuplicates() {
	seen := make(map[int]struct{})

	node := ll.head
	// Save the count ahead of time as elements may be erased.
	count := ll.count
	for i := 0; i < count; i++ {
		// Check if value has already been seen.
		if _, ok := seen[node.value]; ok {
			ll.removeNode(node)
		} else {
			seen[node.value] = struct{}{}
		}

		// Move onto the next node.
		if node.next == nil {
			break
		}
		node = node.next
	}
}

func (ll *DoublyLinkedList) RemoveDuplicatesBrute() {
	current := ll.head

	if current == nil {
		return
	}

	for current != nil {

		runner := current.next
		for runner != nil {

			if runner.value == current.value {
				next := runner.next
				ll.removeNode(runner)
				runner = next
			} else {
				runner = runner.next
			}
		}

		current = current.next
	}
}
