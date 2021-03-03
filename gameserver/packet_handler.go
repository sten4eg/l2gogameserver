package gameserver

import (
	"fmt"
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"log"
)

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
			clientpackets.NewprotocolVersion(data, client)
		case 43:
			clientpackets.NewAuthLogin(data, client, g.database)
		case 19:
			serverpackets.NewCharacterSuccess(client)
		case 12:
			clientpackets.NewCharacterCreate(data, g.database, client)
		case 18:
			clientpackets.NewCharSelected(data, client)

			_ = serverpackets.NewCharSelected(client.Account.Char[client.Account.CharSlot], client) // return charId , unused ? remove?

			x, y, _ := client.CurrentChar.GetXYZ()
			rg := models.GetRegion(x, y)
			rg.AddVisibleObject(client.CurrentChar)
			client.CurrentChar.CurrentRegion = rg
			g.addOnlineChar(client.CurrentChar)

		case 208:
			if len(data) >= 2 {
				switch data[0] {
				case 1:
					serverpackets.NewExSendManorList(client)
				case 54:
					serverpackets.NewCharSelectionInfo(g.database, client)
				}
			}

		case 193:
			serverpackets.NewObservationReturn(client.CurrentChar, client)
		case 108:
			serverpackets.NewShowMiniMap(client)
		case 17:
			serverpackets.NewUserInfo(client.CurrentChar, client)

			serverpackets.NewExBrExtraUserInfo(client)

			serverpackets.NewSendMacroList(client)

			serverpackets.NewItemList(client, g.database)

			serverpackets.NewExQuestItemList(client)

			serverpackets.NewGameGuardQuery(client)

			serverpackets.NewExGetBookMarkInfoPacket(client)

			serverpackets.NewShortCutInit(client)

			serverpackets.NewExBasicActionList(client)

			serverpackets.NewSkillList(client, g.database)

			serverpackets.NewHennaInfo(client)

			serverpackets.NewQuestList(client)

			serverpackets.NewStaticObject(client)

			var info models.PacketByte
			pkg := serverpackets.NewCharInfo(client.CurrentChar)
			info.SetB(pkg)
			g.BroadToAroundPlayers(client, info)

			//todo вынести это отсюдова
			charIds := models.GetAroundPlayers(client.CurrentChar)
			for _, v := range charIds {
				tt := g.OnlineCharacters.Char[v]
				pkgs := serverpackets.NewCharInfo(tt)
				client.Buffer.WriteSlice(pkgs)
			}

			log.Println("Send NewUserInfo")
		case 166:
			pkg := serverpackets.NewSkillCoolTime()
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
		case 15:
			location := clientpackets.NewMoveBackwardToLocation(data)
			pkg := serverpackets.NewMoveToLocation(location, client)
			var info models.PacketByte
			info.SetB(pkg)
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			g.BroadToAroundPlayers(client, info)

			log.Println("Send NewMoveToLocation")
		case 73:
			clientpackets.NewSay(data, g.OnlineCharacters, client.CurrentChar)

			//info.B = serverpackets.NewCharInfo(client.CurrentChar)
			//Broad(client, info)
		case 89:
			clientpackets.NewValidationPosition(data, client.CurrentChar)
		case 31:
			clientpackets.NewAction(data, client)
		case 72:
			clientpackets.NewRequestTargetCanceld(data, client)
		case 1:
			clientpackets.NewAttack(data, client)
		case 25:
			clientpackets.NewUseItem(data, client, g.database)

			//todo нужно подумать как это вынести и отправлять =((
			var info models.PacketByte
			pkg := serverpackets.NewCharInfo(client.CurrentChar)
			info.SetB(pkg)
			g.BroadToAroundPlayers(client, info)
		case 87:
			clientpackets.NewRequestRestart(data, client, g.database)
		case 57:
			clientpackets.NewRequestMagicSkillUse(data, client)
		default:
			log.Println("Not Found case with opcode: ", opcode)
		}

		client.SentToSend()
	}
}
