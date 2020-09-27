package serverpackets

import "l2gogameserver/packets"

func NewGameGuardQuery() []byte {
	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x74)
	buffer.WriteD(0x27533DD9)
	buffer.WriteD(0x2E72A51D)
	buffer.WriteD(0x2017038B)
	buffer.WriteDU(0xC35B1EA3)

	return buffer.Bytes()
}
