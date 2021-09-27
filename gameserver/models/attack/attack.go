package attack

import (
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/models"
	"log"
)

//Действие при /attack
func IsAttack(data []byte, client *models.Client) (models.Npc, int32, int32, int32, int32, byte, bool, error) {
	pkg, targetObjectId, actionId, reAppeal := clientpackets.Action(data, client)
	client.SSend(pkg)
	npc, npcx, npcy, npcz, err := models.GetNpcObject(targetObjectId)
	if err != nil {
		//Тут нужно проверить, является ли объект предметом...
		//Объект может являться NPC
		log.Println(err)
		return models.Npc{}, 0, 0, 0, 0, 0, false, nil
	}
	return npc, npcx, npcy, npcz, targetObjectId, actionId, reAppeal, nil
}

//  Повторная нажатие атаки
func ReAppeal(npc models.Npc, npcx, npcy, npcz, targetObjectId int32, data []byte, client *models.Client) {
	//client.SSend(clientpackets.ChangeMoveType(client, targetObjectId))
	if models.GetDialogNPC(npc.Type) == 0 { //Это диалоговый NPC
		OpenDialogHTML(client, npc)
	} else { //NPC для атаки
		AttackNPC(client, data)
	}
}

//Дейтвие атаки
func AttackNPC(client *models.Client, data []byte) {
	client.SSend(clientpackets.Attack(data, client))
}

//Открытие HTML
func OpenDialogHTML(client *models.Client, npc models.Npc) {
	NpcHtmlMessage := clientpackets.NpcHtmlMessage(client, npc)
	client.SSend(NpcHtmlMessage)
}

// DistanceToNpc Дистанция до NPC
func DistanceToNpc(client *models.Client, npcx, npcy, npcz int32) float64 {
	x, y, z := client.CurrentChar.GetXYZ()
	distance := models.CalculateDistance(npcx, npcy, npcz, x, y, z, false, false)
	return distance
}
