package lmth

import (
	"maps"
	"slices"
	"strings"

	"hawx.me/code/lmth/escape"
)

var (
	// See https://html.spec.whatwg.org/multipage/indices.html#attributes-1
	urlAttrs = map[string]struct{}{
		"action":     struct{}{},
		"cite":       struct{}{},
		"data":       struct{}{},
		"formaction": struct{}{},
		"href":       struct{}{},
		"itemid":     struct{}{},
		"poster":     struct{}{},
		"src":        struct{}{},
	}
)

type attrToggle bool

const (
	// AttrSet can be used as an attribute value where the attribute doesn't require
	// a value, e.g.
	//
	//	Attr{"selected": "selected"} => <... selected="selected">
	//	Attr{"selected": AttrSet}    => <... selected>
	AttrSet = attrToggle(true)

	// AttrUnset can be used as an attribute value where the attribute should not be
	// written. The real use-case for this is [AttrToggle].
	AttrUnset = attrToggle(false)
)

// RawAttr will cause the string to be written verbatim as the attribute value.
type RawAttr string

// Attr contains HTML attributes to write. Use [escape.Attr] if you are passing
// in data that came from a dubious source.
type Attr map[string]any

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
		switch v := value.(type) {
		case attrToggle:
			// nothing

		case RawAttr:
			w.Write([]byte{'=', '"'})
			w.Write([]byte(v))
			w.Write([]byte{'"'})

		case string:
			w.Write([]byte{'=', '"'})
			if attrTypeURL(k) {
				w.Write([]byte(escape.URL(v)))
			} else {
				w.Write([]byte(escape.Attr(v)))
			}
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
func AttrToggleSet(on bool) any {
	if on {
		return AttrSet
	} else {
		return AttrUnset
	}
}

// AttrToggle will return value if on is true, otherwise AttrUnset. It is useful
// for conditional attributes.
func AttrToggle(on bool, value string) any {
	if on {
		return value
	} else {
		return AttrUnset
	}
}

func attrTypeURL(name string) bool {
	// copy 'html/template' and strip /^data-/ and /.*:/ to apply these rules to
	// more
	if strings.HasPrefix(name, "data-") {
		name = name[5:]
	} else if before, after, ok := strings.Cut(name, ":"); ok {
		if before == "xmlns" {
			return true
		}
		name = after
	}

	_, ok := urlAttrs[name]

	// copy the heuristic to check for /src|uri|url/ as used by
	// 'html/template', it is sensible
	return ok || strings.Contains(name, "src") || strings.Contains(name, "uri") || strings.Contains(name, "url")
}
