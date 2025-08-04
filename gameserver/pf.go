package gameserver

import (
	"fmt"
	"os"
	"time"
	"unsafe"
)

type PathFindingNode struct {
	x, y, z                        int32
	n1, n2, n3, n4, n5, n6, n7, n8 int32
}

var Global []PathFindingNode

func ParsePathFin() {
	start := time.Now()
	file, err := os.ReadFile("gameserver/pathnode.bin")
	if err != nil {
		panic(err)
	}

	Global = unsafe.Slice((*PathFindingNode)(unsafe.Pointer(&file[0])), 12706146)
	qwe := Global
	_ = qwe

	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(len(Global))
	fmt.Println(cap(Global))

	fmt.Printf("ERERER")
}
