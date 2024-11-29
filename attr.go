package lmth

import (
	"maps"
	"slices"
)

type Attr map[string]string

func (a Attr) internWriteTo(w *errWriter) {
	keys := slices.Collect(maps.Keys(a))
	slices.Sort(keys)

	for _, k := range keys {
		w.Write([]byte{' '})
		w.Write([]byte(k))
		w.Write([]byte{'=', '"'})
		w.Write([]byte(a[k]))
		w.Write([]byte{'"'})
	}
}

func (a Attr) Merge(other Attr) Attr {
	for k, v := range other {
		a[k] = v
	}
	return a
}
