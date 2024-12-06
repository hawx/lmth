package escape

import (
	"bytes"
	"html/template"
	"strings"
	"testing"

	"hawx.me/code/assert"
	"hawx.me/code/lmth"
	"hawx.me/code/lmth/elements"
)

func renderLMTH(node lmth.Node) string {
	var buf bytes.Buffer
	node.WriteTo(&buf)
	return buf.String()
}

func renderTemplate(tmpl string, data any) string {
	var buf bytes.Buffer
	template.Must(template.New("").Parse(tmpl)).Execute(&buf, data)
	return buf.String()
}

func TestAttr(t *testing.T) {
	s := "O'Reilly: How are <i>you</i>?"

	t.Run("title", func(t *testing.T) {
		us := renderLMTH(elements.A(lmth.Attr{"title": Attr(s)}))
		them := renderTemplate(`<a title="{{.}}"></a>`, s)

		assert.Equal(t, them, us)
	})

	t.Run("href url", func(t *testing.T) {
		us := renderLMTH(elements.A(lmth.Attr{"href": URL(s)}))
		them := renderTemplate(`<a href="{{.}}"></a>`, s)

		assert.Equal(t, strings.Replace(them, "ZgotmplZ", "ZlmthunsafeurlZ", 1), us)
	})

	t.Run("href path", func(t *testing.T) {
		us := renderLMTH(elements.A(lmth.Attr{"href": "/" + Path(s)}))
		them := renderTemplate(`<a href="/{{.}}"></a>`, s)

		assert.Equal(t, them, us)
	})

	t.Run("href query", func(t *testing.T) {
		us := renderLMTH(elements.A(lmth.Attr{"href": "?q=" + Query(s)}))
		them := renderTemplate(`<a href="?q={{.}}"></a>`, s)

		assert.Equal(t, them, us)
	})
}
