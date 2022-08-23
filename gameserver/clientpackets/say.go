package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/chat"
	"l2gogameserver/packets"
	"strings"
)

func Say(client interfaces.ReciverAndSender, data []byte) models.Say {
	var packet = packets.NewReader(data)
	var say models.Say

	say.Text = packet.ReadString()
	say.Type = packet.ReadInt32()
	//todo say.Target = реализовать

	buffer := packets.Get()
	if strings.HasPrefix(say.Text, ".") {
		say.Type = chat.SpecialCommand
		say.Text = "tok"
	}
	defer packets.Put(buffer)

	switch say.Type {
	case chat.All:
		return say
	case chat.Tell:
		say.To = packet.ReadString()
		return say
	case chat.Shout:
		return say
	}
	//BroadCastChat(client, say)

	return say
}

//func BroadCastChat(me *models.ClientCtx, say models.Say) {
//	pb := utils.GetPacketByte()
//	defer pb.Release()
//
//	switch say.tType {
//	case chat.All:
//		cs := serverpackets.CreatureSay(&say, me.CurrentChar)
//		pb.SetData(cs)
//		me.Send(me.CryptAndReturnPackageReadyToShip(pb.GetData()))
//		q := models.GetAroundPlayersObjIdInRadius(me.CurrentChar, chat.AllChatRange)
//		broadcast.BBBroadCastToAroundPlayersInRadius(q, pb)
//	}
//}
