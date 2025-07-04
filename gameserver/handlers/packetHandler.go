package handlers

import (
	"fmt"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/listeners"
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/packets"
)

type GameServerInterface interface {
	AddClient(string, interfaces.ClientInterface) bool
	AddWaitClient(string, interfaces.ClientInterface)
	RemoveClient(string)
	SendLogout(string)
}

func Handler(client interfaces.ClientInterface, gs GameServerInterface) {
	//defer kickClient(client)
	for {
		opcode, data, err := client.Receive()
		if err != nil {
			fmt.Println(err)
			gameserver.CharOffline(client) //todo если чар офф то надо менять state
			return                         // todo  return ?
		}
		logger.Info.Println("Client->Server: #", opcode, packets.GetNamePacket(opcode))

		state := client.GetState()
		switch state {
		case clientStates.Connected:
			switch opcode {
			default:
				fmt.Printf("Неопознаный опкод {%x} при state Connection\n", opcode)
			case 0x0e:
				clientpackets.ProtocolVersion(client, data)
			case 0x2b:
				clientpackets.AuthLogin(data, client, gs)
			}
		case clientStates.Authed:
			switch opcode {
			default:
				fmt.Printf("Неопознаный опкод %x при state Authed\n", opcode)
			case 0x00:
				clientpackets.Logout(client, state, gs)
			case 0x0c:
				clientpackets.CharacterCreate(client, data)
			case 0x0d:
				clientpackets.CharacterDelete(client, data)
			case 0x12:
				clientpackets.CharSelected(data, client)
				gameserver.AddOnlineChar(client.GetCurrentChar()) //todo проверить зачем еще одна мапа с чарами онлайн, есть мапа с клиентами
			case 0x13:
				clientpackets.RequestNewCharacter(client)
			case 0xd0:
				if len(data) >= 2 {
					switch data[0] {
					default:
						fmt.Printf("Неопознаный второй опкод %x при state Authed\n", data[0])
					case 0x36:
						clientpackets.RequestGoToLobby(client)
					}
				}
			}
		case clientStates.Joining:
			switch opcode {
			default:
				fmt.Printf("Неопознаный опкод %x при state Joining\n", opcode)
			case 0x11:
				clientpackets.RequestEnterWorld(client, data)
				broadcast.BroadCastUserInfoInRadius(client, 2000)
				//рассылка при входе в игру происходит в setWorldRegion // broadcast.SendCharInfoAboutCharactersInRadius(client, 2000)
				go listeners.StartClientListener(client) //todo  надо зпускать не отсюда
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
			character := client.GetCurrentChar()

			switch opcode {
			default:
				fmt.Printf("Неопознаный опкод {%x} при state InGame\n", opcode)
			case 0x00:
				clientpackets.Logout(character, state, gs)
			case 0x01:
				clientpackets.Attack(data, character)
			case 0x1a: //Запрос другому персонажу на желание торговать
				clientpackets.TradeRequest(data, client)
			case 0x55: //AnswerTradeRequest (если пользователь отвечает Да/Нет на предложение торговли)
				clientpackets.AnswerTradeRequest(data, client)
			case 0x1b: //AddTradeItem
				clientpackets.AddTradeItem(data, client)
			case 0x1c: //tradeDone
				clientpackets.TradeDone(data, client)
			case 0x17:
				clientpackets.DropItem(client, data)
			//client.Send(pkg)
			//
			//pkgInventoryUpdate := clientpackets.InventoryUpdate(client, &item, models.UpdateTypeRemove)
			//client.Send(pkgInventoryUpdate)
			case 0x14:
				clientpackets.RequestItemList(data, character)
			case 0x23: //todo пересмотреть
				clientpackets.BypassToServer(data, client)
			case 0x19:
				clientpackets.UseItem(character, data)
			case 0x31:
				clientpackets.SetPrivateStoreListSell(client, data)
			case 0x39:
				clientpackets.RequestMagicSkillUse(data, client)
			case 0x3d:
				clientpackets.RequestShortCutReg(data, client)
			case 0x3f:
				clientpackets.RequestShortCutDel(data, client)
			case 0x57:
				clientpackets.RequestRestart(client)
				//gameserver.CharOffline(client)
			case 0x60:
				clientpackets.DestroyItem(data, client)
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
				clientpackets.RequestJoinParty(client, data)
			case 0x43:
				clientpackets.RequestAnswerJoinParty(client, data)
			case 0x44:
				clientpackets.RequestWithDrawalParty(client)
			case 0x49:
				say := clientpackets.Say(client, data)
				broadcast.BroadCastChat(client, say)

			case 0x59:
				clientpackets.ValidationPosition(data, client.GetCurrentChar())
				//broadcast.Checkaem(client, pkg)
			case 0x50:
				clientpackets.RequestSkillList(client, data)
			case 0x1f:
				pkg := clientpackets.Action(data, client, broadcast.Checkaem)
				if pkg != nil {
					broadcast.Checkaem(client, *pkg)
				}
			case 0x48:
				clientpackets.RequestTargetCancel(data, client)
			case 0xcd:
				clientpackets.RequestMakeMacro(client, data)
			case 0xce:
				clientpackets.RequestDeleteMacro(client, data)
			case 0x56:
				clientpackets.RequestActionUse(client, data)
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
				clientpackets.SendBypassBuildCmd(character, data)
			case 0x83:
				clientpackets.RequestPrivateStoreBuy(client, data)
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
				clientpackets.RequestPrivateStoreSell(client, data)
			}
		}

		//todo куда кинуть оставшийся
		switch opcode {
		case 114:
			clientpackets.MoveToPawn(client, data)
		}

	}
}
