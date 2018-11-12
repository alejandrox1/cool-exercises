package main

import (
    "testing"
)

var g StringGraph

func fillGraph() {
    nA := Node{"A"}
    nB := Node{"B"}
    nC := Node{"C"}
    nD := Node{"D"}
    nE := Node{"E"}
    nF := Node{"F"}

    g.AddNode(&nA)
    g.AddNode(&nB)
    g.AddNode(&nC)
    g.AddNode(&nD)
    g.AddNode(&nE)
    g.AddNode(&nF)

    g.AddEdge(&nA, &nB)
    g.AddEdge(&nA, &nC)
    g.AddEdge(&nB, &nE)
    g.AddEdge(&nC, &nE)
    g.AddEdge(&nE, &nF)
    g.AddEdge(&nD, &nA)
}

func TestStringGraph(t *testing.T) {
    fillGraph()

    t.Logf("%s", g.String())
}

func TestTraverse(t *testing.T) {
    g.Traverse(func(n *Node) {
        t.Logf("%v\n", n)
    })
}
