package main

import (
	"l2gogameserver/packets"
	"testing"
)

var Data []byte

func init() {
	Data = []byte{123, 0, 0, 0}
}

//func BenchmarkNew(b *testing.B) {
//	for i:=0;i<b.N;i++ {
//		var read = packets.NewReader(Data)
//	//	_ = read.NormTema()
//	}
//
//}
func BenchmarkOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var read = packets.NewReader(Data)
		_ = read.ReadInt32()
	}
}
