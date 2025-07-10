package listeners

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"log"
	"strconv"
)

type GsInterf interface {
	GetChar(string) (interfaces.CharacterI, bool)
	GetNetConnByCharObjectId(objectId int32) interfaces.CharacterI
	GetClientByLogin(login string) interfaces.NewClientCtxInterface
}

func StartClientListener(client interfaces.NewClientCtxInterface, g GsInterf) {
	go startListener(client, g)
}

func startListener(client interfaces.NewClientCtxInterface, g GsInterf) {
	currChar := client.GetAccount().GetCurrentChar()
	for {
		select {
		case charsAround := <-currChar.GetChannelCharInfoTo():

			charInfoAboutMe := utils.GetPacketByte()
			extraUserInfoAboutMe := utils.GetPacketByte()

			charInfoAboutMe.SetData(serverpackets.CharInfo(currChar))
			extraUserInfoAboutMe.SetData(serverpackets.ExBrExtraUserInfo(currChar))

			for index := range charsAround {
				strKey := strconv.Itoa(int(charsAround[index]))
				charAround, ok := g.GetChar(strKey)
				if !ok {
					log.Println("Персонаж не найден")
					continue
				}
				pkg2 := utils.GetPacketByte()
				exUi2 := utils.GetPacketByte()

				pkg2.SetData(serverpackets.CharInfo(charAround))
				exUi2.SetData(serverpackets.ExBrExtraUserInfo(charAround))

				client.EncryptAndSend(pkg2.GetData())
				client.EncryptAndSend(exUi2.GetData())

				pkg2.Release()
				exUi2.Release()

				charAround.EncryptAndSend(charInfoAboutMe.GetData())
				charAround.EncryptAndSend(extraUserInfoAboutMe.GetData())
			}
			charInfoAboutMe.Release()
			extraUserInfoAboutMe.Release()

		case npcsList := <-currChar.GetChannelNpcInfo():
			buff := packets.Get()
			for npcInfo := range npcsList {
				pkg := serverpackets.NpcInfo(npcsList[npcInfo])
				buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
			}
			client.Send(buff.Bytes())
		case shadowItemInfo := <-currChar.GetChannelUpdateShadowItem():
			pkg := serverpackets.ItemUpdate(client, shadowItemInfo.UpdateType, shadowItemInfo.ObjId)
			client.EncryptAndSend(pkg)
			if shadowItemInfo.UpdateType == models.UpdateTypeRemove {
				broadcast.BroadCastUserInfoInRadius(client, 2000, g)
			}
		case droppedItems := <-currChar.GetChannelDropItemsInfo():
			buff := packets.Get()
			for droppedItem := range droppedItems {
				pkg := serverpackets.DropItem(droppedItems[droppedItem], 0)
				buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg.Bytes()))
			}
			client.Send(buff.Bytes())
		case deleteObjects := <-currChar.GetChannelDeleteObjectTo():
			pkg := utils.GetPacketByte()
			pkg.SetDataBuf(serverpackets.DeleteObject(currChar.GetObjectId()))
			for index := range deleteObjects {

				strKey := strconv.Itoa(int(deleteObjects[index]))
				char, ok := g.GetChar(strKey)
				if !ok {
					log.Println("Персонаж не найден")
					continue
				}
				char.EncryptAndSend(pkg.GetData())

			}
			pkg.Release()
		//case res := <-currChar.SkillQueue:
		//	fmt.Println("SKILL V QUEUE")
		//	fmt.Println(res)
		case _ = <-currChar.GetChannelEndChannel():
			return
		}
	}
}
