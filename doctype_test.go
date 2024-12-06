package lmth_test

import (
	"bytes"
	"testing"

	"hawx.me/code/assert"
	"hawx.me/code/lmth"
	"hawx.me/code/lmth/elements"
)

func TestDoctype(t *testing.T) {
	assert := assert.Wrap(t)

	doc := lmth.Doctype(elements.Html(lmth.Attr{}))

	var buf bytes.Buffer
	doc.WriteTo(&buf)

	expected := `<!DOCTYPE html><html></html>`
	assert(buf.String()).Equal(expected)
}
