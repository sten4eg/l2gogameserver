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

	activeChar := client.GetCurrentChar()
	if activeChar == nil {
		return
	}

	switch actionId {
	default:
		fmt.Printf("Неопознаный второй опкод %v в RequestActionUse\n", data[0])
	case 0:
		ChangeWaitType(client)
	case 10:
		tryOpenPrivateSellShop(client, false)
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
			c.EncryptAndSend(sysmsg.SystemMessage(sysmsg.NoPrivateStoreHere))
		}
		pkg := serverpackets.ActionFailed(client)
		client.EncryptAndSend(pkg)

	}
}
