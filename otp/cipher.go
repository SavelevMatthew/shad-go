//go:build !solution

package otp

import (
	"io"
)

type XORReader struct {
	r    io.Reader
	prng io.Reader
}

func (x XORReader) Read(p []byte) (n int, err error) {
	n, err = x.r.Read(p)
	if n == 0 {
		return
	}
	code := make([]byte, n)
	_, err = x.prng.Read(code)
	if err != nil {
		return
	}

	for i, v := range code {
		p[i] = p[i] ^ v
	}

	return
}

type XORWriter struct {
	w    io.Writer
	prng io.Reader
}

func (x XORWriter) Write(p []byte) (n int, err error) {
	code := make([]byte, len(p))
	n, err = x.prng.Read(code)
	if err != nil {
		return
	}

	result := make([]byte, len(p))
	for i := range result {
		result[i] = p[i] ^ code[i]
	}

	return x.w.Write(result)
}

func NewReader(r io.Reader, prng io.Reader) io.Reader {
	return XORReader{r, prng}
}

func NewWriter(w io.Writer, prng io.Reader) io.Writer {
	return XORWriter{w, prng}
}
