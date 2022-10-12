package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"strconv"
)

// AddTradeItem Когда игрок добавляет предмет в трейде
func AddTradeItem(data []byte, client interfaces.ReciverAndSender) {
	var packet = packets.NewReader(data)

	tradeId := packet.ReadInt32()
	objectId := packet.ReadInt32() // objectId предмета
	count := packet.ReadInt64()

	trade := client.GetCurrentChar().GetActiveTradeList()
	if trade == nil {
		logger.Warning.Println("Character: " + client.GetCurrentChar().GetName() + " requested item:" + strconv.Itoa(int(objectId)) + " add without active tradelist:" + strconv.Itoa(int(tradeId)))
		return
	}
	partner := trade.GetPartner()
	if partner == nil || broadcast.GetCharacterByObjectId(partner.GetObjectId()) == nil || partner.GetActiveTradeList() == nil {
		if partner == nil {
			logger.Warning.Println("Character:" + client.GetCurrentChar().GetName() + " requested invalid trade object: " + strconv.Itoa(int(objectId)))
		}
		sm := sysmsg.TargetIsNotFoundInTheGame
		client.EncryptAndSend(sysmsg.SystemMessage(sm))

		canceledTradeForMe, canceledTradeForPartner := client.GetCurrentChar().CancelActiveTrade()
		if canceledTradeForMe {
			buff := packets.Get()

			pkg := serverpackets.TradeDone(0)
			buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

			smg := sysmsg.C1CanceledTrade
			smg.AddString(client.GetCurrentChar().GetActiveTradeList().GetPartner().GetName())
			buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(sysmsg.SystemMessage(smg)))

			client.Send(buff.Bytes())

		}

		if canceledTradeForPartner {
			realPartner := client.GetCurrentChar().GetActiveTradeList().GetPartner()

			pkg := serverpackets.TradeDone(0)
			realPartner.EncryptAndSend(pkg)

			smg := sysmsg.C1CanceledTrade
			smg.AddString(realPartner.GetActiveTradeList().GetPartner().GetName())

			realPartner.EncryptAndSend(sysmsg.SystemMessage(smg))
		}
		return
	}

	if trade.IsConfirmed() || partner.GetActiveTradeList().IsConfirmed() {
		client.EncryptAndSend(sysmsg.SystemMessage(sysmsg.CannotAdjustItemsAfterTradeConfirmed))
		return
	}
	if !client.GetCurrentChar().ValidateItemManipulation(objectId) {
		client.EncryptAndSend(sysmsg.SystemMessage(sysmsg.NothingHappened))
		return
	}

	tItem := trade.AddItem(objectId, count, client.GetCurrentChar(), 0)
	if tItem != nil {
		client.EncryptAndSend(serverpackets.TradeOwnOAdd(tItem))
		trade.GetPartner().EncryptAndSend(serverpackets.TradeOtherAdd(tItem))
	}

}
