package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
)

func TradeOtherDone(client interfaces.CharacterI) {
	buffer := packets.Get()
	defer packets.Put(buffer)
	buffer.WriteSingleByte(0x82)

	ut := utils.GetPacketByte()
	ut.SetData(buffer.Bytes())
	client.EncryptAndSend(ut.GetData())
}
