package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

// TradeDone Игрок подтвердил сделку
func TradeDone(data []byte, client interfaces.ReciverAndSender) {
	var packet = packets.NewReader(data)
	response := packet.ReadInt32() // 1 - пользователь нажал ОК, 0 пользователь отменил трейд

	player := client.GetCurrentChar()
	if player == nil {
		return
	}
	// TODO проверка на флуд трейдами
	trade := player.GetActiveTradeList()
	if trade == nil {
		logger.Warning.Println("player.GetActiveTradeList == nil")
		return
	}
	if trade.IsLocked() {
		return
	}

	if response == 1 {
		if trade.GetPartner() == nil {
			needSendCancelToMe, _ := player.CancelActiveTrade()
			if needSendCancelToMe {
				player.EncryptAndSend(sysmsg.SystemMessage(sysmsg.TargetIsNotFoundInTheGame))
				return
			}
		}

		if trade.GetOwner().GetActiveEnchantItemId() != models.IdNone || trade.GetPartner().GetActiveEnchantItemId() != models.IdNone {
			return
		}

		//todo тут еще несколько проверок

		partner := trade.GetPartner()
		if player.CalculateDistanceTo(partner.GetX(), partner.GetY(), partner.GetZ(), true, false) > 150 {
			endTrade(client)
			return
		}
		_, isTradeConfirmed, tradeDone, success := trade.Confirmed()
		if isTradeConfirmed {
			tradeConfirmed(trade.GetPartner(), trade.GetOwner())
		}
		if tradeDone {
			finishTrade(player, success)
			finishTrade(partner, success)
		}
	} else {
		partner := player.GetActiveTradeList().GetPartner()
		needSendCancelToMe, _ := client.GetCurrentChar().CancelActiveTrade()
		if needSendCancelToMe {
			cancelTrade(player, partner)
			cancelTrade(partner, player)
		}
	}

}

func tradeConfirmed(client interfaces.CharacterI, partner interfaces.CharacterI) {
	buff := packets.Get()
	msg := sysmsg.C1ConfirmedTrade
	msg.AddString(partner.GetName())
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(sysmsg.SystemMessage(msg)))
	client.Send(buff.Bytes())
	serverpackets.TradeOtherDone(client)
	packets.Put(buff)
}
func endTrade(client interfaces.ReciverAndSender) {
	buff := packets.Get()

	pkg := serverpackets.TradeDone(0)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	smg := sysmsg.C1CanceledTrade
	smg.AddString(client.GetCurrentChar().GetActiveTradeList().GetPartner().GetName())
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(sysmsg.SystemMessage(smg)))

	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(sysmsg.SystemMessage(sysmsg.TargetIsNotFoundInTheGame)))
	client.Send(buff.Bytes())

	packets.Put(buff)
}

func finishTrade(c interfaces.CharacterI, successful bool) {
	//c.EncryptAndSend(serverpackets.InventoryUpdate(c.GetInventory().GetItemsWithUpdatedType()))

	c.OnTradeFinish()
	c.EncryptAndSend(serverpackets.TradeDone(1))
	if successful {
		c.EncryptAndSend(sysmsg.SystemMessage(sysmsg.TradeSuccessful))
	}
}

func cancelTrade(c interfaces.CharacterI, partner interfaces.CharacterI) {
	buff := packets.Get()

	pkg := serverpackets.TradeDone(0)
	buff.WriteSlice(c.CryptAndReturnPackageReadyToShip(pkg))

	msg := sysmsg.C1CanceledTrade
	msg.AddString(partner.GetName())
	buff.WriteSlice(c.CryptAndReturnPackageReadyToShip(sysmsg.SystemMessage(msg)))
	c.Send(buff.Bytes())
	packets.Put(buff)

}
