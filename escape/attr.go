package escape

import (
	"html/template"
)

func Attr(s string) string {
	return template.HTMLEscapeString(s)
}

func URL(s string) string {
	if !isSafeURL(s) {
		return "#ZlmthunsafeurlZ"
	}

	return s
}

func Path(s string) string {
	return urlProcessor(true, s)
}

func Query(s string) string {
	return urlProcessor(false, s)
}
