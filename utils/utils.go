package utils

import (
	"golang.org/x/exp/constraints"
	"l2gogameserver/packets"
	"reflect"
	"sync"
	"unsafe"
)

const maxCap = 1 << 11 // 2 kB

var PacketBytePool = sync.Pool{
	New: func() interface{} {
		return new(PacketByte)
	},
}

type PacketByte struct {
	data []byte
}

// Release если cap() больше maxCap то лучше его не ложить обратно в пул
// а дождаться пока GC его уничтожит,
// использование packetByte с cap() большого размера не эффективно
func (b *PacketByte) Release() {
	if cap(b.data) <= maxCap {
		b.data = b.data[:0]
		PacketBytePool.Put(b)
	}
}

func (b *PacketByte) Free() {
	b.data = b.data[:0]
}

func GetPacketByte() (b *PacketByte) {
	return PacketBytePool.Get().(*PacketByte)
}

// GetData получение массива байт из packetByte
func (b *PacketByte) GetData() []byte {
	cl := make([]byte, len(b.data))
	_ = copy(cl, b.data)
	return cl
}

// SetData копирует массив байт в packetByte
func (b *PacketByte) SetData(v []byte) {
	cl := make([]byte, len(v))
	b.data = cl
	copy(b.data, v)
}

func (b *PacketByte) SetDataBuf(v *packets.Buffer) {
	cl := make([]byte, len(v.Bytes()))
	b.data = cl
	copy(b.data, v.Bytes())
	packets.Put(v)
}

func Contains[T constraints.Integer](slice []T, need T) bool {
	for i := range slice {
		if slice[i] == need {
			return true
		}
	}
	return false
}

func BoolToInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

func I2B[T constraints.Integer](i T) bool {
	return i != 0
}

func B2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func S2b(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}
