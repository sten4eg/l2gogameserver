package serverpackets

import "l2gogameserver/packets"

func NewCharSelected(user *User) []byte {

	buffer := new(packets.Buffer)
	//cSdSddddddddddFFdQddddddddddddddddBd

	buffer.WriteSingleByte(0x0b) // 1

	buffer.WriteS(user.CharName)     // 11
	buffer.WriteD(user.CharId)       // objId 15
	buffer.WriteS(user.Title.String) //title // 21 2 нуля
	buffer.WriteD(0)                 //TODO sessionId //25
	buffer.WriteD(user.ClanId)       //clanId // 29
	buffer.WriteD(0)                 // ? //33
	buffer.WriteD(user.Sex)          //sex// 37
	buffer.WriteD(user.Race)         //race 41
	buffer.WriteD(user.ClassId)      //classId 45
	buffer.WriteD(0x1)               // ? 49
	buffer.WriteD(user.X)            //x 53
	buffer.WriteD(user.Y)            //y 57
	buffer.WriteD(user.Z)            //z 61

	//buffer.WriteD(83306)  //x 53
	//buffer.WriteD(148115) //y 57
	//buffer.WriteD(-3405)  //z 61

	buffer.WriteF(float64(user.CurHp)) //currentHP 69
	buffer.WriteF(float64(user.CurMp)) //currentMP 77

	buffer.WriteD(user.Sp)         // SP 81
	buffer.WriteQ(int64(user.Exp)) // EXP 89
	buffer.WriteD(user.Level)      // level 93
	buffer.WriteD(user.Karma)      // karma 97
	buffer.WriteD(user.PkKills)    // pk 101

	buffer.WriteD(21) //int 105
	buffer.WriteD(40) //str 109
	buffer.WriteD(43) //con 113
	buffer.WriteD(25) //men 117
	buffer.WriteD(30) //dex 121
	buffer.WriteD(11) //wit 125

	buffer.WriteD(user.OnlineTime) //inGameTime 129
	buffer.WriteD(0)               // ?? 133

	buffer.WriteD(user.ClassId) // 137 classId

	buffer.WriteD(0) // 141
	buffer.WriteD(0) // 145
	buffer.WriteD(0) // 149
	buffer.WriteD(0) // 153

	m := make([]byte, 64)
	buffer.WriteSlice(m) //217

	buffer.WriteD(0) // 221

	return buffer.Bytes()

}
