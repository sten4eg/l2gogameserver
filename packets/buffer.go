package packets

import (
	"encoding/binary"
	"l2gogameserver/config"
	"l2gogameserver/data/logger"
	"math"
	"path/filepath"
	"runtime"
	"strings"
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
	if config.GetDebug().IsShowPackets() {
		// Получаем информацию о вызывающей функции
		pc, file, line, ok := runtime.Caller(1)
		if ok && value != 0x00 {
			funcName := runtime.FuncForPC(pc).Name()
			parts := strings.Split(funcName, ".")
			shortFuncName := parts[len(parts)-1]
			fileName := filepath.Base(file)
			logger.Info.Printf("Server->Client: %s (%s:%d) (опкод: 0x%02X)",
				shortFuncName, fileName, line, value)
		}
	}
	b.b = append(b.b, value)
}

func (b *Buffer) WriteS(value string) {
	runes := []rune(value)
	// Предварительное выделение памяти
	b.b = append(b.b, make([]byte, len(runes)*2+2)...)

	offset := len(b.b) - len(runes)*2 - 2
	for i, r := range runes {
		v := utf16.Encode([]rune{r})[0]
		binary.LittleEndian.PutUint16(b.b[offset+i*2:], v)
	}
}
