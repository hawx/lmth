package lmth

import (
	"bytes"
	"fmt"
	"testing"

	"hawx.me/code/assert"
)

func TestMap(t *testing.T) {
	res := Map(func(s string) Node {
		return Text(s)
	}, []string{"a", "b", "c"})

	var buf bytes.Buffer
	res.WriteTo(&buf)

	assert.Equal(t, "abc", buf.String())
}

func TestMap2(t *testing.T) {
	args := []string{"a", "b", "c"}
	res := Map2(func(i int, s string) Node {
		end := ""
		if i < len(args)-1 {
			end = ", "
		}
		return Text(fmt.Sprintf("(%d) %s%s", i+1, s, end))
	}, args)

	var buf bytes.Buffer
	res.WriteTo(&buf)

	assert.Equal(t, "(1) a, (2) b, (3) c", buf.String())
}

func TestToggle(t *testing.T) {
	var bufOn bytes.Buffer
	Toggle(true, Text("hi")).WriteTo(&bufOn)

	var bufOff bytes.Buffer
	Toggle(false, Text("hi")).WriteTo(&bufOff)

	assert.Equal(t, "hi", bufOn.String())
	assert.Equal(t, "", bufOff.String())
}

func TestJoin(t *testing.T) {
	var buf bytes.Buffer
	Join(Text("a"), Text("b"), Text("c")).WriteTo(&buf)

	assert.Equal(t, "abc", buf.String())
}
