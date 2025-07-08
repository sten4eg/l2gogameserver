package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestSkillCoolTime(client interfaces.NewClientCtxInterface, data []byte) {
	pkg := serverpackets.SkillCoolTime()
	client.EncryptAndSend(pkg)
}
