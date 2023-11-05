package badgo

import (
	"bufio"
	"bytes"
	"encoding/binary"
	fmt "fmt"
	"io"
	"testing"
	"unsafe"
)

func readString(b *Buffer) (string, error) {
	l, n := binary.Varint(b.Peek())
	for n == 0 {
		// Not enough data to read the varint. Can we get more?
		if err := b.Refill(); err != nil {
			return "", err
		}
		l, n = binary.Varint(b.Peek())
	}
	if n < 0 {
		return "", fmt.Errorf("blah")
	}
	b.i += n

	if l < 0 {
		return "", fmt.Errorf("negative length")
	}

	s, err := b.Next(int(l))
	return string(s), err
}

func readStringB(r *bufio.Reader) (string, error) {
	l, err := binary.ReadVarint(r)
	if err != nil {
		return "", err
	}
	b := make([]byte, l)
	_, err = io.ReadFull(r, b)
	return *(*string)(unsafe.Pointer(&b)), err
}

func readStringC(data []byte) (string, int, error) {
	l, n := binary.Varint(data)
	if n == 0 {
		return "", 0, io.ErrUnexpectedEOF
	}
	if n < 0 {
		return "", 0, fmt.Errorf("invalid length")
	}
	if n+int(l) > len(data) {
		return "", 0, io.ErrUnexpectedEOF
	}

	// Casting []byte to string causes an allocation, but we want that here
	return string(data[n : n+int(l)]), n + int(l), nil
}

func BenchmarkReadStringC(b *testing.B) {
	data := bytes.Repeat([]byte{16, 'c', 'h', 'e', 'e', 's', 'e', 'i', 't'}, 1000)
	b.ReportAllocs()
	for i := 0; i < b.N; i += 1000 {
		offset := 0
		for j := 0; j < 1000; j++ {
			_, n, err := readStringC(data[offset:])
			if err != nil {
				b.Fatal(err)
			}
			offset += n
		}
	}
}

func BenchmarkReadStringB(b *testing.B) {
	data := bytes.Repeat([]byte{16, 'c', 'h', 'e', 'e', 's', 'e', 'i', 't'}, 1000)
	b.ReportAllocs()
	r := bytes.NewReader(nil)
	buf := bufio.NewReader(nil)
	for i := 0; i < b.N; i += 1000 {
		r.Reset(data)
		buf.Reset(r)
		for j := 0; j < 1000; j++ {
			if _, err := readStringB(buf); err != nil {
				b.Fatal(err)
			}
		}
	}
}

func BenchmarkReadString(b *testing.B) {
	data := bytes.Repeat([]byte{16, 'c', 'h', 'e', 'e', 's', 'e', 'i', 't'}, 1000)
	b.ReportAllocs()
	r := bytes.NewReader(nil)
	buf := NewBuffer()
	for i := 0; i < b.N; i += 1000 {
		r.Reset(data)
		buf.Reset(r)
		for j := 0; j < 1000; j++ {
			if _, err := readString(buf); err != nil {
				b.Fatal(err)
			}
		}
	}
}

func NewBuffer() *Buffer {
	return &Buffer{
		data: make([]byte, 1000),
	}
}

type Buffer struct {
	data []byte
	i    int
	r    io.Reader
	err  error
}

func (b *Buffer) Reset(r io.Reader) {
	b.data = b.data[:0]
	b.i = 0
	b.err = nil
	b.r = r
}

func (b *Buffer) Next(l int) ([]byte, error) {
	if b.i+l > len(b.data) {
		// Asking for more data than we have. refill
		if err := b.refill(l); err != nil {
			return nil, err
		}
	}

	b.i += l
	return b.data[b.i-l : b.i], nil
}

// Peek allows direct access to the current remaining buffer
func (b *Buffer) Peek() []byte {
	return b.data[b.i:]
}

// Dicard consumes data in the current buffer
func (b *Buffer) Discard(n int) {
	b.i += n
}

// Refill forces the buffer to try to put at least one more byte into its buffer
func (b *Buffer) Refill() error {
	return b.refill(1)
}

func (b *Buffer) refill(l int) error {
	if b.err != nil {
		// We already know we can't get more data
		return b.err
	}

	// fill the rest of the buffer from the reader
	if b.r != nil {
		// shift existing data down over the read portion of the buffer
		n := copy(b.data[:cap(b.data)], b.data[b.i:])
		b.i = 0

		read, err := io.ReadFull(b.r, b.data[n:cap(b.data)])

		b.data = b.data[:n+read]
		if err == io.ErrUnexpectedEOF {
			err = io.EOF
		}
		b.err = err
	}

	if b.i+l > len(b.data) {
		// Still not enough data
		return io.ErrUnexpectedEOF
	}

	return nil
}
