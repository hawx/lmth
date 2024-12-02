package lmth

func Map[T any](f func(T) Node, ts []T) Node {
	nodes := make([]Node, len(ts))
	for i, t := range ts {
		nodes[i] = f(t)
	}
	return Node{nodeType: nodeTypeMulti, nodes: nodes}
}

func Toggle(on bool, node Node) Node {
	if !on {
		node.nodeType = nodeTypeEmpty
	}

	return node
}
