package packets

import (
	"encoding/binary"
	"math"
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

	buf := make([]byte, 0, len(value)*2+2)
	if len(value) != 0 {
		for _, v := range []byte(value) {
			buf = append(buf, v, EmptyByte)
		}
	}

	buf = append(buf, EmptyByte, EmptyByte)

	b.B = append(b.B, buf...)

}
