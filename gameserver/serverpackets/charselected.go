package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func CharSelected(user *models.Character, client *models.Client) {

	x, y, z := user.GetXYZ()
	client.Buffer.WriteSingleByte(0x0b) // 1

	client.Buffer.WriteS(user.CharName)    // 11
	client.Buffer.WriteD(user.CharId)      // objId 15
	client.Buffer.WriteS(user.Title)       //title // 21 2 нуля
	client.Buffer.WriteD(0)                //TODO sessionId //25
	client.Buffer.WriteD(user.ClanId)      //clanId // 29
	client.Buffer.WriteD(0)                // ? //33
	client.Buffer.WriteD(user.Sex)         //sex// 37
	client.Buffer.WriteD(int32(user.Race)) //race 41
	client.Buffer.WriteD(user.ClassId)     //classId 45
	client.Buffer.WriteD(0x1)              // ? 49
	client.Buffer.WriteD(x)                //x 53
	client.Buffer.WriteD(y)                //y 57
	client.Buffer.WriteD(z)                //z 61

	client.Buffer.WriteF(float64(user.CurHp)) //currentHP 69
	client.Buffer.WriteF(float64(user.CurMp)) //currentMP 77

	client.Buffer.WriteD(user.Sp)         // SP 81
	client.Buffer.WriteQ(int64(user.Exp)) // EXP 89
	client.Buffer.WriteD(user.Level)      // level 93
	client.Buffer.WriteD(user.Karma)      // karma 97
	client.Buffer.WriteD(user.PkKills)    // pk 101

	client.Buffer.WriteD(21) //int 105
	client.Buffer.WriteD(40) //str 109
	client.Buffer.WriteD(43) //con 113
	client.Buffer.WriteD(25) //men 117
	client.Buffer.WriteD(30) //dex 121
	client.Buffer.WriteD(11) //wit 125

	client.Buffer.WriteD(user.OnlineTime) //inGameTime 129
	client.Buffer.WriteD(0)               // ?? 133

	client.Buffer.WriteD(user.ClassId) // 137 classId

	client.Buffer.WriteD(0) // 141
	client.Buffer.WriteD(0) // 145
	client.Buffer.WriteD(0) // 149
	client.Buffer.WriteD(0) // 153

	m := make([]byte, 64)
	client.Buffer.WriteSlice(m) //217

	client.Buffer.WriteD(0) // 221

	client.SaveAndCryptDataInBufferToSend(true)
	client.CurrentChar = client.Account.Char[client.Account.CharSlot]

	// загрузка всех данных выбранного чара
	client.CurrentChar.Load()
}
