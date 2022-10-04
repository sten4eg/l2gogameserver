package clientpackets

import (
	"fmt"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestActionUse(client interfaces.ReciverAndSender, data []byte) {
	packet := packets.NewReader(data)

	actionId := packet.ReadInt32()
	ctrlPressed := packet.ReadInt32() == 1
	shiftPressed := packet.ReadInt32() == 1

	_, _ = ctrlPressed, shiftPressed

	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	switch actionId {
	default:
		fmt.Printf("Неопознаный второй опкод %v в RequestActionUse\n", data[0])
	case 0:
		ChangeWaitType(client)
	case 10:
		tryOpenPrivateSellShop(client, false)
	case 28:
		tryOpenPrivateBuyStore(client)
	}

}

func tryOpenPrivateSellShop(client interfaces.ReciverAndSender, isPackageSale bool) {
	c := client.GetCurrentChar()
	if true { //TODO проверка на возможность создания магазина
		if c.GetPrivateStoreType() == privateStoreType.SELL || c.GetPrivateStoreType() == privateStoreType.SELL_MANAGE || c.GetPrivateStoreType() == privateStoreType.PACKAGE_SELL {
			c.SetPrivateStoreType(privateStoreType.NONE)
		}

		if c.GetPrivateStoreType() == privateStoreType.NONE {
			if c.IsSittings() {
				ChangeWaitType(client)
			}
			c.SetPrivateStoreType(privateStoreType.SELL_MANAGE)
			pkg := serverpackets.PrivateStoreManageListSell(c, isPackageSale)
			client.SendBuf(pkg)
		}

	} else {
		if false { //TODO проверка что персонаж находится в зоне, в которой нельзя торговать
			c.SendSysMsg(sysmsg.NoPrivateStoreHere)
		}
		pkg := serverpackets.ActionFailed(client)
		client.EncryptAndSend(pkg)

	}
}

func tryOpenPrivateBuyStore(client interfaces.ReciverAndSender) {
	c := client.GetCurrentChar()
	if true { //TODO проверка на возможность создания магазина
		if c.GetPrivateStoreType() == privateStoreType.BUY || c.GetPrivateStoreType() == privateStoreType.BUY_MANAGE {
			c.SetPrivateStoreType(privateStoreType.NONE)
		}
		if c.GetPrivateStoreType() == privateStoreType.NONE {
			if c.IsSittings() {
				ChangeWaitType(client)
			}
			c.SetPrivateStoreType(privateStoreType.BUY_MANAGE)
			pkg := serverpackets.PrivateStoreManageListBuy(c)
			client.SendBuf(pkg)
		}
	} else {
		if false { //TODO проверка что персонаж находится в зоне, в которой нельзя торговать
			c.SendSysMsg(sysmsg.NoPrivateStoreHere)
		}
		pkg := serverpackets.ActionFailed(client)
		client.EncryptAndSend(pkg)
	}
}
