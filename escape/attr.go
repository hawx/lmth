package escape

import (
	"html/template"
	"strings"
)

func Attr(s string) string {
	return template.HTMLEscapeString(s)
}

func URL(s string) string {
	if len(s) > 0 && s[0] == '/' {
		return Path(s)
	} else if len(s) > 0 && s[0] == '?' {
		if before, after, ok := strings.Cut(s, "="); ok {
			return before + "=" + Query(after)
		}
	}

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
