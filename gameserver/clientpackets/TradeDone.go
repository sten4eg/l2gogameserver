package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

// TradeDone Игрок подтвердил сделку
func TradeDone(data []byte, client interfaces.ReciverAndSender) {
	var packet = packets.NewReader(data)
	response := packet.ReadInt32() // 1 - пользователь нажал ОК, 0 пользователь отменил трейд

	player := client.GetCurrentChar()
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
				endTrade(client)
				return
			}
		}

		//todo тут еще несколько проверок

		partner := trade.GetPartner()
		if player.CalculateDistanceTo(partner.GetX(), partner.GetY(), partner.GetZ(), true, false) > 150 {
			endTrade(client)
			return
		}
		_, isTradeConfirmed := trade.Confirmed()
		if isTradeConfirmed {
			tradeConfirmed(client, trade.GetPartner())
		}
	} else {
		needSendCancelToMe, _ := client.GetCurrentChar().CancelActiveTrade()
		if needSendCancelToMe {
			endTrade(client)
		}
	}

}

func tradeConfirmed(client interfaces.ReciverAndSender, partner interfaces.CharacterI) {
	buff := packets.Get()
	smg := sysmsg.C1ConfirmedTrade
	smg.AddString(partner.GetName())
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(sysmsg.SystemMessage(smg)))
	client.Send(buff.Bytes())
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
