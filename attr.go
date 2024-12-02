package lmth

import (
	"maps"
	"slices"
)

const AttrSet = "__hawx.me/code/lmth:AttrSet__"
const AttrUnset = "__hawx.me/code/lmth:AttrUnset__"

type Attr map[string]string

func (a Attr) internWriteTo(w *errWriter) {
	keys := slices.Collect(maps.Keys(a))
	slices.Sort(keys)

	for _, k := range keys {
		value := a[k]
		if value == AttrUnset {
			continue
		}

		w.Write([]byte{' '})
		w.Write([]byte(k))
		if value != AttrSet {
			w.Write([]byte{'=', '"'})
			w.Write([]byte(value))
			w.Write([]byte{'"'})
		}
	}
}

func (a Attr) Merge(other Attr) Attr {
	for k, v := range other {
		a[k] = v
	}
	return a
}

func AttrToggleSet(on bool) string {
	if on {
		return AttrSet
	} else {
		return AttrUnset
	}
}

func AttrToggle(on bool, value string) string {
	if on {
		return value
	} else {
		return AttrUnset
	}
}
