package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestSkillList(client interfaces.NewClientCtxInterface, data []byte) {
	pkg := serverpackets.SkillList(client.GetAccount().GetCurrentChar())
	client.EncryptAndSend(pkg)
}
