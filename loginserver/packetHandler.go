package loginserver

import (
	"fmt"
	"l2gogameserver/loginserver/network/gs2ls"
	"l2gogameserver/loginserver/network/ls2gs"
)

func (ls *LoginServer) HandlePacket(data []byte) {
	opCode := data[0]
	data = data[1:]
	fmt.Println(opCode)

	switch opCode {
	default:
		fmt.Printf("неопознаный опкод от логинсервера: %v\n", opCode)
	case 0x00:
		pubKey := ls2gs.InitLs(data)
		bfk := generateNewBlowFish()
		buf := gs2ls.BlowFishKey(pubKey, bfk)

		ls.Send(buf)
		ls.setBlowFish(bfk)
		buf = gs2ls.AuthRequest()
		ls.Send(buf)
	case 0x02:
		ls2gs.AuthResponse(data)
		buf := gs2ls.ServerStatus()
		ls.Send(buf)
	}
}
