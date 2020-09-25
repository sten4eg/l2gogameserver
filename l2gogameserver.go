package main

import (
	"l2gogameserver/gameserver"
)

func main() {
	//q := new(packets.Buffer)
	//q.WriteD(-75122)
	//q.WriteD(258213)
	//q.WriteD(-3108)
	//log.Fatal(1)
	x := gameserver.New()
	x.Init()
	x.Start()
}
