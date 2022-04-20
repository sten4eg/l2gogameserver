package drops

import (
	"l2gogameserver/gameserver/models"
)

//Дроп предмета персонажа
func DropItemCharacter(client interfaces.ReciverAndSender, objectId int32, count int64, x, y, z int32) {
	//TODO наверно для выкидывания на пол надо придумать другую функцию
	// models.RemoveItemCharacter(client.CurrentChar, objectId, count)

}

//Дроп предметов из моба
func DropItemMobs() {

}
