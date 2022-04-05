package drops

import "l2gogameserver/gameserver/models"

//Дроп предмета персонажа
func DropItemCharacter(client *models.Client, objectId int32, count int64, x, y, z int32) {
	models.RemoveItemCharacter(client.CurrentChar, objectId, count)
}

//Дроп предметов из моба
func DropItemMobs() {

}
