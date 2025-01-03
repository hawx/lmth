package escape_test

import (
	"bytes"
	"html/template"

	"hawx.me/code/lmth"
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

// func TestAttr(t *testing.T) {
//	s := "O'Reilly: How are <i>you</i>?"

//	t.Run("title", func(t *testing.T) {
//		us := renderLMTH(elements.A(lmth.Attr{"title": s}))
//		them := renderTemplate(`<a title="{{.}}"></a>`, s)

//		assert.Equal(t, them, us)
//	})

//	t.Run("href url", func(t *testing.T) {
//		us := renderLMTH(elements.A(lmth.Attr{"href": escape.URL(s)}))
//		them := renderTemplate(`<a href="{{.}}"></a>`, s)

//		assert.Equal(t, strings.Replace(them, "ZgotmplZ", "ZlmthunsafeurlZ", 1), us)
//	})

//	t.Run("href path", func(t *testing.T) {
//		us := renderLMTH(elements.A(lmth.Attr{"href": "/" + escape.Path(s)}))
//		them := renderTemplate(`<a href="/{{.}}"></a>`, s)

//		assert.Equal(t, them, us)
//	})

//	t.Run("href query", func(t *testing.T) {
//		us := renderLMTH(elements.A(lmth.Attr{"href": "?q=" + escape.Query(s)}))
//		them := renderTemplate(`<a href="?q={{.}}"></a>`, s)

//		assert.Equal(t, them, us)
//	})
// }
