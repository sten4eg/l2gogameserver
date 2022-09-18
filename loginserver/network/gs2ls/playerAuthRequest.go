package gs2ls

import "l2gogameserver/packets"

func PlayerAuthRequest(login string, playOk1, playOk2, loginOk1, loginOk2 uint32) *packets.Buffer {
	buf := packets.Get()
	buf.WriteSingleByte(0x05)
	buf.WriteS(login)
	buf.WriteDU(playOk1)
	buf.WriteDU(playOk2)
	buf.WriteDU(loginOk1)
	buf.WriteDU(loginOk2)
	return buf
}
