package listeners

import (
	"fmt"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"log"
	"strconv"
)

func StartClientListener(client interfaces.ReciverAndSender) {
	go channelListener(client)
	go npcListener(client)
	go moveListener(client)
	go dropItemListener(client)
	go deleteObjectListener(client)
	go listenSkillQueue(client)

}

func channelListener(client interfaces.ReciverAndSender) {
	ch, ok := client.(*models.ClientCtx)
	if !ok {
		logger.Error.Panicln("ChannelListenerlogger.Error.Panicln")
	}
	for {
		select {
		case q := <-ch.CurrentChar.ChannelUpdateShadowItem:
			pkg := serverpackets.ItemUpdate(client, q.UpdateType, q.ObjId)
			client.EncryptAndSend(pkg)
			if q.UpdateType == models.UpdateTypeRemove {
				broadcast.BroadCastUserInfoInRadius(client, 2000)
			}
		case _ = <-ch.CurrentChar.EndChannel:
			return
		}
	}

}

func npcListener(client interfaces.ReciverAndSender) {
	ch, ok := client.(*models.ClientCtx)
	if !ok {
		logger.Error.Panicln("NpcListenerlogger.Error.Panicln")
	}
	for {
		select {
		case q := <-ch.CurrentChar.NpcInfo:
			buff := packets.Get()
			for i := range q {
				pkg := serverpackets.NpcInfo(q[i])
				buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
			}
			client.Send(buff.Bytes())
		case _ = <-ch.CurrentChar.EndChannel:
			return
		}
	}

}
func dropItemListener(client interfaces.ReciverAndSender) {
	ch, ok := client.(*models.ClientCtx)
	if !ok {
		logger.Error.Panicln("NpcListenerlogger.Error.Panicln")
	}
	for {
		select {
		case q := <-ch.CurrentChar.DropItemsInfo:
			buff := packets.Get()
			for i := range q {
				pkg := serverpackets.DropItem(q[i], 0)
				buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg.Bytes()))
			}
			client.Send(buff.Bytes())
		case _ = <-ch.CurrentChar.EndChannel:
			return
		}
	}

}
func moveListener(client interfaces.ReciverAndSender) {
	ch, ok := client.(*models.ClientCtx)
	if !ok {
		logger.Error.Panicln("NpcListenerlogger.Error.Panicln")
	}

	pkg := utils.GetPacketByte()
	defer pkg.Release()

	for {
		select {
		case to := <-ch.CurrentChar.CharInfoTo:
			pkg.SetData(serverpackets.CharInfo(ch.CurrentChar))
			for index := range to {
				strKey := strconv.Itoa(int(to[index]))
				char, ok := gameserver.OnlineCharacters.Load(strKey)
				if !ok {
					log.Println("Персонаж не найден")
					continue
				}
				char.EncryptAndSend(pkg.GetData())
			}
		case _ = <-ch.CurrentChar.EndChannel:
			return
		}
	}
}

func deleteObjectListener(client interfaces.ReciverAndSender) {
	ch, ok := client.(*models.ClientCtx)
	if !ok {
		logger.Error.Panicln("NpcListenerlogger.Error.Panicln")
	}

	pkg := utils.GetPacketByte()
	defer pkg.Release()

	for {
		select {
		case to := <-ch.CurrentChar.DeleteObjectTo:
			pkg.SetDataBuf(serverpackets.DeleteObject(ch.CurrentChar.GetObjectId()))
			for index := range to {

				strKey := strconv.Itoa(int(to[index]))
				char, ok := gameserver.OnlineCharacters.Load(strKey)
				if !ok {
					log.Println("Персонаж не найден")
					continue
				}
				char.EncryptAndSend(pkg.GetData())

			}
		case _ = <-ch.CurrentChar.EndChannel:
			return
		}
	}

}

func listenSkillQueue(client interfaces.ReciverAndSender) {
	ch, ok := client.(*models.ClientCtx)
	if !ok {
		logger.Error.Panicln("NpcListenerlogger.Error.Panicln")
	}

	for {
		select {
		case res := <-ch.CurrentChar.SkillQueue:
			fmt.Println("SKILL V QUEUE")
			fmt.Println(res)
		case _ = <-ch.CurrentChar.EndChannel:
			return
		}
	}
}
