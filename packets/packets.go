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

func (b *Buffer) WriteH(value uint16) {
	err := binary.Write(b, binary.LittleEndian, value)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *Buffer) WriteD(value uint32) {
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

type Reader struct {
	*bytes.Reader
}

func NewReader(buffer []byte) *Reader {
	return &Reader{bytes.NewReader(buffer)}
}

func (r *Reader) ReadBytes(number int) []byte {
	buffer := make([]byte, number)
	n, _ := r.Read(buffer)
	if n < number {
		return []byte{}
	}

	return buffer
}

func (r *Reader) ReadUInt64() uint64 {
	var result uint64

	buffer := make([]byte, 8)
	n, err := r.Read(buffer)
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

func (r *Reader) ReadUInt32() uint32 {
	var result uint32

	buffer := make([]byte, 4)
	n, err := r.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	if n < 4 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	err = binary.Read(buf, binary.LittleEndian, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (r *Reader) ReadUInt16() uint16 {
	var result uint16

	buffer := make([]byte, 2)
	n, err := r.Read(buffer)
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
	n, err := r.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	if n < 1 {
		return 0
	}

	buf := bytes.NewBuffer(buffer)

	err = binary.Read(buf, binary.LittleEndian, &result)

	return result
}

func (r *Reader) ReadString() string {
	var result []byte
	var secondByte byte
	for {
		firstByte, err := r.ReadByte()
		if err != nil {
			log.Fatal(err)
		}
		secondByte, err = r.ReadByte()
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
