package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
)

func TradeDone(num int32) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)
	buffer.WriteSingleByte(0x1C)
	buffer.WriteD(num)
	return buffer.Bytes()
}

//Отмена торга
func TradeCancel(client, player2 interfaces.CharacterI) {
	tradeSend(client, false)
	tradeSend(player2, false)
}

//Пакет обмена
func TradeOK(client, player2 interfaces.CharacterI) {
	tradeSend(client, true)
	tradeSend(player2, true)
}

func tradeSend(client interfaces.CharacterI, tradeOK bool) {
	var trade int32
	if tradeOK {
		trade = 1
	} else {
		trade = 0
	}
	buffer := packets.Get()
	defer packets.Put(buffer)
	buffer.WriteSingleByte(0x1C)
	buffer.WriteD(trade)

	ut := utils.GetPacketByte()
	ut.SetData(buffer.Bytes())
	client.EncryptAndSend(ut.GetData())
}
