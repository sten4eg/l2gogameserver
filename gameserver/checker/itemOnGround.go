package checker

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"time"
)

func DropItemChecker(region interfaces.WorldRegioner) {
	for {
		ids := region.DropItemChecker()
		for _, id := range ids {
			buffer := serverpackets.DeleteObject(id)
			pb := utils.GetPacketByte()
			pb.SetData(buffer.Bytes())
			for _, r := range region.GetNeighbors() {
				for _, character := range r.GetCharsInRegion() {
					logger.Info.Println(character.GetName())
					character.EncryptAndSend(pb.GetData())
				}
			}
			pb.Release()
			packets.Put(buffer)
		}
		time.Sleep(time.Second * 5)
	}
}
