package lmth

import (
	"io"
	"text/template"
)

type nodeType uint8

const (
	nodeTypeEmpty nodeType = iota
	nodeTypeText
	nodeTypeNode
	nodeTypeMulti
)

var voidElements = map[string]struct{}{
	"area":   struct{}{},
	"base":   struct{}{},
	"br":     struct{}{},
	"col":    struct{}{},
	"embed":  struct{}{},
	"hr":     struct{}{},
	"img":    struct{}{},
	"input":  struct{}{},
	"link":   struct{}{},
	"meta":   struct{}{},
	"param":  struct{}{},
	"source": struct{}{},
	"track":  struct{}{},
	"wbr":    struct{}{},
}

// Element creates a [Node] representing a HTML element.
func Element(el string, attr Attr, children ...Node) Node {
	return Node{nodeType: nodeTypeNode, data: el, attr: attr, nodes: children}
}

// Text creates a [Node] representing text, HTML escaping is applied.
func Text(s string) Node {
	return Node{nodeType: nodeTypeText, data: template.HTMLEscapeString(s)}
}

// RawText creates a [Node] representing text, with no escaping. Be careful.
func RawText(s string) Node {
	return Node{nodeType: nodeTypeText, data: s}
}

// A Node represents some HTML structure.
type Node struct {
	nodeType nodeType
	data     string
	attr     Attr
	nodes    []Node
}

func (n Node) internWriteTo(w *errWriter) {
	switch n.nodeType {
	case nodeTypeEmpty:
		//

	case nodeTypeText:
		w.Write([]byte(n.data))

	case nodeTypeNode:
		w.Write([]byte{'<'})
		w.Write([]byte(n.data))
		n.attr.internWriteTo(w)
		w.Write([]byte{'>'})

		if _, ok := voidElements[n.data]; !ok {
			for _, child := range n.nodes {
				child.internWriteTo(w)
			}

			w.Write([]byte{'<', '/'})
			w.Write([]byte(n.data))
			w.Write([]byte{'>'})
		}

	case nodeTypeMulti:
		for _, node := range n.nodes {
			node.internWriteTo(w)
		}
	}
}

// WriteTo writes the [Node] to w.
func (n Node) WriteTo(w io.Writer) (int64, error) {
	ew := &errWriter{w: w}
	n.internWriteTo(ew)
	return ew.Done()
}
