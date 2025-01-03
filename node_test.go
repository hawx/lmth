package lmth_test

import (
	"bytes"
	"html/template"
	"net/url"
	"strings"
	"testing"

	"hawx.me/code/assert"
	"hawx.me/code/lmth"
	. "hawx.me/code/lmth/elements"
	"hawx.me/code/lmth/escape"
)

func TestReadme(t *testing.T) {
	assert := assert.Wrap(t)

	fruits := []string{"apple", "pear"}
	disabled := true

	doc := Html(lmth.Attr{},
		Head(lmth.Attr{},
			Title(lmth.Attr{}, lmth.Text("A test page")),
			Style(lmth.Attr{}, lmth.RawText(`body > h1 { font-weight: 100px; }`)),
		),
		Body(lmth.Attr{},
			H1(lmth.Attr{"class": "heading-xl"},
				lmth.Text("My web page"),
			),
			P(lmth.Attr{},
				lmth.Text("This is some fruit"),
			),
			A(lmth.Attr{"href": "bad://url"}),
			A(lmth.Attr{"href": lmth.RawAttr(escape.Path("/url ok this"))}),
			A(lmth.Attr{"href": lmth.RawAttr("?url=" + url.QueryEscape(" ok this"))}),
			A(lmth.Attr{"href": lmth.RawAttr("bad://url")}),
			Ul(lmth.Attr{},
				lmth.Map(func(s string) lmth.Node {
					return Li(lmth.Attr{}, lmth.Text(s), Img(lmth.Attr{"src": escape.Attr(s) + ".jpg"}))
				}, fruits),
			),
			lmth.Toggle(!disabled, P(lmth.Attr{"class": "help"},
				lmth.Text("You can save your progress"),
			)),
			Button(lmth.Attr{"disabled": lmth.AttrToggleSet(disabled)}, lmth.Text("Save")),
		),
	)

	var buf bytes.Buffer
	doc.WriteTo(&buf)

	expected := `<html><head><title>A test page</title><style>body > h1 { font-weight: 100px; }</style></head><body><h1 class="heading-xl">My web page</h1><p>This is some fruit</p><a href="#ZlmthunsafeurlZ"></a><a href="/url%20ok%20this"></a><a href="?url=+ok+this"></a><a href="bad://url"></a><ul><li>apple<img src="apple.jpg"></li><li>pear<img src="pear.jpg"></li></ul><button disabled>Save</button></body></html>`
	assert(buf.String()).Equal(expected)
}

func TestEscaping(t *testing.T) {
	s := "O'Reilly: How are <i>you</i>?"

	t.Run("title", func(t *testing.T) {
		us := renderLMTH(A(lmth.Attr{"title": s}))
		them := renderTemplate(`<a title="{{.}}"></a>`, s)

		assert.Equal(t, them, us)
	})

	t.Run("href url", func(t *testing.T) {
		us := renderLMTH(A(lmth.Attr{"href": s}))
		them := renderTemplate(`<a href="{{.}}"></a>`, s)

		assert.Equal(t, strings.Replace(them, "ZgotmplZ", "ZlmthunsafeurlZ", 1), us)
	})

	t.Run("href path", func(t *testing.T) {
		us := renderLMTH(A(lmth.Attr{"href": "/" + s}))
		them := renderTemplate(`<a href="/{{.}}"></a>`, s)

		assert.Equal(t, them, us)
	})

	t.Run("href query", func(t *testing.T) {
		us := renderLMTH(A(lmth.Attr{"href": "?q=" + s}))
		them := renderTemplate(`<a href="?q={{.}}"></a>`, s)

		assert.Equal(t, them, us)
	})
}

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
