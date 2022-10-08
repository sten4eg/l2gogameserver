package broadcast

import (
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/chat"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
)

// ToAroundPlayerInRadius отправляет всем персонажам в радиусе radius
// информацию из пакета pkg
func ToAroundPlayerInRadius(my interfaces.ReciverAndSender, pkg *utils.PacketByte, radius int32) {
	charsIds := models.GetAroundPlayersInRadius(my.GetCurrentChar(), radius, float64(radius))
	for i := range charsIds {
		charsIds[i].EncryptAndSend(pkg.GetData())
	}
}

func BroadCastPkgToAroundPlayer(my interfaces.ReciverAndSender, pkg []byte) {
	pb := utils.GetPacketByte()
	pb.SetData(pkg)
	BroadCastToAroundPlayers(my, pb)
	pb.Release()
}

func BroadCastToAroundPlayers(my interfaces.ReciverAndSender, pkg *utils.PacketByte) {
	charsIds := models.GetAroundPlayer(my.GetCurrentChar())
	for i := range charsIds {
		charsIds[i].EncryptAndSend(pkg.GetData())
	}
}

func BroadCastToAroundPlayersWithoutSelf(my interfaces.ReciverAndSender, pkg *utils.PacketByte) {
	chardIds := models.GetAroundPlayer(my.GetCurrentChar())
	for i := range chardIds {
		if chardIds[i].GetObjectId() != my.GetCurrentChar().GetObjectId() {
			chardIds[i].EncryptAndSend(pkg.GetData())
		}
	}
}

func BroadCastPkgToAroundPlayersWithoutSelf(my interfaces.ReciverAndSender, pkg []byte) {
	pb := utils.GetPacketByte()
	pb.SetData(pkg)
	BroadCastToAroundPlayersWithoutSelf(my, pb)
	pb.Release()
}

func BroadCastBufferToAroundPlayers(my interfaces.ReciverAndSender, buffer *packets.Buffer) {
	pb := utils.GetPacketByte()
	pb.SetData(buffer.Bytes())

	BroadCastToAroundPlayers(my, pb)

	pb.Release()
	packets.Put(buffer)
}

func BroadCastBufferToAroundPlayersWithoutSelf(my interfaces.ReciverAndSender, buffer *packets.Buffer) {
	pb := utils.GetPacketByte()
	pb.SetData(buffer.Bytes())

	BroadCastToAroundPlayersWithoutSelf(my, pb)

	pb.Release()
	packets.Put(buffer)
}

// BroadCastUserInfoInRadius отправляет всем персонажам в радиусе radius
// информацию о персонаже, Самому персонажу отправляет полный UserInfo
func BroadCastUserInfoInRadius(me interfaces.ReciverAndSender, radius int32) {
	ui := serverpackets.UserInfo(me.GetCurrentChar())
	me.EncryptAndSend(ui)

	charsIds := models.GetAroundPlayersInRadius(me.GetCurrentChar(), radius, float64(radius))
	if len(charsIds) == 0 {
		return
	}

	ci := utils.GetPacketByte()
	defer ci.Release()

	ci.SetData(serverpackets.CharInfo(me.GetCurrentChar()))

	exUi := utils.GetPacketByte()
	defer exUi.Release()
	exUi.SetData(serverpackets.ExBrExtraUserInfo(me.GetCurrentChar()))

	//g.OnlineCharacters.Mu.Lock()
	for i := range charsIds {
		gameserver.OnlineCharacters.Char[charsIds[i].GetObjectId()].Conn.EncryptAndSend(ci.GetData())
		gameserver.OnlineCharacters.Char[charsIds[i].GetObjectId()].Conn.EncryptAndSend(exUi.GetData())
	}
	//g.OnlineCharacters.Mu.Unlock()
}

