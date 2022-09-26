package gs2ls

import "l2gogameserver/packets"

func PlayerLogout(login string) *packets.Buffer {
	buf := packets.Get()
	buf.WriteSingleByte(0x03)
	buf.WriteS(login)
	return buf
}
