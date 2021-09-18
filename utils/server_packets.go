package utils

import (
	"l2gogameserver/gameserver/models"
)

var ServerPackets map[uint8][]func(client *models.Client)
var ClientPackets map[uint8]func(client *models.Client)

func SetupServerPackets() {
	//ServerPackets = make(map[uint8][]func(client *models.Client))
	////ServerPackets[17] = append(ServerPackets[17], serverpackets.UserInfo)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.ExBrExtraUserInfo)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.SendMacroList)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.ItemList)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.ExQuestItemList)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.GameGuardQuery)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.ExGetBookMarkInfoPacket)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.ExStorageMaxCount)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.ShortCutInit)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.ExBasicActionList)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.SkillList)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.HennaInfo)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.QuestList)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.StaticObject)
	//ServerPackets[17] = append(ServerPackets[17], serverpackets.ShortBuffStatusUpdate)
}

func SetupClientPackets() {
	ClientPackets = make(map[uint8]func(client *models.Client))

}
