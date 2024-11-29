package lmth

func Map[T any](f func(T) Node, ts []T) Node {
	nodes := make([]Node, len(ts))
	for i, t := range ts {
		nodes[i] = f(t)
	}
	return Node{nodes: nodes}
}
