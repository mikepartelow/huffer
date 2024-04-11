package bitstream

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/bits"
)

// TODO: read buffer should be a different type from write buffer, since we can't read and write from the same one
//	in fact, use io.Reader and io.Writer

type Buffer struct {
	b   byte
	l   int
	buf bytes.Buffer
}

func NewBuffer(s []byte) *Buffer {
	return &Buffer{
		buf: *bytes.NewBuffer(s),
	}
}

// Read reads a bit from the buffer, and returns io.EOF if the buffer is empty
func (b *Buffer) Read() (int, error) {
	if b.l == 0 {
		var err error
		b.b, err = b.buf.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return 0, io.EOF
			}
			return 0, fmt.Errorf("error reading byte: %w", err)
		}
		b.b = bits.Reverse8(b.b)
	}

	bit := int((b.b >> b.l) & 1)

	if b.l++; b.l == 8 {
		b.l = 0
	}

	return bit, nil

}

// Write writes a bit to the Buffer. Only the least significant bit of "bit" is used.
func (b *Buffer) Write(bit int) error {
	b.b = (b.b << 1) | byte(bit&0b1)
	b.l++

	if b.l == 8 {
		return b.flush()
	}

	return nil
}

// Bytes returns a slice of right-padded bytes representing the bitstream
func (b *Buffer) Bytes() []byte {
	err := b.flush()
	if err != nil {
		panic(err)
	}
	r := b.buf.Bytes()
	b.buf.Reset()
	return r
}

func (b *Buffer) flush() error {
	if b.l == 0 {
		return nil
	}

	defer func() { b.b, b.l = 0, 0 }()
	return b.buf.WriteByte(b.b << (8 - b.l))
}
