package gameserver

import (
	"fmt"
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/utils"
	"log"
)

// loop клиента в ожидании входящих пакетов
func (g *GameServer) handler(client *models.Client) {
	defer kickClient(client)

	for {
		opcode, data, err := client.Receive()

		if err != nil {
			fmt.Println(err)
			fmt.Println("Closing the connection...")
			break
		}
		log.Println("income ", opcode)
		switch opcode {
		case 14:
			clientpackets.ProtocolVersion(data, client)
		case 43:
			clientpackets.AuthLogin(data, client)
		case 19:
			serverpackets.CharacterSuccess(client)
		case 12:
			clientpackets.CharacterCreate(data, client)
		case 18:
			clientpackets.CharSelected(data, client)

			g.addOnlineChar(client.CurrentChar)

		case 208:
			if len(data) >= 2 {
				switch data[0] {
				case 1:
					serverpackets.ExSendManorList(client)
				case 54:
					serverpackets.CharSelectionInfo(client)
				case 13:
					clientpackets.RequestAutoSoulShot(data, client)
				default:
					log.Println("Не реализованный пакет: ", data[0])
				}
			}

		case 193:
			serverpackets.ObservationReturn(client.CurrentChar, client)
		case 108:
			serverpackets.ShowMiniMap(client)
		case 17:
			serverpackets.UserInfo(client)

			serverpackets.ExBrExtraUserInfo(client)

			serverpackets.SendMacroList(client)

			serverpackets.ItemList(client)

			serverpackets.ExQuestItemList(client)

			serverpackets.GameGuardQuery(client)

			serverpackets.ExGetBookMarkInfoPacket(client)

			serverpackets.ExStorageMaxCount(client)

			serverpackets.ShortCutInit(client)

			serverpackets.ExBasicActionList(client)

			serverpackets.SkillList(client)

			serverpackets.HennaInfo(client)

			serverpackets.QuestList(client)

			serverpackets.StaticObject(client)

			var info utils.PacketByte
			pkg := serverpackets.CharInfo(client.CurrentChar)
			info.SetB(pkg)
			g.BroadToAroundPlayers(client, info)

			//todo вынести это отсюдова
			charIds := models.GetAroundPlayers(client.CurrentChar)
			for _, v := range charIds {
				tt := g.OnlineCharacters.Char[v]
				pkgs := serverpackets.CharInfo(tt)
				client.Buffer.WriteSlice(pkgs)
			}

			log.Println("Send UserInfo")
		case 166:
			pkg := serverpackets.SkillCoolTime()
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
		case 15:
			location := clientpackets.MoveBackwardToLocation(data)
			pkg := serverpackets.MoveToLocation(location, client)
			var info utils.PacketByte
			info.SetB(pkg)
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			g.BroadToAroundPlayers(client, info)

			log.Println("Send MoveToLocation")
		case 73:
			clientpackets.Say(data, g.OnlineCharacters, client.CurrentChar)

			//info.B = serverpackets.CharInfo(client.CurrentChar)
			//Broad(client, info)
		case 89:
			clientpackets.ValidationPosition(data, client.CurrentChar)
			serverpackets.NpcInfo(client)

		case 31:
			clientpackets.Action(data, client)
		case 72:
			clientpackets.RequestTargetCancel(data, client)
		case 1:
			clientpackets.Attack(data, client)
		case 25:
			clientpackets.UseItem(data, client)
			serverpackets.NpcHtmlMessage(client)
			//todo нужно подумать как это вынести и отправлять =((
			var info utils.PacketByte
			pkg := serverpackets.CharInfo(client.CurrentChar)
			info.SetB(pkg)
			g.BroadToAroundPlayers(client, info)
		case 87:
			clientpackets.RequestRestart(data, client)
		case 57:
			clientpackets.RequestMagicSkillUse(data, client)
		case 61:
			clientpackets.RequestShortCutReg(data, client)
		case 63:
			clientpackets.RequestShortCutDel(data, client)
		case 80:
			serverpackets.SkillList(client)
		default:
			log.Println("Not Found case with opcode: ", opcode)
		}

		client.SentToSend()
	}
}
