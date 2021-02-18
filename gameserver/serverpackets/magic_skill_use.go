package serverpackets

import "l2gogameserver/gameserver/models"

func NewMagicSkillUse(client *models.Client) {
	client.Buffer.WriteH(0) // reserve

	client.Buffer.WriteSingleByte(0x48)
	client.Buffer.WriteD(client.CurrentChar.CharId) // activeChar id
	client.Buffer.WriteD(client.CurrentChar.CharId) // targetChar id
	client.Buffer.WriteD(1216)                      // skillId
	client.Buffer.WriteD(1)                         // skillLevel
	client.Buffer.WriteD(5000)                      // hitTime
	client.Buffer.WriteD(3000)                      // reuseDelay

	x, y, z := client.CurrentChar.GetXYZ()
	client.Buffer.WriteD(x)
	client.Buffer.WriteD(y)
	client.Buffer.WriteD(z)

	client.Buffer.WriteH(0) //size???
	// for  by size ???

	client.Buffer.WriteH(0) // _groundLocations.size()
	// for by _groundLocations.size()

	//location target
	client.Buffer.WriteD(x)
	client.Buffer.WriteD(y)
	client.Buffer.WriteD(z)
	client.SimpleSend(client.Buffer.Bytes(), true)

}
