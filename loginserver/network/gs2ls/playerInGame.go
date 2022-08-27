package gs2ls

import "l2gogameserver/packets"

func PlayerInGame(accounts []string) *packets.Buffer {
	buf := packets.Get()
	buf.WriteSingleByte(0x02)
	buf.WriteH(int16(len(accounts)))
	for i := range accounts {
		buf.WriteS(accounts[i])
	}
	return buf
}
