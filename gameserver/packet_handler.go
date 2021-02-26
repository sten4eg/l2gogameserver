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

			rg := models.GetRegion(client.CurrentChar.Coordinates.X, client.CurrentChar.Coordinates.Y)
			rg.AddVisibleObject(client.CurrentChar)
			client.CurrentChar.CurrentRegion = rg
			g.addOnlineChar(client.CurrentChar)
			client.SimpleSend(client.Buffer.Bytes(), true)

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
			pkg := serverpackets.NewUserInfo(client.CurrentChar)
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			pkg = serverpackets.NewExBrExtraUserInfo()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			pkg = serverpackets.NewSendMacroList()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			serverpackets.NewItemList(client, g.database)

			pkg = serverpackets.NewExQuestItemList()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewGameGuardQuery()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewExGetBookMarkInfoPacket()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewShortCutInit()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewExBasicActionList()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewSkillList(client.CurrentChar.CharId, g.database)
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewHennaInfo()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewQuestList()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewStaticObject()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
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
			g.Broad(client, info)

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
		case 87:
			clientpackets.NewRequestRestart(data, client, g.database)
		case 57:
			clientpackets.NewRequestMagicSkillUse(data, client)
		default:
			log.Println("Not Found case with opcode: ", opcode)
		}
	}
}
