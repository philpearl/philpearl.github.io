package badgo

import (
	"encoding/binary"
	"io"
	"strings"
)

type Reader interface {
	io.Reader
	io.ByteReader
}

func readStringA(r Reader) (string, error) {
	l, err := binary.ReadVarint(r)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	b.Grow(int(l))
	_, err = io.CopyN(&b, r, l)
	return b.String(), err
}
