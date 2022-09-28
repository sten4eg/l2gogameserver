package clientpackets

import (
	"fmt"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
)

func DropItem(client interfaces.ReciverAndSender, data []byte) {
	var read = packets.NewReader(data)
	objectId := read.ReadInt32()
	count := read.ReadInt64()
	x := read.ReadInt32()
	y := read.ReadInt32()
	z := read.ReadInt32()

	activeChar := client.GetCurrentChar()
	//item := drops.DropItemCharacter(client, objectId, count, x, y, z)
	item := activeChar.DropItem(objectId, count)

	items := []interfaces.MyItemInterface{item}
	msg := serverpackets.InventoryUpdate(items)
	client.EncryptAndSend(msg)

	pkg := dropItem(item, activeChar.GetObjectId(), x, y, z)
	err := client.SendBuf(pkg)
	if err != nil {
		fmt.Println("spok")
	}
	//activeChar.Send(pkg)

	//buffer := packets.Get()
	//defer packets.Put(buffer)

	//pkg := serverpackets.DropItem(client, objectId, count, x, y, z)
	//buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
	//
	//return buffer.Bytes(), item
	//return []byte{}, models.MyItem{}
}

func dropItem(item interfaces.MyItemInterface, charObjectId int32, x int32, y int32, z int32) *packets.Buffer {
	buf := packets.Get()

	buf.WriteSingleByte(0x16)
	buf.WriteD(charObjectId)
	buf.WriteD(item.GetObjectId())
	buf.WriteD(item.GetId())

	buf.WriteD(x)
	buf.WriteD(y)
	buf.WriteD(z)

	buf.WriteD(utils.BoolToInt32(item.IsStackable()))
	buf.WriteQ(item.GetCount())

	buf.WriteD(0x01)

	return buf
}
