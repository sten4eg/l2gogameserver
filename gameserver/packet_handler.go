package gameserver

import (
	"fmt"
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/models"
	"log"
)

// loop клиента в ожидании входящих пакетов
func (g *GameServer) handler(client *models.Client) {
	defer kickClient(client)

	for {
		opcode, data, err := client.Receive()

		if err != nil {
			fmt.Println(err)
			fmt.Println("Коннект закрыт")
			break // todo  return ?
		}
		log.Println("income ", opcode)
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
			go g.ChannelListener(client)

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
					log.Println("Не реализованный пакет: ", data[0])
				}
			}

		case 193:
			pkg := clientpackets.RequestObserverEnd(client, data)
			client.SSend(pkg)
		case 108:
			pkg := clientpackets.RequestShowMiniMap(client, data)
			client.SSend(pkg)
		case 17:
			pkg := clientpackets.RequestEnterWorld(client, data)
			client.SSend(pkg)
			g.BroadCastUserInfoInRadius(client, 2000)
		case 166:
			pkg := clientpackets.RequestSkillCoolTime(client, data)
			client.SSend(pkg)
		case 15:
			pkg := clientpackets.MoveBackwardToLocation(client, data)
			client.SSend(pkg)
			//var info utils.PacketByte
			//info.SetB(pkg)
			//
			//client.Buffer.WriteSlice(pkg)
			//
			//
			//client.SaveAndCryptDataInBufferToSend(true)
			//
			//g.BroadToAroundPlayers(client, info)
			//
			//log.Println("Send MoveToLocation")
		case 73:
			_ = clientpackets.Say(client, data) //todo

			//info.B = serverpackets.CharInfo(client.CurrentChar)
			//Broad(client, info)
		case 89:
			pkg := clientpackets.ValidationPosition(data, client.CurrentChar)
			client.SSend(pkg)
		case 31:
			pkg := clientpackets.Action(data, client)
			client.SSend(pkg)
		case 72:
			pkg := clientpackets.RequestTargetCancel(data, client)
			client.SSend(pkg)
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
		default:
			log.Println("Not Found case with opcode: ", opcode)
		}

	}
}
