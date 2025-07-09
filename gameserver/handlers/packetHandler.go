package handlers

import (
	"database/sql"
	"fmt"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/listeners"
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/packets"
)

type GameServerInterface interface {
	AddClient(string, interfaces.NewClientCtxInterface) bool
	AddWaitClient(string, uint32, uint32, uint32, uint32)
	RemoveClient(string)
	SendLogout(string)
	GetDbConn() *sql.DB
	CharOffline(ctxInterface interfaces.NewClientCtxInterface)
	AddOnlineChar(character interfaces.CharacterI)
	GetChar(string) (interfaces.CharacterI, bool)
	GetNetConnByCharObjectId(objectId int32) interfaces.CharacterI
	GetClientByLogin(login string) interfaces.NewClientCtxInterface
}

func Handler(client interfaces.NewClientCtxInterface, gs GameServerInterface) {
	//defer kickClient(client)
	for {
		opcode, data, err := client.Receive()
		if err != nil {
			fmt.Println(err)
			gs.CharOffline(client) //todo если чар офф то надо менять state
			return                 // todo  return ?
		}
		logger.Info.Println("Client->Server: #", opcode, packets.GetNamePacket(opcode))

		state := client.GetState()
		switch state {
		case clientStates.Connected: //TODO в этом стейте использовать ClientCtxInterface
			switch opcode {
			default:
				fmt.Printf("Неопознаный опкод {%x} при state Connection\n", opcode)
			case 0x0e:
				clientpackets.ProtocolVersion(client, data) // интерфейсный
			case 0x2b:
				clientpackets.AuthLogin(data, client, gs)
			}
		case clientStates.Authed: //TODO в этом стейте использовать ClientCtxInterface
			switch opcode {
			default:
				fmt.Printf("Неопознаный опкод %x при state Authed\n", opcode)
			case 0x00:
				clientpackets.Logout(client, state, gs)
			case 0x0c:
				clientpackets.CharacterCreate(client, data, gs.GetDbConn())
			case 0x0d:
				clientpackets.CharacterDelete(client, data, gs.GetDbConn())
			case 0x12:
				clientpackets.CharSelected(data, client, gs.GetDbConn())
				gs.AddOnlineChar(client.GetAccount().GetCurrentChar()) //todo проверить зачем еще одна мапа с чарами онлайн, есть мапа с клиентами
			case 0x13:
				clientpackets.RequestNewCharacter(client)
			case 0xd0:
				if len(data) >= 2 {
					switch data[0] {
					default:
						fmt.Printf("Неопознаный второй опкод %x при state Authed\n", data[0])
					case 0x36:
						clientpackets.RequestGoToLobby(client, gs.GetDbConn())
					}
				}
			}
		case clientStates.Joining: // TODO в этом стейте можно использовать CharacterInterface
			switch opcode {
			default:
				fmt.Printf("Неопознаный опкод %x при state Joining\n", opcode)
			case 0x11:
				clientpackets.RequestEnterWorld(client, data, gs.GetDbConn())
				//broadcast.BroadCastUserInfoInRadius(client, 2000)
				//рассылка при входе в игру происходит в setWorldRegion // broadcast.SendCharInfoAboutCharactersInRadius(client, 2000)
				go listeners.StartClientListener(client, gs) //todo  надо зпускать не отсюда
			case 0xd0:
				if len(data) >= 2 {
					switch data[0] {
					default:
						fmt.Printf("Неопознаный второй опкод %x при state Joining\n", data[0])
					case 0x01:
						clientpackets.RequestManorList(client, data)
					}
				}
			}
		case clientStates.InGame:
			character := client.GetAccount().GetCurrentChar()

			switch opcode {
			default:
				fmt.Printf("Неопознаный опкод {%x} при state InGame\n", opcode)
			case 0x00:
				clientpackets.Logout(client, state, gs)
			case 0x01:
				clientpackets.Attack(data, client)
			case 0x1a: //Запрос другому персонажу на желание торговать
				clientpackets.TradeRequest(data, client, gs)
			case 0x55: //AnswerTradeRequest (если пользователь отвечает Да/Нет на предложение торговли)
				clientpackets.AnswerTradeRequest(data, client, gs)
			case 0x1b: //AddTradeItem
				clientpackets.AddTradeItem(data, client, gs)
			case 0x1c: //tradeDone
				clientpackets.TradeDone(data, client, gs.GetDbConn(), gs)
			case 0x17:
				clientpackets.DropItem(client, data, gs.GetDbConn())
			//client.Send(pkg)
			//
			//pkgInventoryUpdate := clientpackets.InventoryUpdate(client, &item, models.UpdateTypeRemove)
			//client.Send(pkgInventoryUpdate)
			case 0x14:
				clientpackets.RequestItemList(data, client)
			case 0x23: //todo пересмотреть
				clientpackets.BypassToServer(data, client, gs.GetDbConn())
			case 0x19:
				clientpackets.UseItem(character, data, gs.GetDbConn(), gs)
			case 0x31:
				clientpackets.SetPrivateStoreListSell(client, data, gs)
			case 0x39:
				clientpackets.RequestMagicSkillUse(data, client)
			case 0x3d:
				clientpackets.RequestShortCutReg(data, client, gs.GetDbConn())
			case 0x3f:
				clientpackets.RequestShortCutDel(data, client, gs.GetDbConn())
			case 0x57:
				clientpackets.RequestRestart(client, gs, gs.GetDbConn())
				//gameserver.CharOffline(client)
			case 0x60:
				clientpackets.DestroyItem(data, client, gs.GetDbConn())
			case 0xc1:
				clientpackets.RequestObserverEnd(client, data)
			case 0x5c:
				clientpackets.FinishRotating(client, data)
			case 0x6c:
				clientpackets.RequestShowMiniMap(client, data)
			case 0xa6: //TODO На java сборках пакет deprecated и не реализован
				clientpackets.RequestSkillCoolTime(client, data)
			case 0x0f:
				pkg := clientpackets.MoveBackwardToLocation(client, data)
				broadcast.Checkaem(client, pkg)
			case 0x42:
				clientpackets.RequestJoinParty(client, data, gs)
			case 0x43:
				clientpackets.RequestAnswerJoinParty(client, data, gs)
			case 0x44:
				clientpackets.RequestWithDrawalParty(client, gs)
			case 0x49:
				say := clientpackets.Say(client, data)
				broadcast.BroadCastChat(client, say, gs)

			case 0x59:
				clientpackets.ValidationPosition(data, client.GetAccount().GetCurrentChar())
				//broadcast.Checkaem(client, pkg)
			case 0x50:
				clientpackets.RequestSkillList(client, data)
			case 0x1f:
				pkg := clientpackets.Action(data, client, broadcast.Checkaem, gs.GetDbConn())
				if pkg != nil {
					broadcast.Checkaem(client, *pkg)
				}
			case 0x48:
				clientpackets.RequestTargetCancel(data, client)
			case 0xcd:
				clientpackets.RequestMakeMacro(client, data)
			case 0xce:
				clientpackets.RequestDeleteMacro(client, data, gs.GetDbConn())
			case 0x56:
				clientpackets.RequestActionUse(client, data, gs)
			case 0xd0:
				switch data[0] {
				default:
					fmt.Printf("Неопознаный второй опкод {%x} при state InGame, первый опкод {%x}\n", data[0], opcode)
				case 0x24:
					clientpackets.RequestSaveInventoryOrder(client, data)
				case 0x0d:
					clientpackets.RequestAutoSoulShot(data, client)
				case 0x7a:
					clientpackets.AnswerCoupleAction(client, data)

				}
			case 0x74:
				clientpackets.SendBypassBuildCmd(character, data, gs)
			case 0x83:
				clientpackets.RequestPrivateStoreBuy(client, data, gs.GetDbConn())
			case 0x96:
				clientpackets.RequestPrivateStoreQuitSell(client)
			case 0x97:
				clientpackets.SetPrivateStoreMsgSell(client, data)
			case 0x9a:
				clientpackets.SetPrivateStoreListBuy(client, data)
			case 0x9c:
				clientpackets.RequestPrivateStoreQuitBuy(client)
			case 0x9d:
				clientpackets.SetPrivateStoreMsgBuy(client, data)
			case 0x9f:
				clientpackets.RequestPrivateStoreSell(client, data, gs.GetDbConn())
			}
		}

		//todo куда кинуть оставшийся
		switch opcode {
		case 114:
			clientpackets.MoveToPawn(client, data)
		}

	}
}
