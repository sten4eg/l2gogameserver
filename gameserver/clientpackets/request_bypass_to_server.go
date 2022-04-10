package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/htm"
	"l2gogameserver/packets"
	"log"
	"math"
	"strconv"
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
	numberOfSplits := int(math.Ceil(float64(len(htmlDialog)) / 8180))
	for i := 1; i < numberOfSplits+1; i++ {
		bufferDialog := packets.Get()
		bufferDialog.WriteSingleByte(0x7b)
		bufferDialog.WriteSingleByte(0x01)
		for _, s := range bbslist {
			bufferDialog.WriteS(s)
		}
		bufferDialog.WriteS("10" + strconv.Itoa(i) + "\u0008")
		if i == 1 {
			bufferDialog.WriteS(htmlDialog + "\u0008")
		} else if i == 2 {
			bufferDialog.WriteS(htmlDialog[8180:] + "\u0008")
		} else if i == 3 {
			bufferDialog.WriteS(htmlDialog[8180*2:] + "\u0008")
		} else {
			log.Println("Html is too long!")
			return
		}
		buffer := packets.Get()
		buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(bufferDialog.Bytes()))
		client.SSend(buffer.Bytes())
	}

}
