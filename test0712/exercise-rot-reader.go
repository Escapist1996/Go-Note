package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
func rot13(b byte) byte {
	switch {
		case 'A' <= b && b <= 'M':
		b = b + 13
		case 'N' <= b && b <= 'Z':
		b = b - 13
		case 'a' <= b && b <= 'm':
		b = b + 13
		case 'n' <= b && b <= 'z':
		b = b - 13
	}
	return b
}

func (mr rot13Reader) Read(b []byte) (int, error) {
	n, e := mr.r.Read(b)
	for i := 0; i < 0; i++ {
		b[i] = rot13(b[i])
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
