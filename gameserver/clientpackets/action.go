package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func Action(data []byte, clientI interfaces.ReciverAndSender) *models.BackwardToLocation {
	client, ok := clientI.(*models.Client)
	if !ok {
		return nil
	}
	reAppeal := false // повторное обращение к объекту
	var packet = packets.NewReader(data)
	objectId := packet.ReadInt32() //Target
	originX := packet.ReadInt32()
	originY := packet.ReadInt32()
	originZ := packet.ReadInt32()
	actionId := packet.ReadSingleByte() // Action identifier : 0-Simple click, 1-Shift click
	if objectId == client.CurrentChar.Target {
		reAppeal = true
	} else {
		client.CurrentChar.Target = objectId
	}

	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.TargetSelected(client.CurrentChar.ObjectId, objectId, originX, originY, originZ)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	client.SSend(buffer.Bytes())

	return MoveToLocation(client, originX, originY, originZ)

	_, _ = actionId, reAppeal
	/*
		npc, npcx, npcy, npcz, err := models.GetNpcObject(objectId)
		if err != nil {
			log.Println(err)
		}

		//Прост тест вызова HTML при клике
		if actionId == 1 {
			NpcHtmlMessage := NpcHtmlMessage(client, npc.NpcId)
			client.SSend(NpcHtmlMessage)
		}
		//Если повторный клик по нпц
		if reAppeal {
			//npc, npcx, npcy, npcz, err := models.GetNpcObject(objectId)
			//if err != nil {
			//	log.Println(err)
			//}
			x, y, z := client.CurrentChar.GetXYZ()
			distance := models.CalculateDistance(npcx, npcy, npcz, x, y, z, false, false)
			_, _ = distance, npc

			//подбегаем
			if distance <= 60 {
				log.Println("Расстояние до NPC подходящее")
				if models.GetDialogNPC(npc.Type) == 0 {
					//НПЦ для разговора, открываем диалог
					//Пускай макс. дистанция разговора будет 60 поинтов
					//Пока откроем ID нпц
					NpcHtmlMessage := NpcHtmlMessage(client, npc.NpcId)
					client.SSend(NpcHtmlMessage)
				} else {
					//бьем нпц
					client.SSend(Attack(data, client))
				}
			} else {
				log.Println("Расстояние до NPC слишком больше, необходимо подбежать")
				return MoveToLocation(client, npcx, npcy, npcz)

			}

		}

	*/
	return nil
}
