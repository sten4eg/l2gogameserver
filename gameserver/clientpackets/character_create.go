package clientpackets

import (
	"errors"
	"github.com/jackc/pgx"
	"l2gogameserver/packets"
	"log"
)

type CharCreate struct {
	Name      string
	Race      int32
	Sex       int32
	ClassId   int32
	Int       int32
	Str       int32
	Con       int32
	Men       int32
	Dex       int32
	Wit       int32
	HairStyle int32
	HairColor int32
	Face      int32
}

func NewCharacterCreate(data []byte, db *pgx.Conn) (int32, error) {
	var packet = packets.NewReader(data)

	var charCreate CharCreate

	charCreate.Name = packet.ReadString()
	charCreate.Race = packet.ReadInt32()
	charCreate.Sex = packet.ReadInt32()
	charCreate.ClassId = packet.ReadInt32()
	charCreate.Int = packet.ReadInt32()
	charCreate.Str = packet.ReadInt32()
	charCreate.Con = packet.ReadInt32()
	charCreate.Men = packet.ReadInt32()
	charCreate.Dex = packet.ReadInt32()
	charCreate.Wit = packet.ReadInt32()
	charCreate.HairStyle = packet.ReadInt32()
	charCreate.HairColor = packet.ReadInt32()
	charCreate.Face = packet.ReadInt32()
	reason, err := charCreate.validate(db)
	if err != nil {
		return reason, errors.New("fail")
	}

	return 0, nil
}

var (
	ReasonCreationFailed       int32 = 0x00
	REASON_TOO_MANY_CHARACTERS       = 0x01
	ReasonNameAlreadyExists    int32 = 0x02
	Reason16EngChars           int32 = 0x03
	REASON_INCORRECT_NAME            = 0x04
	REASON_CREATE_NOT_ALLOWED        = 0x05
	REASON_CHOOSE_ANOTHER_SVR        = 0x06
)

func (cc *CharCreate) validate(db *pgx.Conn) (int32, error) {
	lenName := len(cc.Name)
	if (lenName < 1) || (lenName > 16) {
		return Reason16EngChars, errors.New("long name")
	}

	if (cc.Face > 2) || (cc.Face < 0) {
		return ReasonCreationFailed, errors.New("wrong face")
	}

	if (cc.HairStyle < 0) || ((cc.Sex == 0) && (cc.HairStyle > 4)) || ((cc.Sex) != 0 && (cc.HairStyle > 6)) {
		return ReasonCreationFailed, errors.New("wrong sex and hairStyle")
	}

	if (cc.HairStyle > 3) || (cc.HairColor < 0) {
		return ReasonCreationFailed, errors.New("wrong hairColor and hairStyle")
	}

	row := db.QueryRow("(SELECT exists(SELECT char_name from characters WHERE char_name = $1))", cc.Name)
	var exist bool
	err := row.Scan(&exist)
	if err != nil {
		log.Fatal(err)
	}
	if exist {
		return ReasonNameAlreadyExists, errors.New("exist name")
	}
	return 0, nil
}
