package main

import (
	"testing"
)

func BenchmarkSimplest(b *testing.B) {

	for i := 0; i < b.N; i++ {
		w, e := by()
		_, _ = w, e
	}

}

func by() (opcode byte, data []byte) {
	f := make([]byte, 1024*1024)

	opcode = f[0]
	data = f[:1]
	return
}
