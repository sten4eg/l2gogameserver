package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"l2gogameserver/packets"
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"golang.org/x/exp/constraints"
)

const (
	// maxCap максимальная емкость для пула пакетов
	maxCap = 1 << 11 // 2 kB

	// Размеры для генерации ID
	RequestIDLength = 16
)

var (
	// PacketBytePool пул для переиспользования байтовых массивов
	PacketBytePool = sync.Pool{
		New: func() interface{} {
			return new(PacketByte)
		},
	}
)

// PacketByte представляет байтовый массив для пакетов
type PacketByte struct {
	data []byte
}

// Release освобождает ресурсы пакета
// Если cap() больше maxCap то лучше его не ложить обратно в пул
// а дождаться пока GC его уничтожит,
// использование packetByte с cap() большого размера не эффективно
func (b *PacketByte) Release() {
	if cap(b.data) <= maxCap {
		b.data = b.data[:0]
		PacketBytePool.Put(b)
	}
}

// Free очищает данные пакета
func (b *PacketByte) Free() {
	b.data = b.data[:0]
}

// GetPacketByte получает PacketByte из пула
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

// SetDataBuf копирует данные из буфера в packetByte
func (b *PacketByte) SetDataBuf(v *packets.Buffer) {
	cl := make([]byte, len(v.Bytes()))
	b.data = cl
	copy(b.data, v.Bytes())
}

// Contains проверяет наличие элемента в слайсе
func Contains[T constraints.Integer](slice []T, need T) bool {
	for i := range slice {
		if slice[i] == need {
			return true
		}
	}
	return false
}

// ContainsString проверяет наличие строки в слайсе строк
func ContainsString(slice []string, need string) bool {
	for _, item := range slice {
		if item == need {
			return true
		}
	}
	return false
}

// BoolToInt32 конвертирует bool в int32
func BoolToInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// I2B конвертирует целое число в bool
func I2B[T constraints.Integer](i T) bool {
	return i != 0
}

// BoolToByte конвертирует bool в byte
func BoolToByte(flag bool) byte {
	if flag {
		return 1
	}
	return 0
}

// B2s конвертирует []byte в string без копирования
func B2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// S2b конвертирует string в []byte без копирования
func S2b(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

// GenerateRequestID генерирует уникальный ID для запроса
func GenerateRequestID() string {
	bytes := make([]byte, RequestIDLength)
	if _, err := rand.Read(bytes); err != nil {
		// В случае ошибки используем timestamp
		return fmt.Sprintf("%d", runtime.NumGoroutine())
	}
	return hex.EncodeToString(bytes)
}

// Min возвращает минимальное значение из двух
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max возвращает максимальное значение из двух
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Clamp ограничивает значение в заданном диапазоне
func Clamp[T constraints.Ordered](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Abs возвращает абсолютное значение
func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// IsPowerOfTwo проверяет, является ли число степенью двойки
func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// NextPowerOfTwo возвращает следующую степень двойки
func NextPowerOfTwo(n int) int {
	if n <= 0 {
		return 1
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	return n + 1
}

// ReverseString переворачивает строку
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// RemoveDuplicates удаляет дубликаты из слайса
func RemoveDuplicates[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}

	seen := make(map[T]bool)
	result := make([]T, 0, len(slice))

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// Chunk разбивает слайс на части указанного размера
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return nil
	}

	var chunks [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

// Filter фильтрует слайс по условию
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map применяет функцию к каждому элементу слайса
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = fn(item)
	}
	return result
}

// Reduce применяет функцию к элементам слайса, накапливая результат
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, item := range slice {
		result = fn(result, item)
	}
	return result
}

// SafeString возвращает безопасную строку (заменяет nil на пустую строку)
func SafeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// SafeInt возвращает безопасное целое число (заменяет nil на 0)
func SafeInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// SafeInt32 возвращает безопасное int32 (заменяет nil на 0)
func SafeInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// SafeInt64 возвращает безопасное int64 (заменяет nil на 0)
func SafeInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// SafeFloat64 возвращает безопасное float64 (заменяет nil на 0)
func SafeFloat64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

// SafeBool возвращает безопасное bool (заменяет nil на false)
func SafeBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
