package gameserver

import (
	"fmt"
	"l2gogameserver/config"
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
	"log"
)

// loop клиента в ожидании входящих пакетов
func (g *GameServer) handler(client *models.Client) {
	for {
		opcode, data, err := client.Receive()
		//defer kickClient(client)
		if err != nil {
			fmt.Println(err)
			g.charOffline(client)
			break // todo  return ?
		}
		if config.Get().Debug.ShowPackets == true {
			log.Println("Client->Server: #", opcode, packets.GetNamePacket(opcode))
		}
		switch opcode {
		case 14:
			pkg := clientpackets.ProtocolVersion(data, client)
			client.SSend(pkg)
		case 43:
			pkg := clientpackets.AuthLogin(data, client)
			client.SSend(pkg)
		case 19:
			pkg := clientpackets.RequestNewCharacter(client, data)
			client.SSend(pkg)
		case 12:
			pkg := clientpackets.CharacterCreate(data, client)
			client.SSend(pkg)
		case 18:
			pkg := clientpackets.CharSelected(data, client)
			client.SSend(pkg)
			g.addOnlineChar(client.CurrentChar)
		case 208:
			if len(data) >= 2 {
				switch data[0] {
				case 1:
					pkg := clientpackets.RequestManorList(client, data)
					client.SSend(pkg)
				case 54:
					pkg := clientpackets.RequestGoToLobby(client, data)
					client.SSend(pkg)
				case 13:
					pkg := clientpackets.RequestAutoSoulShot(data, client)
					client.SSend(pkg)
				case 36:
					clientpackets.RequestSaveInventoryOrder(client, data)
				default:
					if config.Get().Debug.ShowPackets == true {
						log.Println("Не реализованный пакет: ", data[0], packets.GetNamePacket(data[0]))
					}
				}
			}

		case 86:
			if len(data) >= 2 {
				log.Println(data[0])
				switch data[0] {
				case 0: //посадить персонажа на жопу
					pkg0 := clientpackets.ChangeWaitType(client)
					client.SSend(pkg0)
				}

			}

		case 23:
			pkg := clientpackets.DropItem(client, data)
			client.SSend(pkg)

			pkgInventoryUpdate := clientpackets.InventoryUpdate(client, client.CurrentChar.ObjectId, models.UpdateTypeModify)
			client.SSend(pkgInventoryUpdate)

		case 193:
			pkg := clientpackets.RequestObserverEnd(client, data)
			client.SSend(pkg)
		case 108:
			pkg := clientpackets.RequestShowMiniMap(client, data)
			client.SSend(pkg)
		case 17:
			pkg := clientpackets.RequestEnterWorld(client, data)
			client.SSend(pkg)
			//g.BroadCastUserInfoInRadius(client, 2000)
			g.GetCharInfoAboutCharactersInRadius(client, 2000)
			go g.ChannelListener(client)
			go g.MoveListener(client)
			go g.NpcListener(client)
		case 166:
			pkg := clientpackets.RequestSkillCoolTime(client, data)
			client.SSend(pkg)
		case 15:
			pkg := clientpackets.MoveBackwardToLocation(client, data)
			g.Checkaem(client, pkg)

		case 73:
			say := clientpackets.Say(client, data)
			g.BroadCastChat(client, say)
		case 89:
			pkg := clientpackets.ValidationPosition(data, client.CurrentChar)
			//g.Checkaem(client, pkg)
			client.SSend(pkg)
		case 31:
			pkg, objectId, actionId, reAppeal := clientpackets.Action(data, client)
			client.SSend(pkg)

			npc, npcx, npcy, npcz, err := models.GetNpcObject(objectId)
			if err != nil {
				log.Println(err)
			}

			//Прост тест вызова HTML при клике
			if actionId == 1 {
				NpcHtmlMessage := clientpackets.NpcHtmlMessage(client, npc.NpcId)
				client.SSend(NpcHtmlMessage)
			}
			//Если повторный клик по нпц
			if reAppeal {
				//npc, npcx, npcy, npcz, err := models.GetNpcObject(objectId)
				//if err != nil {
				//	log.Println(err)
				//}
				x, y, z := client.CurrentChar.GetXYZ()
				distance := models.CalculateDistance(npcx, npcy, npcz, x, y, z, false, false)
				_, _ = distance, npc

				//подбегаем
				if distance <= 60 {
					log.Println("Расстояние до NPC подходящее")
					if models.GetDialogNPC(npc.Type) == 0 {
						//НПЦ для разговора, открываем диалог
						//Пускай макс. дистанция разговора будет 60 поинтов
						//Пока откроем ID нпц
						NpcHtmlMessage := clientpackets.NpcHtmlMessage(client, npc.NpcId)
						client.SSend(NpcHtmlMessage)
					} else {
						//бьем нпц
						client.SSend(clientpackets.Attack(data, client))
					}
				} else {
					log.Println("Расстояние до NPC слишком больше, необходимо подбежать")
					pkg2 := clientpackets.MoveToLocation(client, npcx, npcy, npcz)
					g.Checkaem(client, pkg2)
				}

			}

		case 72:
			pkg := clientpackets.RequestTargetCancel(data, client)
			client.SSend(pkg)
		case 114:
			log.Println(data)
			clientpackets.MoveToPawn(client, data)
		case 1:
			pkg := clientpackets.Attack(data, client)
			client.SSend(pkg)
		case 25:
			pkg := clientpackets.UseItem(client, data)
			client.SSend(pkg)
		case 87:
			pkg := clientpackets.RequestRestart(data, client)
			client.SSend(pkg)
		case 57:
			pkg := clientpackets.RequestMagicSkillUse(data, client)
			client.SSend(pkg)
		case 61:
			pkg := clientpackets.RequestShortCutReg(data, client)
			client.SSend(pkg)
		case 63:
			pkg := clientpackets.RequestShortCutDel(data, client)
			client.SSend(pkg)
		case 80:
			pkg := clientpackets.RequestSkillList(client, data)
			client.SSend(pkg)
		case 20:
			pkg := clientpackets.RequestItemList(client, data)
			client.SSend(pkg)
		case 205:
			pkg := clientpackets.RequestMakeMacro(client, data)
			client.SSend(pkg)
		default:
			if config.Get().Debug.ShowPackets == true {
				log.Println("Not Found case with opcode: ", opcode, packets.GetNamePacket(opcode))
			}
		}

	}
}