func BroadcastUserInfo(client interfaces.ReciverAndSender) {
	pkg := serverpackets.UserInfo(client.GetCurrentChar())
	client.EncryptAndSend(pkg)

	pkg2 := serverpackets.CharInfo(client.GetCurrentChar())
	BroadCastPkgToAroundPlayersWithoutSelf(client, pkg2)

	pkg3 := serverpackets.ExBrExtraUserInfo(client.GetCurrentChar())
	BroadCastPkgToAroundPlayer(client, pkg3)
}

func BroadCastChat(me interfaces.ReciverAndSender, say models.Say) {
	pb := utils.GetPacketByte()
	defer pb.Release()

	switch say.Type {
	case chat.All:
		cs := serverpackets.CreatureSay(&say, me.GetCurrentChar())
		pb.SetData(cs)
		me.EncryptAndSend(pb.GetData())
		ToAroundPlayerInRadius(me, pb, chat.AllChatRange)
	case chat.Tell:
		cs := serverpackets.CreatureSay(&say, me.GetCurrentChar())
		pb.SetData(cs)
		ok := BroadCastToCharacterByName(pb, say.To)
		if ok {
			me.EncryptAndSend(pb.GetData())
		} else {
			// systemMSG что не найден перс
		}
	case chat.Shout:
		cs := serverpackets.CreatureSay(&say, me.GetCurrentChar())
		pb.SetData(cs)
		me.EncryptAndSend(pb.GetData())

		ToAroundPlayerInRadius(me, pb, chat.ShoutChatRange)
		//todo что за SpecialCommand ?
		//case chat.SpecialCommand:
		//	if me.CurrentChar.Target == 0 {
		//		return
		//	}
		//	qwe := g.OnlineCharacters.Char[me.CurrentChar.Target]
		//	q := models.CalculateDistance(qwe.Coordinates.X, qwe.Coordinates.Y, qwe.Coordinates.Z, me.CurrentChar.Coordinates.X, me.CurrentChar.Coordinates.Y, me.CurrentChar.Coordinates.Z, false, false)
		//	say.Text = fmt.Sprintf("%f", q)
		//	say.tType = chat.All
		//
		//	cs := serverpackets.CreatureSay(&say, me.CurrentChar)
		//	pb.SetData(cs)
		//	me.Send(me.CryptAndReturnPackageReadyToShip(pb.GetData()))
		//	ToAroundPlayerInRadius(me, pb, chat.AllChatRange)
	}
}

// BroadCastToCharacterByName отправляет pkg персонажу с ником to
// true если отправлен, false если персонаж не найден
func BroadCastToCharacterByName(pkg *utils.PacketByte, to string) bool {
	gameserver.OnlineCharacters.Mu.Lock()
	defer gameserver.OnlineCharacters.Mu.Unlock()

	conn := gameserver.GetNetConnByCharacterName(to)
	if conn != nil {
		conn.EncryptAndSend(pkg.GetData())
		return true
	}

	return false
}

// SendCharInfoAboutCharactersInRadius отправляет me CharInfo персонажей
// в радиусе radius
func SendCharInfoAboutCharactersInRadius(me interfaces.ReciverAndSender, radius int32) {
	charsIds := models.GetAroundPlayersInRadius(me.GetCurrentChar(), radius, float64(radius))
	for i := range charsIds {
		me.EncryptAndSend(serverpackets.CharInfo(charsIds[i]))
	}
}

// SendCharInfoAboutCharactersAround отправляет me CharInfo персонажей
func SendCharInfoAboutCharactersAround(me *models.ClientCtx) {
	charsIds := models.GetAroundPlayer(me.CurrentChar)
	for i := range charsIds {
		me.EncryptAndSend(serverpackets.CharInfo(charsIds[i]))
	}
}

func Checkaem(client interfaces.ReciverAndSender, l models.BackwardToLocation) {
	ut := utils.GetPacketByte()
	ut.SetData(serverpackets.MoveToLocation(&l, client.GetCurrentChar()))

	client.EncryptAndSend(ut.GetData())
	BroadCastToAroundPlayers(client, ut)
}

func GetCharacterByObjectId(id int32) interfaces.CharacterI {
	return gameserver.GetNetConnByCharObjectId(id)
}
