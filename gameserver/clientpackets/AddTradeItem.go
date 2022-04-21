package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/trade"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"log"
)

//AddTradeItem Когда игрок добавляет предмет в трейде
func AddTradeItem(data []byte, client interfaces.ReciverAndSender) {
	var packet = packets.NewReader(data)

	_ = packet.ReadInt32()
	objectId := packet.ReadInt32() // objectId предмета
	count := packet.ReadUInt64()

	item, toUser, ok := trade.AddItemTrade(client.GetCurrentChar(), objectId, int64(count))
	if !ok {
		log.Println("Не добавлен предмет")
		return
	}

	pkg := serverpackets.TradeOwnOAdd(item, count)
	client.EncryptAndSend(pkg)

	TradeOtherAdd(toUser, item, count)

}

//Шлется инфа для другого игрока в обмене
func TradeOtherAdd(toUser interfaces.CharacterI, item *models.MyItem, count uint64) {

	pkg := serverpackets.TradeOtherAdd(item, count)
	ut := utils.GetPacketByte()
	ut.SetData(pkg)
	toUser.EncryptAndSend(ut.GetData())

}
