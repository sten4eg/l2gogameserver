package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewCharSelected(user *Character, client *models.Client) {

	client.Buffer.WriteH(0) //reserve

	client.Buffer.WriteSingleByte(0x0b) // 1

	client.Buffer.WriteS(user.CharName)     // 11
	client.Buffer.WriteD(user.CharId)       // objId 15
	client.Buffer.WriteS(user.Title.String) //title // 21 2 нуля
	client.Buffer.WriteD(0)                 //TODO sessionId //25
	client.Buffer.WriteD(user.ClanId)       //clanId // 29
	client.Buffer.WriteD(0)                 // ? //33
	client.Buffer.WriteD(user.Sex)          //sex// 37
	client.Buffer.WriteD(user.Race)         //race 41
	client.Buffer.WriteD(user.ClassId)      //classId 45
	client.Buffer.WriteD(0x1)               // ? 49
	client.Buffer.WriteD(user.X)            //x 53
	client.Buffer.WriteD(user.Y)            //y 57
	client.Buffer.WriteD(user.Z)            //z 61

	//buffer.WriteD(83306)  //x 53
	//buffer.WriteD(148115) //y 57
	//buffer.WriteD(-3405)  //z 61

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

}
