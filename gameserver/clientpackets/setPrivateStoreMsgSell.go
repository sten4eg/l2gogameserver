package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

const MaxMsgLength = 29

func SetPrivateStoreMsgSell(client interfaces.NewClientCtxInterface, data []byte) {
	reader := packets.NewReader(data)

	storeMsg := reader.ReadString()

	character := client.GetAccount().GetCurrentChar()
	if character == nil || character.GetSellList() == nil {
		return
	}

	if storeMsg != "" && len(storeMsg) > MaxMsgLength {
		//TODO читер бан
		return
	}

	character.GetSellList().SetTitle(storeMsg)
	pkg := serverpackets.PrivateStoreMsgSell(character)
	client.SendBuf(pkg)
}
