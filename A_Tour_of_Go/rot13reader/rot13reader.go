package main

import (
	"io"
	"os"
	"strings"
)

type rot13reader struct {
	r io.Reader
}

func (rot13 rot13reader) Read(b []byte) (int, error) {
	result, err := rot13.r.Read(b)
	for i, char := range b {
		if 'a' <= char && char <= 'm' || 'A' <= char && char <= 'M' {
			b[i] += 13
		} else if 'n' <= char && char <= 'z' || 'N' <= char && char <= 'Z' {
			b[i] -= 13
		}
	}
	return result, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13reader{s}
	io.Copy(os.Stdout, &r)
}
