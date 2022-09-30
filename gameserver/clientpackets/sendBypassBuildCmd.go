package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"os"
)

func SendBypassBuildCmd(client interfaces.CharacterI, data []byte) {
	var packet = packets.NewReader(data)
	command := packet.ReadString()
	_ = command
	data, err := os.ReadFile("./datapack/html/admin/main.html")
	if err != nil {
		return
	}
	d := utils.B2s(data)
	client.SendBuf(serverpackets.NpcHtmlMessage2(0, d, 0))
}
