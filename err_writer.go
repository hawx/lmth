package lmth

import "io"

type errWriter struct {
	w     io.Writer
	err   error
	total int64
}

func (w *errWriter) Write(data []byte) {
	if w.err == nil {
		c, err := w.w.Write(data)
		w.total += int64(c)
		w.err = err
	}
}

func (w *errWriter) Done() (int64, error) {
	return w.total, w.err
}
