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
			//log.Println("Send NewKeyPacket")
		case 43:
			clientpackets.NewAuthLogin(data, client, g.database)
			//log.Println("Send NewCharSelectionInfo")
		case 19:
			serverpackets.NewCharacterSuccess(client) //toDo 4eknyt'
			client.SimpleSend(client.Buffer.Bytes(), true)
		//	log.Println("Send NewCharacterSuccess")
		case 12:
			reason := clientpackets.NewCharacterCreate(data, g.database, client.Account.Login)
			if reason != clientpackets.ReasonOk {
				serverpackets.NewCharCreateFail(client, reason)
				err := client.SimpleSend(client.Buffer.Bytes(), true)
				if err != nil {
					log.Println(err)
				}
			} else {
				serverpackets.NewCharCreateOk(client)
				err = client.SimpleSend(client.Buffer.Bytes(), true)
				if err != nil {
					log.Println(err)
				}
				log.Println("send NewCharCreateOk")
			}
		case 18:
			client.Account.CharSlot = clientpackets.NewCharSelected(data)
			pkg := serverpackets.NewSSQInfo()
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			log.Println("sendSSQ")

			_ = serverpackets.NewCharSelected(client.Account.Char[client.Account.CharSlot], client) // return charId
			client.CurrentChar = client.Account.Char[client.Account.CharSlot]

			rg := models.GetRegion(client.CurrentChar.Coordinates.X, client.CurrentChar.Coordinates.Y)
			rg.AddVisibleObject(client.CurrentChar)
			client.CurrentChar.CurrentRegion = rg
			g.addOnlineChar(client.CurrentChar)
			err = client.SimpleSend(client.Buffer.Bytes(), true)
			if err != nil {
				log.Println(err)
			}
			log.Println("Send CharSelected")
		case 208:
			if len(data) >= 2 {
				switch data[0] {
				case 1:
					serverpackets.NewExSendManorList(client)
					err := client.SimpleSend(client.Buffer.Bytes(), true)
					if err != nil {
						log.Println(err)
					}
					log.Println("Send ExSendManorList")
				case 54:
					client.Account = serverpackets.NewCharSelectionInfo(g.database, client)
					err := client.SimpleSend(client.Buffer.Bytes(), true)
					if err != nil {
						log.Println(err)
					}
					log.Println("Send NewCharSelectionInfo")
				}

			}

		case 193:
			serverpackets.NewObservationReturn(client.CurrentChar, client)
			err := client.SimpleSend(client.Buffer.Bytes(), true)
			if err != nil {
				log.Println(err)
			}
		case 108:
			serverpackets.NewShowMiniMap(client)
			err := client.SimpleSend(client.Buffer.Bytes(), true)
			if err != nil {
				log.Println(err)
			}
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

			pkg = serverpackets.NewItemList()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

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

			pkg = serverpackets.NewSkillList()
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
			var info PacketByte
			info.SetB(pkg)
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			Broad(g, client.CurrentChar, info)

			log.Println("Send NewMoveToLocation")
		case 73:
			//	say := clientpackets.NewSay(data)
			var info PacketByte
			//info.b = serverpackets.NewCreatureSay(say, client.CurrentChar)
			//err := client.Send(info.GetB(), true)
			//if err != nil {
			//	log.Println(err)
			//}
			info.b = serverpackets.NewCharInfo(client.CurrentChar)
			Broad(g, client.CurrentChar, info)
		case 89:
			clientpackets.NewValidationPosition(data, client.CurrentChar)
		default:
			log.Println("Not Found case with opcode: ", opcode)
		}
	}
}