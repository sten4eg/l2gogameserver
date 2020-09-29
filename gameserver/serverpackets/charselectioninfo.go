package serverpackets

import (
	"database/sql"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"l2gogameserver/gameserver/models"
	"log"
)

type Character struct {
	Login      pgtype.Bytea
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
	CharName   pgtype.Bytea
}

func NewCharSelectionInfo(db *pgx.Conn, client *models.Client) *Character {
	var character Character
	ll := []byte{49, 0, 50, 0}
	rows, err := db.Query("SELECT * FROM characters WHERE Login = $1", ll)
	if err != nil {
		log.Fatal(err)
	}
	Characters := make([]Character, 0)

	for rows.Next() {
		err = rows.Scan(
			&character.Login,
			&character.CharId,
			&character.Level,
			&character.MaxHp,
			&character.CurHp,
			&character.MaxMp,
			&character.CurMp,
			&character.Face,
			&character.HairStyle,
			&character.HairColor,
			&character.Sex,
			&character.X,
			&character.Y,
			&character.Z,
			&character.Exp,
			&character.Sp,
			&character.Karma,
			&character.PvpKills,
			&character.PkKills,
			&character.ClanId,
			&character.Race,
			&character.ClassId,
			&character.BaseClass,
			&character.Title,
			&character.OnlineTime,
			&character.Nobless,
			&character.Vitality,
			&character.CharName,
		)
		if err != nil {
			log.Fatal(err)
		}
		Characters = append(Characters, character)
	}

	//client.Buffer := new(packets.client.Buffer)
	client.Buffer.WriteH(0)
	client.Buffer.WriteSingleByte(0x09)
	client.Buffer.WriteD(int32(len(Characters))) //size char in account

	// Can prevent players from creating new characters (if 0); (if 1, the client will ask if chars may be created (0x13) Response: (0x0D) )
	client.Buffer.WriteD(7)          //char max number
	client.Buffer.WriteSingleByte(0) // delim

	//todo блок который должен повторяться

	for _, char := range Characters {

		client.Buffer.WriteS(string(char.CharName.Bytes)) // Pers name

		client.Buffer.WriteD(char.CharId)              // objId
		client.Buffer.WriteS(string(char.Login.Bytes)) // loginName

		client.Buffer.WriteD(0)           //TODO sessionId
		client.Buffer.WriteD(char.ClanId) //clanId
		client.Buffer.WriteD(0)           // Builder Level

		client.Buffer.WriteD(char.Sex)       //sex
		client.Buffer.WriteD(char.Race)      // race
		client.Buffer.WriteD(char.BaseClass) // baseclass

		client.Buffer.WriteD(0) // active ??

		client.Buffer.WriteD(char.X) //x 53
		client.Buffer.WriteD(char.Y) //y 57
		client.Buffer.WriteD(char.Z) //z 61

		client.Buffer.WriteF(float64(char.CurHp)) //currentHP
		client.Buffer.WriteF(float64(char.CurMp)) //currentMP

		client.Buffer.WriteD(char.Sp)         // SP
		client.Buffer.WriteQ(int64(char.Exp)) // EXP
		client.Buffer.WriteF(0)               // percent ?
		client.Buffer.WriteD(char.Level)      // level

		client.Buffer.WriteD(char.Karma)    // karma
		client.Buffer.WriteD(char.PkKills)  // pk
		client.Buffer.WriteD(char.PvpKills) //pvp

		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)
		//
		//
		m := make([]byte, 104)
		client.Buffer.WriteSlice(m)

		client.Buffer.WriteD(char.HairStyle) //hairStyle
		client.Buffer.WriteD(char.HairColor) //hairColor
		client.Buffer.WriteD(char.Face)      // face

		client.Buffer.WriteF(float64(char.MaxHp)) //max hp
		client.Buffer.WriteF(float64(char.MaxMp)) // max mp

		client.Buffer.WriteD(0)            // days left before
		client.Buffer.WriteD(char.ClassId) //classId

		client.Buffer.WriteD(1)          //auto-selected
		client.Buffer.WriteSingleByte(0) // enchanted
		client.Buffer.WriteD(0)          //augumented

		client.Buffer.WriteD(0) // Currently on retail when you are on character select you don't see your transformation.

		// Implementing it will be waster of resources.
		client.Buffer.WriteD(0)             // Pet ID
		client.Buffer.WriteD(0)             // Pet Level
		client.Buffer.WriteD(0)             // Pet Max Food
		client.Buffer.WriteD(0)             // Pet Current Food
		client.Buffer.WriteF(0)             // Pet Max HP
		client.Buffer.WriteF(0)             // Pet Max MP
		client.Buffer.WriteD(char.Vitality) // H5 Vitality

	}

	return &character
}
