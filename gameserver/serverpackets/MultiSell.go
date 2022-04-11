package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/multisell"
	"l2gogameserver/packets"
)

//Отправка пакета на открытие мультиселла
func MultisellShow(client *models.Client, item []multisell.Item) {
	buffer := packets.Get()
	defer packets.Put(buffer)
	pkg := MultiSell(client, item)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
	client.SSend(buffer.Bytes())
}

//Отправка пакета
func MultiSell(client *models.Client, itemsMultisell []multisell.Item) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0xD0)
	buffer.WriteH(100) // listid
	buffer.WriteH(1)   // page
	buffer.WriteH(1)   // finished
	buffer.WriteH(40)  // size of pages
	buffer.WriteH(9)   // list length (предметов в продаже)

	for _, itemStack := range itemsMultisell {
		buffer.WriteD(int32(itemStack.ID))
		buffer.WriteSingleByte(0) // stackable?
		buffer.WriteH(0)          //
		buffer.WriteD(0)          //
		buffer.WriteD(0)          //

		buffer.WriteH(0)
		buffer.WriteH(0) // T1
		buffer.WriteH(0) // T1
		buffer.WriteH(0) // T1
		buffer.WriteH(0) // T1
		buffer.WriteH(0) // T1
		buffer.WriteH(0) // T1
		buffer.WriteH(0) // T1

		buffer.WriteH(40) // T1
		buffer.WriteH(40) // T1

		for _, getProduct := range itemStack.Production {
			buffer.WriteD(int32(getProduct.Id))
			buffer.WriteD(0) //getBodyPart
			buffer.WriteH(0) //getType2ForPackets

			buffer.WriteQ(int64(getProduct.Count)) //getItemCount

			buffer.WriteH(0) // attack element
			buffer.WriteH(0) // element power
			buffer.WriteH(0) // fire
			buffer.WriteH(0) // water
			buffer.WriteH(0) // wind
			buffer.WriteH(0) // earth
			buffer.WriteH(0) // holy
			buffer.WriteH(0) // dark

		}

		for _, getIngredient := range itemStack.Ingredient {

			buffer.WriteD(int32(getIngredient.Id))
			buffer.WriteH(0)
			buffer.WriteQ(int64(getIngredient.Count))

			buffer.WriteH(int16(getIngredient.Enchant)) // enchant level
			buffer.WriteD(0)                            // augment id
			buffer.WriteD(0)                            // mana
			buffer.WriteH(0)                            // attack element
			buffer.WriteH(0)                            // element power
			buffer.WriteH(0)                            // fire
			buffer.WriteH(0)                            // water
			buffer.WriteH(0)                            // wind
			buffer.WriteH(0)                            // earth
			buffer.WriteH(0)                            // holy
			buffer.WriteH(0)                            // dark

		}
	}
	return buffer.Bytes()
}
