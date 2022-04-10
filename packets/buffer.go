package packets

import (
	"encoding/binary"
	"math"
	"unicode/utf16"
)

type Buffer struct {
	B []byte
}

func (b *Buffer) Len() int {
	return len(b.B)
}

func (b *Buffer) Bytes() []byte {
	return b.B
}

func (b *Buffer) Reset() {
	b.B = b.B[:0]
}

func (b *Buffer) WriteF(value float64) {
	b.B = append(b.B, float64ToByte(value)...)
}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

func (b *Buffer) WriteH(value int16) {
	var h, l = byte(value >> 8), byte(value & 0xff)
	b.B = append(b.B, l, h)
}

func (b *Buffer) WriteQ(value int64) {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], uint64(value))
	b.B = append(b.B, buf[:]...)
}

func (b *Buffer) WriteD(value int32) {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], uint32(value))
	b.B = append(b.B, buf[:]...)
}
func (b *Buffer) WriteDU(value uint32) {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], value)
	b.B = append(b.B, buf[:]...)
}

func (b *Buffer) WriteSlice(value []byte) {
	b.B = append(b.B, value...)
}

func (b *Buffer) WriteSingleByte(value byte) {
	b.B = append(b.B, value)
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

	b.B = append(b.B, buf...)
}
