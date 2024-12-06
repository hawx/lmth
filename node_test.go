package lmth_test

import (
	"bytes"
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

	expected := `<html><head><title>A test page</title><style>body > h1 { font-weight: 100px; }</style></head><body><h1 class="heading-xl">My web page</h1><p>This is some fruit</p><ul><li>apple<img src="apple.jpg"></li><li>pear<img src="pear.jpg"></li></ul><button disabled>Save</button></body></html>`
	assert(buf.String()).Equal(expected)
}
