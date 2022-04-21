package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/trade"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"log"
)

//Когда игрок добавляет предмет в трейде
func AddTradeItem(data []byte, client *models.Client) []byte {
	var packet = packets.NewReader(data)

	_ = packet.ReadInt32()
	objectid := packet.ReadInt32()
	count := packet.ReadUInt64()

	item, toUser, ok := trade.AddItemTrade(client, objectid, int64(count))
	if !ok {
		log.Println("Не добавлен предмет")
		return nil
	}
	buff := packets.Get()
	defer packets.Put(buff)
	pkg := serverpackets.TradeOwnOAdd(client, item, count)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	TradeOtherAdd(toUser, item, count)

	return buff.Bytes()

}

//Шлется инфа для другого игрока в обмене
func TradeOtherAdd(toUser *models.Client, item *models.MyItem, count uint64) {

	pkg := serverpackets.TradeOtherAdd(toUser, item, count)
	ut := utils.GetPacketByte()
	ut.SetData(pkg)
	toUser.EncryptAndSend(ut.GetData())

}
