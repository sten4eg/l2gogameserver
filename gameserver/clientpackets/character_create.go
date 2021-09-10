package clientpackets

import (
	"context"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

type CharCreate struct {
	Name      string
	Race      int32
	Sex       uint8
	ClassId   int32
	Int       int32
	Str       int32
	Con       int32
	Men       int32
	Dex       int32
	Wit       int32
	HairStyle uint8
	HairColor uint8
	Face      uint8
	X         int32
	Y         int32
	Z         int32
	MaxHp     int32
	CutHp     int32
	MaxMp     int32
	CurMp     int32
}

func CharacterCreate(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)
	var charCreate CharCreate

	charCreate.Name = packet.ReadString()

	charCreate.Race = packet.ReadInt32()
	charCreate.Sex = byte(packet.ReadInt32())
	charCreate.ClassId = packet.ReadInt32()
	// зачем клиент присылает статы - ХЗ, они всё равно не используются
	charCreate.Int = packet.ReadInt32()
	charCreate.Str = packet.ReadInt32()
	charCreate.Con = packet.ReadInt32()
	charCreate.Men = packet.ReadInt32()
	charCreate.Dex = packet.ReadInt32()
	charCreate.Wit = packet.ReadInt32()
	//////////////////////
	charCreate.HairStyle = byte(packet.ReadInt32())
	charCreate.HairColor = byte(packet.ReadInt32())
	charCreate.Face = byte(packet.ReadInt32())
	charCreate.validate(client)

}

var (
	ReasonCreationFailed      int32 = 0x00
	ReasonTooManyCharacters   int32 = 0x01
	ReasonNameAlreadyExists   int32 = 0x02
	Reason16EngChars          int32 = 0x03
	ReasonIncorrectName       int32 = 0x04
	ReasonCreateNotAllowed    int32 = 0x05
	REASON_CHOOSE_ANOTHER_SVR int32 = 0x06
	ReasonOk                  int32 = 99
)

func (cc *CharCreate) validate(client *models.Client) {
	lenName := len(cc.Name)
	if (lenName < 1) || (lenName > 16) {
		serverpackets.CharCreateFail(client, Reason16EngChars)
		return
	}

	if cc.Face > 2 {
		serverpackets.CharCreateFail(client, ReasonCreationFailed)
		return
	}

	if ((cc.Sex == 0) && (cc.HairStyle > 4)) || ((cc.Sex) != 0 && (cc.HairStyle > 6)) {
		serverpackets.CharCreateFail(client, ReasonCreationFailed)
		return
	}

	if cc.HairStyle > 3 {
		serverpackets.CharCreateFail(client, ReasonCreationFailed)
		return
	}

	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	row := dbConn.QueryRow(context.Background(), "(SELECT exists(SELECT char_name from characters WHERE char_name = $1))", cc.Name)
	var exist bool
	err = row.Scan(&exist)
	if err != nil {
		serverpackets.CharCreateFail(client, ReasonCreateNotAllowed)
		return
	}
	if exist {
		serverpackets.CharCreateFail(client, ReasonNameAlreadyExists)
		return
	}

	row = dbConn.QueryRow(context.Background(), "SELECT count(*) FROM characters where login = $1", client.Account.Login)
	var i int
	err = row.Scan(&i)
	if err != nil {
		serverpackets.CharCreateFail(client, ReasonCreateNotAllowed)
		return
	}
	if i > 6 {
		serverpackets.CharCreateFail(client, ReasonTooManyCharacters)
		return
	}
	x, y, z := models.GetCreationCoordinates(cc.ClassId)
	_, err = dbConn.Exec(context.Background(), "INSERT INTO characters (char_name, race, sex, class_id, hair_style, hair_color, face,x,y,z,login, base_class, title) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)",
		cc.Name,
		cc.Race,
		cc.Sex,
		cc.ClassId,
		cc.HairStyle,
		cc.HairColor,
		cc.Face,
		x,
		y,
		z,
		client.Account.Login,
		cc.ClassId,
		"")
	if err != nil {
		serverpackets.CharCreateFail(client, ReasonCreateNotAllowed)
		return
	}

	//createChar
	serverpackets.CharCreateOk(client)
}
