package bitstream

import "bytes"

type Buffer struct {
	b   byte
	l   int
	buf bytes.Buffer
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
