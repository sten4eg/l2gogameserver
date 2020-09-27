package serverpackets

import (
	"database/sql"
	"github.com/jackc/pgx"
	"l2gogameserver/packets"
	"log"
)

type User struct {
	Login      string
	CharId     int32
	Level      int32
	MaxHp      int32
	CurHp      int32
	MaxMp      int32
	CurMp      int32
	Face       int32
	HairStyle  int32
	HairColor  int32
	Sex        int32
	X          int32
	Y          int32
	Z          int32
	Exp        int32
	Sp         int32
	Karma      int32
	PvpKills   int32
	PkKills    int32
	ClanId     int32
	Race       int32
	ClassId    int32
	BaseClass  int32
	Title      sql.NullString
	OnlineTime int32
	Nobless    int32
	Vitality   int32
	CharName   string
}

func NewCharSelectionInfo(db *pgx.Conn) ([]byte, *User) {
	var user User
	row := db.QueryRow("SELECT * FROM characters WHERE Login = $1", "12")
	err := row.Scan(
		&user.Login,
		&user.CharId,
		&user.Level,
		&user.MaxHp,
		&user.CurHp,
		&user.MaxMp,
		&user.CurMp,
		&user.Face,
		&user.HairStyle,
		&user.HairColor,
		&user.Sex,
		&user.X,
		&user.Y,
		&user.Z,
		&user.Exp,
		&user.Sp,
		&user.Karma,
		&user.PvpKills,
		&user.PkKills,
		&user.ClanId,
		&user.Race,
		&user.ClassId,
		&user.BaseClass,
		&user.Title,
		&user.OnlineTime,
		&user.Nobless,
		&user.Vitality,
		&user.CharName,
	)
	if err != nil {
		log.Fatal(11)
	}

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0x09)
	buffer.WriteD(1) //size char in account

	// Can prevent players from creating new characters (if 0); (if 1, the client will ask if chars may be created (0x13) Response: (0x0D) )
	buffer.WriteD(7)          //char max number
	buffer.WriteSingleByte(0) // delim

	//todo блок который должен повторяться

	buffer.WriteS(user.CharName) // Pers name

	buffer.WriteD(user.CharId) // objId
	buffer.WriteS(user.Login)  // loginName

	buffer.WriteD(0)           //TODO sessionId
	buffer.WriteD(user.ClanId) //clanId
	buffer.WriteD(0)           // Builder Level

	buffer.WriteD(user.Sex)       //sex
	buffer.WriteD(user.Race)      // race
	buffer.WriteD(user.BaseClass) // baseclass

	buffer.WriteD(0) // active ??

	buffer.WriteD(user.X) //x 53
	buffer.WriteD(user.Y) //y 57
	buffer.WriteD(user.Z) //z 61

	buffer.WriteF(float64(user.CurHp)) //currentHP
	buffer.WriteF(float64(user.CurMp)) //currentMP

	buffer.WriteD(user.Sp)         // SP
	buffer.WriteQ(int64(user.Exp)) // EXP
	buffer.WriteF(0)               // percent ?
	buffer.WriteD(user.Level)      // level

	buffer.WriteD(user.Karma)    // karma
	buffer.WriteD(user.PkKills)  // pk
	buffer.WriteD(user.PvpKills) //pvp

	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	//
	//
	m := make([]byte, 104)
	buffer.WriteSlice(m)

	buffer.WriteD(user.HairStyle) //hairStyle
	buffer.WriteD(user.HairColor) //hairColor
	buffer.WriteD(user.Face)      // face

	buffer.WriteF(float64(user.MaxHp)) //max hp
	buffer.WriteF(float64(user.MaxMp)) // max mp

	buffer.WriteD(0)            // days left before
	buffer.WriteD(user.ClassId) //classId

	buffer.WriteD(1)          //auto-selected
	buffer.WriteSingleByte(0) // enchanted
	buffer.WriteD(0)          //augumented

	buffer.WriteD(0) // Currently on retail when you are on character select you don't see your transformation.

	// Implementing it will be waster of resources.
	buffer.WriteD(0)             // Pet ID
	buffer.WriteD(0)             // Pet Level
	buffer.WriteD(0)             // Pet Max Food
	buffer.WriteD(0)             // Pet Current Food
	buffer.WriteF(0)             // Pet Max HP
	buffer.WriteF(0)             // Pet Max MP
	buffer.WriteD(user.Vitality) // H5 Vitality
	return buffer.Bytes(), &user
}
