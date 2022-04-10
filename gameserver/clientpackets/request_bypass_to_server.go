package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/htm"
	"l2gogameserver/packets"
	"log"
)

var bbslist = [8]string{
	"bypass _bbshome",
	"bypass _bbsgetfav",
	"bypass _bbsloc",
	"bypass _bbsclan",
	"bypass _bbsmemo",
	"bypass _maillist_0_1_0_",
	"bypass _bbsfriends",
	"bypass _bbsaddfav",
}

func BypassToServer(data []byte, client *models.Client) {
	//var bbsname = packets.NewReader(data).ReadString()
	//log.Println(bbsname)

	htmlDialog, err := htm.Open("./server/html/community/index.htm")
	if err != nil {
		log.Println(err)
		return
	}
	bufferDialog := packets.Get()
	defer packets.Put(bufferDialog)
	bufferDialog1 := packets.Get()
	defer packets.Put(bufferDialog1)
	bufferDialog2 := packets.Get()
	defer packets.Put(bufferDialog2)

	if len(htmlDialog) < 16250 {
		bufferDialog.WriteSlice(models.QWERT(htmlDialog, "101"))
		bufferDialog1.WriteSlice(models.QWERT("", "102"))
		bufferDialog2.WriteSlice(models.QWERT("", "103"))
	}
	buffer := packets.Get()

	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(bufferDialog.Bytes()))
	//	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(bufferDialog1.Bytes()))
	//	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(bufferDialog2.Bytes()))
	client.SSend(buffer.Bytes())
}
