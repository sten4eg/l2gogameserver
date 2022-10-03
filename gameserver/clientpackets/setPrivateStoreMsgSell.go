package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

const MaxMsgLength = 29

func SetPrivateStoreMsgSell(client interfaces.ReciverAndSender, data []byte) {
	reader := packets.NewReader(data)

	storeMsg := reader.ReadString()

	activeChar := client.GetCurrentChar()
	if activeChar == nil || activeChar.GetSellList() == nil {
		return
	}

	if storeMsg != "" && len(storeMsg) > MaxMsgLength {
		//TODO читер бан
		return
	}

	activeChar.GetSellList().SetTitle(storeMsg)
	pkg := serverpackets.PrivateStoreMsgSell(activeChar)
	client.SendBuf(pkg)
}
