package main

import (
    "fmt"
    "sync"
)

// Node
type Node struct {
    value string
}

// Node String returns the node's value.
func (n *Node) String() string {
    return fmt.Sprintf("%v", n.value)
}



// NodeQueue is a queue of Nodes.
type NodeQueue struct {
    items []Node
    lock sync.RWMutex
}

// NewNodeQueue creates a NodeQueue.
func NewNodeQueue() *NodeQueue {
    return &NodeQueue{}
}

// Enqueue adds a Node to the end of the queue.
func (q *NodeQueue) Enqueue(n Node) {
    q.lock.Lock()
    q.items = append(q.items, n)
    q.lock.Unlock()
}

// Dequeue removes a Node from the front of the queue.
func (q *NodeQueue) Dequeue() *Node {
    q.lock.Lock()
    item := q.items[0]
    q.items = q.items[1:len(q.items)]
    q.lock.Unlock()
    return &item
}

// Front returns a Node from the front of the queue without removing it.
func (q *NodeQueue) Front() *Node {
    q.lock.RLock()
    item := q.items[0]
    q.lock.RUnlock()
    return &item
}

// IsEmpty checks is the NodeQueue is empty.
func (q *NodeQueue) IsEmpty() bool {
    q.lock.RLock()
    defer q.lock.RUnlock()
    return len(q.items) == 0
}

// Size returns the size of the NodeQueue.
func (q *NodeQueue) Size() int {
    q.lock.RLock()
    defer q.lock.RUnlock()
    return len(q.items)
}



// StringGraph
type StringGraph struct {
    nodes []*Node
    edges map[Node][]*Node
    lock sync.RWMutex
}

// AddNode adds a node to the graph.
func (g *StringGraph) AddNode(n *Node) {
    g.lock.Lock()
    g.nodes = append(g.nodes, n)
    g.lock.Unlock()
}

// AddEdge adds an edge to the graph.
func (g *StringGraph) AddEdge(n1, n2 *Node) {
    g.lock.Lock()

    if g.edges == nil {
        g.edges = make(map[Node][]*Node)
    }

    g.edges[*n1] = append(g.edges[*n1], n2)
    g.edges[*n2] = append(g.edges[*n2], n1)

    g.lock.Unlock()
}

// StringGraph String
func (g *StringGraph) String() string {
    g.lock.RLock()

    s := ""
    for i := 0; i < len(g.nodes); i++ {
        s += g.nodes[i].String() + " -> "

        near := g.edges[*g.nodes[i]]
        for j := 0; j < len(near); j++ {
            s += near[j].String() + " "
        }
        s += "\n"
    }

    g.lock.RUnlock()

    return fmt.Sprintln(s)
}

// Traverse implements a breath-first search.
func (g *StringGraph) Traverse(f func(*Node)) {
    g.lock.RLock()

    q := NewNodeQueue()

    n := g.nodes[0]
    q.Enqueue(*n)

    visited := make(map[*Node]bool)
    for {
        if q.IsEmpty() {
            break
        }

        node := q.Dequeue()
        visited[node] = true
        near := g.edges[*node]

        for i := 0; i < len(near); i++ {
            j := near[i]
            if !visited[j] {
                q.Enqueue(*j)
                visited[j] = true
            }
        }

        if f != nil {
            f(node)
        }
    }

    g.lock.RUnlock()
}
