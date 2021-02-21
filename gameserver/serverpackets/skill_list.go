package serverpackets

import (
	"github.com/jackc/pgx"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func NewSkillList(charId int32, conn *pgx.Conn) []byte {
	buffer := new(packets.Buffer)

	skills := models.GetMySkills(charId, conn)

	buffer.WriteSingleByte(0x5F)

	buffer.WriteD(int32(len(skills))) // skill size

	for _, v := range skills {
		buffer.WriteD(0)           // passiv ?
		buffer.WriteD(1)           // level
		buffer.WriteD(int32(v.ID)) // id
		buffer.WriteD(0)           // disable?
		buffer.WriteD(0)           // enchant ?
	}

	return buffer.Bytes()
}
