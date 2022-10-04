package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func SetPrivateStoreMsgBuy(client interfaces.ReciverAndSender, data []byte) {
	reader := packets.NewReader(data)

	storeMsg := reader.ReadString()

	character := client.GetCurrentChar()
	if character == nil || character.GetBuyList() == nil {
		return
	}

	if storeMsg != "" && len(storeMsg) > MaxMsgLength {
		//Читер
		return
	}

	character.GetBuyList().SetTitle(storeMsg)
	pkg := serverpackets.PrivateStoreMsgBuy(character)
	client.SendBuf(pkg)

}
