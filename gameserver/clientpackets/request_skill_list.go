package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestSkillList(client interfaces.ReciverAndSender, data []byte) {
	pkg := serverpackets.SkillList(client)
	client.EncryptAndSend(pkg)
}
