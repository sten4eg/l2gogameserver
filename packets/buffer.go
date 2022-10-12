package packets

import (
	"encoding/binary"
	"math"
	"unicode/utf16"
)

type Buffer struct {
	b []byte
}

func (b *Buffer) Len() int {
	return len(b.b)
}

func (b *Buffer) Bytes() []byte {
	cl := make([]byte, len(b.b))
	_ = copy(cl, b.b)
	Put(b)
	return cl
}

func (b *Buffer) Reset() {
	b.b = b.b[:0]
}

func (b *Buffer) WriteF(value float64) {
	b.b = append(b.b, float64ToByte(value)...)
}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

func (b *Buffer) WriteH(value int16) {
	var h, l = byte(value >> 8), byte(value & 0xff)
	b.b = append(b.b, l, h)
}

func (b *Buffer) WriteQ(value int64) {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], uint64(value))
	b.b = append(b.b, buf[:]...)
}
func (b *Buffer) WriteHU(value uint16) {
	b.b = append(b.b, byte(value&0xff), byte(value>>8))
}
func (b *Buffer) WriteD(value int32) {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], uint32(value))
	b.b = append(b.b, buf[:]...)
}
func (b *Buffer) WriteDU(value uint32) {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], value)
	b.b = append(b.b, buf[:]...)
}

func (b *Buffer) WriteSlice(value []byte) {
	b.b = append(b.b, value...)
}

func (b *Buffer) WriteSingleByte(value byte) {
	b.b = append(b.b, value)
}

const EmptyByte byte = 0

func (b *Buffer) WriteS(value string) {
	utf16Slice := utf16.Encode([]rune(value))

	var buf []byte
	for _, v := range utf16Slice {
		if v < math.MaxInt8 {
			buf = append(buf, byte(v), 0)
		} else {
			f, s := uint8(v&0xff), uint8(v>>8)
			buf = append(buf, f, s)
		}
	}

	buf = append(buf, EmptyByte, EmptyByte)

	b.b = append(b.b, buf...)
}
