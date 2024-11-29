package lmth_test

import (
	"bytes"
	"testing"

	"hawx.me/code/assert"
	"hawx.me/code/lmth"
	. "hawx.me/code/lmth/elements"
)

func TestReadme(t *testing.T) {
	assert := assert.Wrap(t)

	fruits := []string{"apple", "pear"}

	doc := Body(lmth.Attr{},
		H1(lmth.Attr{"class": "heading-xl"},
			lmth.Text("My web page"),
		),
		P(lmth.Attr{},
			lmth.Text("This is some fruit"),
		),
		Ul(lmth.Attr{},
			lmth.Map(func(s string) lmth.Node {
				return Li(lmth.Attr{}, lmth.Text(s), Img(lmth.Attr{"src": s + ".jpg"}))
			}, fruits),
		),
	)

	var buf bytes.Buffer
	doc.WriteTo(&buf)

	expected := `<body><h1 class="heading-xl">My web page</h1><p>This is some fruit</p><ul><li>apple<img src="apple.jpg"></li><li>pear<img src="pear.jpg"></li></ul></body>`
	assert(buf.String()).Equal(expected)
}
