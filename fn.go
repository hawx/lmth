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

// Toggle will hide the node from being written if on is false.
func Toggle(on bool, node Node) Node {
	if !on {
		node.nodeType = nodeTypeEmpty
	}

	return node
}
