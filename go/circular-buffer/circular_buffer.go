package circular

import "fmt"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
	data []byte
	size int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{data: make([]byte, 0, size), size: size}
}

func (b *Buffer) ReadByte() (byte, error) {
	if size := len(b.data); size < 1 {
		return 0, fmt.Errorf("buffer with size: %d, is empty", b.size)
	} else {
		value := b.data[size-1]
		b.data = b.data[0 : size-1]
		return value, nil
	}
}

func (b *Buffer) WriteByte(c byte) error {
	if added := len(b.data); added >= b.size {
		return fmt.Errorf("buffer with size: %d, is full", b.size)
	}

	temp := []byte{c}
	temp = append(temp, b.data...)
	b.data = temp

	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if added := len(b.data); added >= b.size {
		b.ReadByte()
	}
	b.WriteByte(c)
}

func (b *Buffer) Reset() {
	b.data = make([]byte, 0, b.size)
}
