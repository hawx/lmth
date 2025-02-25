package lmth

// Map returns a [Node] that contains the result of calling f for every element
// of ts.
func Map[T any](f func(T) Node, ts []T) Node {
	nodes := make([]Node, len(ts))
	for i, t := range ts {
		nodes[i] = f(t)
	}
	return Node{nodeType: nodeTypeMulti, nodes: nodes}
}

// Map2 returns a [Node] that contains the result of calling f for every element
// and index of ts.
func Map2[T any](f func(int, T) Node, ts []T) Node {
	nodes := make([]Node, len(ts))
	for i, t := range ts {
		nodes[i] = f(i, t)
	}
	return Node{nodeType: nodeTypeMulti, nodes: nodes}
}

// Toggle will hide the node from being written if on is false.
func Toggle(on bool, node Node) Node {
	if !on {
		node.nodeType = nodeTypeEmpty
	}

	return node
}

// Join returns a [Node] that writes nodes in sequence.
func Join(nodes ...Node) Node {
	return Node{nodeType: nodeTypeMulti, nodes: nodes}
}
