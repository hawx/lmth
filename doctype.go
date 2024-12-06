package lmth

import "io"

type doctypeWriter struct{ w io.WriterTo }

func (d *doctypeWriter) WriteTo(w io.Writer) (int64, error) {
	var total int64
	c, err := w.Write([]byte(`<!DOCTYPE html>`))
	total += int64(c)
	if err != nil {
		return total, err
	}

	ct, err := d.w.WriteTo(w)
	total += ct
	return total, err
}

func Doctype(node Node) io.WriterTo {
	return &doctypeWriter{w: node}
}
