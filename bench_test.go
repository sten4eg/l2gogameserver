package main

import (
	"testing"
)

var Sizze int

func BenchmarkFirst(b *testing.B) {
	header := []byte{25, 250}

	for i := 0; i < b.N; i++ {
		size := 0
		size += int(header[0])
		size += int(header[1]) * 256
		Sizze = size
	}

}
func BenchmarkSecond(b *testing.B) {
	header := []byte{25, 250}

	for i := 0; i < b.N; i++ {
		size := 0
		size += int(header[0])
		size += int(header[1]) << 8
		Sizze = size
	}
}

func BenchmarkThree(b *testing.B) {
	header := []byte{25, 250}

	for i := 0; i < b.N; i++ {
		size := 0
		size = int(header[0]) | int(header[1])<<8
		Sizze = size
	}
}
