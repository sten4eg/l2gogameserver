package main

import (
	"encoding/binary"
	"math"
	"testing"
	"unicode/utf16"
)

func BenchmarkNew(b *testing.B) {
	var bu Buffer
	for i := 0; i < b.N; i++ {
		bu.WriteSOld("test")
	}
}
func BenchmarkSum(b *testing.B) {
	var bu Buffer
	for i := 0; i < b.N; i++ {
		bu.WriteSNew("test")
	}
}

const EmptyByte byte = 0

type Buffer struct {
	b []byte
}

func (b *Buffer) WriteSOld(value string) {
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

func (b *Buffer) WriteSNew(value string) {
	runes := []rune(value)
	// Предварительное выделение памяти
	b.b = append(b.b, make([]byte, len(runes)*2+2)...)

	offset := len(b.b) - len(runes)*2 - 2
	for i, r := range runes {
		v := utf16.Encode([]rune{r})[0]
		binary.LittleEndian.PutUint16(b.b[offset+i*2:], v)
	}
	// Последние 2 байта уже 0
}
