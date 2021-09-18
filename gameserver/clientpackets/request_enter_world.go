package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestEnterWorld(client *models.Client, data []byte) []byte {

	buff := packets.Get()
	defer packets.Put(buff)

	pkg2 := serverpackets.ExBrExtraUserInfo(client.CurrentChar)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg2))

	pkg3 := serverpackets.SendMacroList(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg3))

	pkg4 := serverpackets.ItemList(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg4))

	pkg5 := serverpackets.ExQuestItemList(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg5))

	pkg6 := serverpackets.GameGuardQuery(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg6))

	pkg7 := serverpackets.ExGetBookMarkInfoPacket(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg7))

	pkg8 := serverpackets.ExStorageMaxCount(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg8))

	pkg9 := serverpackets.ShortCutInit(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg9))

	pkg10 := serverpackets.ExBasicActionList(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg10))

	pkg11 := serverpackets.SkillList(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg11))

	pkg12 := serverpackets.HennaInfo(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg12))

	pkg13 := serverpackets.QuestList(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg13))

	pkg14 := serverpackets.StaticObject(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg14))

	pkg15 := serverpackets.ShortBuffStatusUpdate(client) //todo test
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg15))

	return buff.Bytes()
}
