package lmth

import (
	"maps"
	"slices"
)

// AttrSet can be used as an attribute value where the attribute doesn't require
// a value, e.g.
//
//	Attr{"selected": "selected"} => <... selected="selected">
//	Attr{"selected": AttrSet}    => <... selected>
const AttrSet = "__hawx.me/code/lmth:AttrSet__"

// AttrUnset can be used as an attribute value where the attribute should not be
// written. The real use-case for this is [AttrToggle].
const AttrUnset = "__hawx.me/code/lmth:AttrUnset__"

// Attr contains HTML attributes to write. Use [escape.Attr] if you are passing
// in data that came from a dubious source.
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

// Merge adds another set of attributes to this one, returning this again.
func (a Attr) Merge(other Attr) Attr {
	for k, v := range other {
		a[k] = v
	}
	return a
}

// AttrToggleSet will return AttrSet if on is true, otherwise AttrUnset. It is
// useful for toggling things like "selected".
func AttrToggleSet(on bool) string {
	if on {
		return AttrSet
	} else {
		return AttrUnset
	}
}

// AttrToggle will return value if on is true, otherwise AttrUnset. It is useful
// for conditional attributes.
func AttrToggle(on bool, value string) string {
	if on {
		return value
	} else {
		return AttrUnset
	}
}
