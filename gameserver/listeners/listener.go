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
	go channelListener(client, g)
	go npcListener(client)
	go moveListener(client, g)
	go dropItemListener(client)
	go deleteObjectListener(client, g)
	go listenSkillQueue(client)

}

func channelListener(client interfaces.NewClientCtxInterface, g GsInterf) {
	currChar := client.GetAccount().GetCurrentChar()
	for {
		select {
		case q := <-currChar.GetChannelUpdateShadowItem():
			pkg := serverpackets.ItemUpdate(client, q.UpdateType, q.ObjId)
			client.EncryptAndSend(pkg)
			if q.UpdateType == models.UpdateTypeRemove {
				broadcast.BroadCastUserInfoInRadius(client, 2000, g)
			}
		case _ = <-currChar.GetChannelEndChannel():
			return
		}
	}

}

func npcListener(client interfaces.NewClientCtxInterface) {
	//currChar := client.GetAccount().GetCurrentChar()
	//for {
	//	select {
	//	case q := <-currChar.GetChannelNpcInfo():
	//		buff := packets.Get()
	//		for i := range q {
	//			pkg := serverpackets.NpcInfo(q[i])
	//			buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
	//		}
	//		client.Send(buff.Bytes())
	//
	//	case _ = <-currChar.GetChannelEndChannel():
	//		return
	//	}
	//}

}
func dropItemListener(client interfaces.NewClientCtxInterface) {
	currChar := client.GetAccount().GetCurrentChar()

	for {
		select {
		case q := <-currChar.GetChannelDropItemsInfo():
			buff := packets.Get()
			for i := range q {
				pkg := serverpackets.DropItem(q[i], 0)
				buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg.Bytes()))
			}
			client.Send(buff.Bytes())
		case _ = <-currChar.GetChannelEndChannel():
			return
		}
	}

}
func moveListener(client interfaces.NewClientCtxInterface, g GsInterf) {

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

		case q := <-currChar.GetChannelNpcInfo():
			buff := packets.Get()
			for i := range q {
				pkg := serverpackets.NpcInfo(q[i])
				buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
			}
			client.Send(buff.Bytes())
		case _ = <-currChar.GetChannelEndChannel():
			return
		}
	}
}

func deleteObjectListener(client interfaces.NewClientCtxInterface, g GsInterf) {

	pkg := utils.GetPacketByte()
	defer pkg.Release()
	currChar := client.GetAccount().GetCurrentChar()
	for {
		select {
		case to := <-currChar.GetChannelDeleteObjectTo():
			pkg.SetDataBuf(serverpackets.DeleteObject(currChar.GetObjectId()))
			for index := range to {

				strKey := strconv.Itoa(int(to[index]))
				char, ok := g.GetChar(strKey)
				if !ok {
					log.Println("Персонаж не найден")
					continue
				}
				char.EncryptAndSend(pkg.GetData())

			}
		case _ = <-currChar.GetChannelEndChannel():
			return
		}
	}

}

func listenSkillQueue(client interfaces.NewClientCtxInterface) {

	//for {
	//	select {
	//	case res := <-ch.CurrentChar.SkillQueue:
	//		fmt.Println("SKILL V QUEUE")
	//		fmt.Println(res)
	//	case _ = <-ch.CurrentChar.EndChannel:
	//		return
	//	}
	//}
}
