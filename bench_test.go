package main

import (
	"testing"
	"time"
)

var Sizze int64

func BenchmarkFirst(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Sizze = time.Now().Round(time.Millisecond).UnixNano() / 1e6

	}
}
