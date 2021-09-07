package packets

import (
	"bytes"
	"encoding/binary"
)

type Buffer struct {
	bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{}
}

func (b *Buffer) WriteF(value float64) {
	err := binary.Write(b, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}
}

func (b *Buffer) WriteH(value int16) {
	err := binary.Write(b, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}
}

func (b *Buffer) WriteQ(value int64) {
	err := binary.Write(b, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}
}
func (b *Buffer) WriteD(value int32) {
	err := binary.Write(b, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}
}
func (b *Buffer) WriteDU(value uint32) {
	err := binary.Write(b, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}
}
func (b *Buffer) WriteSlice(val []byte) {
	err := binary.Write(b, binary.LittleEndian, val)
	if err != nil {
		panic(err)
	}
}

func (b *Buffer) WriteSingleByte(value byte) {
	err := b.WriteByte(value)
	if err != nil {
		panic(err)
	}
}

const EmptyByte byte = 0

func (b *Buffer) WriteS(value string) {

	if len(value) != 0 {
		for _, v := range []byte(value) {
			err := binary.Write(b, binary.LittleEndian, v)
			if err != nil {
				panic(err)
			}
			_ = binary.Write(b, binary.LittleEndian, EmptyByte)

		}
	}

	_ = binary.Write(b, binary.LittleEndian, []byte{EmptyByte, EmptyByte})
}

type Reader struct {
	r *bytes.Reader
	b *Buffer
}

func NewReader(buffer []byte) *Reader {
	return &Reader{
		r: bytes.NewReader(buffer),
		b: NewBuffer(),
	}
}

func (r *Reader) CurrentIndex() int64 {
	l := r.r.Len()
	s := r.r.Size()
	currIndex := s - int64(l)
	return currIndex
}

func (r *Reader) UnreadBytes(b int) {
	for i := 0; i < b; i++ {
		err := r.r.UnreadByte()
		if err != nil {
			panic(err)
		}
	}
}
func (r *Reader) ReadBytes(number int) []byte {
	buffer := make([]byte, number)
	n, _ := r.r.Read(buffer)
	if n < number {
		return []byte{}
	}

	return buffer
}

func (r *Reader) ReadSingleByte() byte {
	buffer := make([]byte, 1)
	n, _ := r.r.Read(buffer)
	if n < 1 {
		return 0
	}

	return buffer[0]
}
func (r *Reader) ReadUInt64() uint64 {
	var result uint64

	buffer := make([]byte, 8)
	n, err := r.r.Read(buffer)
	if err != nil {
		panic(err)
	}
	if n < 8 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	err = binary.Read(buf, binary.LittleEndian, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func (r *Reader) ReadInt32() int32 {
	var result int32

	buffer := make([]byte, 4)
	n, err := r.r.Read(buffer)
	if err != nil {
		panic(err)
	}
	if n < 4 {
		return 0
	}

	_, _ = r.b.Write(buffer)

	err = binary.Read(r.b, binary.LittleEndian, &result)
	if err != nil {
		panic(err)
	}
	r.b.Reset()
	return result
}

func (r *Reader) ReadUInt16() uint16 {
	var result uint16

	buffer := make([]byte, 2)
	n, err := r.r.Read(buffer)
	if err != nil {
		panic(err)
	}
	if n < 2 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	err = binary.Read(buf, binary.LittleEndian, &result)

	if err != nil {
		panic(err)
	}

	return result
}

func (r *Reader) ReadInt16() int16 {
	var result int16

	buffer := make([]byte, 2)
	n, err := r.r.Read(buffer)
	if err != nil {
		panic(err)
	}
	if n < 2 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	err = binary.Read(buf, binary.LittleEndian, &result)

	if err != nil {
		panic(err)
	}

	return result
}
func (r *Reader) ReadUInt8() uint8 {
	var result uint8

	buffer := make([]byte, 1)
	n, err := r.r.Read(buffer)
	if err != nil {
		panic(err)
	}
	if n < 1 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	err = binary.Read(buf, binary.LittleEndian, &result)
	if err != nil {
		panic(err)
	}
	return result
}
func (r *Reader) ReadInt8() int8 {
	var result int8

	buffer := make([]byte, 1)
	n, err := r.r.Read(buffer)
	if err != nil {
		panic(err)
	}
	if n < 1 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	err = binary.Read(buf, binary.LittleEndian, &result)
	if err != nil {
		panic(err)
	}
	return result
}
func (r *Reader) ReadString() string {
	var result []byte
	var secondByte byte
	for {
		firstByte, err := r.r.ReadByte()
		if err != nil {
			panic(err)
		}
		secondByte, err = r.r.ReadByte()
		if err != nil {
			panic(err)
		}

		if firstByte == 0x00 && secondByte == 0x00 {
			break
		} else {
			result = append(result, firstByte)
		}
	}

	return string(result)
}
