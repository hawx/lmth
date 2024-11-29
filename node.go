package lmth

import (
	"html/template"
	"io"
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

func Element(el string, attr Attr, children ...Node) Node {
	return Node{el: el, attr: attr, children: children}
}

func Text(s string) Node {
	return Node{text: s}
}

type Node struct {
	nodes []Node
	// or
	text string
	// or
	el       string
	attr     Attr
	children []Node
}

func (n Node) internWriteTo(w *errWriter) {
	if n.text != "" {
		w.Write([]byte(template.HTMLEscapeString(n.text)))
		return
	}

	if len(n.nodes) > 0 {
		for _, node := range n.nodes {
			node.internWriteTo(w)
		}

		return
	}

	w.Write([]byte{'<'})
	w.Write([]byte(n.el))
	n.attr.internWriteTo(w)
	w.Write([]byte{'>'})

	if _, ok := voidElements[n.el]; !ok {
		for _, child := range n.children {
			child.internWriteTo(w)
		}

		w.Write([]byte{'<', '/'})
		w.Write([]byte(n.el))
		w.Write([]byte{'>'})
	}
}

func (n Node) WriteTo(w io.Writer) (int64, error) {
	ew := &errWriter{w: w}
	n.internWriteTo(ew)
	return ew.Done()
}
