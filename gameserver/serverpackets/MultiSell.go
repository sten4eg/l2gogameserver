package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/multisell"
	"l2gogameserver/packets"
)

const pageSize = 40

//Отправка пакета на открытие мультиселла
func MultisellShow(client *models.Client, msdata multisell.MultiList) {
	buffer := packets.Get()
	defer packets.Put(buffer)
	pkg := MultiSell(client, msdata)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
	client.SSend(buffer.Bytes())
}

//Отправка пакета
func MultiSell(client *models.Client, msdata multisell.MultiList) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0xD0)
	buffer.WriteD(int32(305986001))        // msdata.ID list id
	buffer.WriteD(1)                       // page started from 1
	buffer.WriteD(1)                       // finished
	buffer.WriteD(pageSize)                // size of pages
	buffer.WriteD(int32(len(msdata.Item))) // list length
	for i, items := range msdata.Item {
		buffer.WriteD(int32((i + 1) * 100000))
		buffer.WriteH(0)  //stack
		buffer.WriteH(0)  // C6
		buffer.WriteD(0)  // C6
		buffer.WriteD(0)  // T1
		buffer.WriteH(-2) // T1
		buffer.WriteH(0)  // T1
		buffer.WriteH(0)  // T1
		buffer.WriteH(0)  // T1
		buffer.WriteH(0)  // T1
		buffer.WriteH(0)  // T1
		buffer.WriteH(0)  // T1
		buffer.WriteH(0)  // T1
		buffer.WriteH(int16(len(items.Production)))
		buffer.WriteH(int16(len(items.Ingredient)))

		for _, item := range items.Production {
			buffer.WriteD(int32(item.Id))
			buffer.WriteD(0)
			buffer.WriteH(-1)
			buffer.WriteQ(int64(item.Count))
			buffer.WriteH(int16(item.Enchant)) // enchant level
			buffer.WriteD(0)                   // augment id
			buffer.WriteD(0)                   // mana
			buffer.WriteH(0)                   // attack element
			buffer.WriteH(0)                   // element power
			buffer.WriteH(0)                   // fire
			buffer.WriteH(0)                   // water
			buffer.WriteH(0)                   // wind
			buffer.WriteH(0)                   // earth
			buffer.WriteH(0)                   // holy
			buffer.WriteH(0)                   // dark

		}
		for _, item := range items.Ingredient {
			buffer.WriteD(int32(item.Id))
			buffer.WriteH(-1)
			buffer.WriteQ(int64(item.Count))
			buffer.WriteH(int16(item.Enchant)) // enchant level
			buffer.WriteD(0)                   // augment id
			buffer.WriteD(0)                   // mana
			buffer.WriteH(0)                   // attack element
			buffer.WriteH(0)                   // element power
			buffer.WriteH(0)                   // fire
			buffer.WriteH(0)                   // water
			buffer.WriteH(0)                   // wind
			buffer.WriteH(0)                   // earth
			buffer.WriteH(0)                   // holy
			buffer.WriteH(0)                   // dark
		}
	}

	return buffer.Bytes()
}
