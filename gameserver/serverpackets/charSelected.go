package serverpackets

import (
	"database/sql"
	"l2gogameserver/gameserver/interfaces"

	"l2gogameserver/packets"
)

func CharSelected(user interfaces.CharacterI, db *sql.DB, login string) []byte {
	//client, ok := clientI.(*models.ClientCtx)
	//if !ok {
	//	return []byte{}
	//}
	//user := client.GetAccount().GetCurrentChar()
	buffer := packets.Get()

	x, y, z := user.GetXYZ()
	buffer.WriteSingleByte(0x0b) // 1

	buffer.WriteS(user.GetName())        // 11
	buffer.WriteD(user.GetObjectId())    // objId 15
	buffer.WriteS(user.GetTitle())       //title // 21 2 нуля
	buffer.WriteD(0)                     //TODO sessionId //25
	buffer.WriteD(user.GetClanId())      //clanId // 29
	buffer.WriteD(0)                     // ? //33
	buffer.WriteD(user.GetSex())         //sex// 37
	buffer.WriteD(int32(user.GetRace())) //race 41
	buffer.WriteD(user.GetClassId())     //classId 45
	buffer.WriteD(0x1)                   // ? 49
	buffer.WriteD(x)                     //x 53
	buffer.WriteD(y)                     //y 57
	buffer.WriteD(z)                     //z 61

	buffer.WriteF(float64(user.GetCurrentHp())) //currentHP 69
	buffer.WriteF(float64(user.GetCurrentMp())) //currentMP 77

	buffer.WriteD(user.GetCurrentSp())         // SP 81
	buffer.WriteQ(int64(user.GetCurrentExp())) // EXP 89
	buffer.WriteD(user.GetLevel())             // level 93
	buffer.WriteD(user.GetKarma())             // karma 97
	buffer.WriteD(user.GetPK())                // pk 101

	buffer.WriteD(21) //int 105
	buffer.WriteD(40) //str 109
	buffer.WriteD(43) //con 113
	buffer.WriteD(25) //men 117
	buffer.WriteD(30) //dex 121
	buffer.WriteD(11) //wit 125

	buffer.WriteD(user.GetOnlineTime()) //inGameTime 129
	buffer.WriteD(0)                    // ?? 133

	buffer.WriteD(user.GetClassId()) // 137 classId TODO уже выше есть ClassId

	buffer.WriteD(0) // 141
	buffer.WriteD(0) // 145
	buffer.WriteD(0) // 149
	buffer.WriteD(0) // 153

	m := make([]byte, 64)
	buffer.WriteSlice(m) //217

	buffer.WriteD(0) // 221

	//	client.SetCurrentChar(client.GetAccount().GetChar(int(client.Account.CharSlotSelected)))
	//client.CurrentChar = client.Account.Char[client.Account.CharSlotSelected]
	//	client.CurrentChar.SockConn = client.GetConn()
	//TODO Load загрузка всех данных выбранного чара
	user.Load(db, login)

	return buffer.Bytes()

}
