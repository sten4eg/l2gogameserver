package serverpackets

import (
	"io/ioutil"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
	"log"
	"strconv"
)

func NpcHtmlMessage(client *models.Client, npc models.Npc) []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x19)

	buffer.WriteD(33)
	buffer.WriteS(openHTML(npc))
	buffer.WriteD(0)

	return buffer.Bytes()
}

func openHTML(npc models.Npc) string {
	npcids := strconv.Itoa(int(npc.NpcId))
	content, err := ioutil.ReadFile("./data/dialogs/" + npc.Type + "/" + npcids + ".htm")

	if err != nil {
		log.Fatal(err)
	}

	return string(content)

}

func openHTML(npc models.Npc) string {
	npcids := strconv.Itoa(int(npc.NpcId))
	content, err := ioutil.ReadFile("./data/dialogs/" + npc.Type + "/" + npcids + ".htm")

	if err != nil {
		log.Fatal(err)
	}

	return string(content)

}
