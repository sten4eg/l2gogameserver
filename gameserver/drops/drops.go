package drops

import "l2gogameserver/gameserver/models"

//Дроп предмета персонажа
func DropItemCharacter(client *models.Client, objectId int32, count int64, x, y, z int32) models.MyItem {
	_, item := models.RemoveItemCharacter(client.CurrentChar, objectId, count)
	return item
}

//Дроп предметов из моба
func DropItemMobs() {

}
