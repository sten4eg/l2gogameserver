package serverpackets

import (
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/models"
)

func NewMoveToLocation(location *clientpackets.Location, client *models.Client, character *Character) {

	client.Buffer.WriteH(0) //reserve for lenght
	client.Buffer.WriteSingleByte(0x2f)

	client.Buffer.WriteD(character.CharId)

	client.Buffer.WriteD(location.TargetX)
	client.Buffer.WriteD(location.TargetY)
	client.Buffer.WriteD(location.TargetZ)

	client.Buffer.WriteD(location.OriginX)
	client.Buffer.WriteD(location.OriginY)
	client.Buffer.WriteD(location.OriginZ)

}
