package packets

import (
	"bytes"
	"encoding/binary"
	"log"
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
		log.Fatal(err)
	}
}

func (b *Buffer) WriteH(value int16) {
	err := binary.Write(b, binary.LittleEndian, value)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *Buffer) WriteQ(value int64) {
	err := binary.Write(b, binary.LittleEndian, value)
	if err != nil {
		log.Fatal(err)
	}
}
func (b *Buffer) WriteD(value int32) {
	err := binary.Write(b, binary.LittleEndian, value)
	if err != nil {
		log.Fatal(err)
	}
}
func (b *Buffer) WriteDU(value uint32) {
	err := binary.Write(b, binary.LittleEndian, value)
	if err != nil {
		log.Fatal(err)
	}
}
func (b *Buffer) WriteSlice(val []byte) {
	err := binary.Write(b, binary.LittleEndian, val)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *Buffer) WriteSingleByte(value byte) {
	err := b.WriteByte(value)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *Buffer) WriteS(value string) {

	val := append([]byte(value), 0, 0)
	err := binary.Write(b, binary.LittleEndian, val)
	if err != nil {
		log.Fatal(err)
	}
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
		log.Fatal(err)
	}
	if n < 8 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	err = binary.Read(buf, binary.LittleEndian, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (r *Reader) ReadInt32() int32 {
	var result int32

	buffer := make([]byte, 4)
	n, err := r.r.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	if n < 4 {
		return 0
	}

	_, _ = r.b.Write(buffer)

	err = binary.Read(r.b, binary.LittleEndian, &result)
	if err != nil {
		log.Fatal(err)
	}
	r.b.Reset()
	return result
}

func (r *Reader) ReadUInt16() uint16 {
	var result uint16

	buffer := make([]byte, 2)
	n, err := r.r.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	if n < 2 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	err = binary.Read(buf, binary.LittleEndian, &result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (r *Reader) ReadInt16() int16 {
	var result int16

	buffer := make([]byte, 2)
	n, err := r.r.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	if n < 2 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	err = binary.Read(buf, binary.LittleEndian, &result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
func (r *Reader) ReadUInt8() uint8 {
	var result uint8

	buffer := make([]byte, 1)
	n, err := r.r.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	if n < 1 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	err = binary.Read(buf, binary.LittleEndian, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (r *Reader) ReadString() string {
	var result []byte
	var secondByte byte
	for {
		firstByte, err := r.r.ReadByte()
		if err != nil {
			log.Fatal(err)
		}
		secondByte, err = r.r.ReadByte()
		if err != nil {
			log.Fatal(err)
		}

		if firstByte == 0x00 && secondByte == 0x00 {
			break
		} else {
			result = append(result, firstByte, secondByte)
		}
	}

	return string(result)
}
