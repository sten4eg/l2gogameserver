package clientpackets

import (
	"context"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"time"
)

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

const CharacterNameMaxLenght = 16
const CharacterMaxNumber = 7
const InsertCharacter = `INSERT INTO characters (object_id, char_name, race, sex, class_id, hair_style, hair_color, face, x, y, z, login, base_class, title) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

func CharacterCreate(client interfaces.ReciverAndSender, data []byte) {
	reader := packets.NewReader(data)

	name := reader.ReadString()
	race := reader.ReadInt32()
	sex := byte(reader.ReadInt32())
	classId := reader.ReadInt32()
	_ = reader.ReadInt32() //int
	_ = reader.ReadInt32() //str
	_ = reader.ReadInt32() //con
	_ = reader.ReadInt32() //men
	_ = reader.ReadInt32() //dex
	_ = reader.ReadInt32() //wit
	hairStyle := byte(reader.ReadInt32())
	hairColor := byte(reader.ReadInt32())
	face := byte(reader.ReadInt32())

	if len(name) < 1 || len(name) > CharacterNameMaxLenght {
		client.EncryptAndSend(serverpackets.CharCreateFail(client, Reason16EngChars))
		return
	}

	if face > 2 || face < 0 {
		client.EncryptAndSend(serverpackets.CharCreateFail(client, ReasonCreationFailed))
		return
	}

	if hairStyle < 0 || (sex == 0 && hairStyle > 4) || (sex != 0 && hairStyle > 6) {
		client.EncryptAndSend(serverpackets.CharCreateFail(client, ReasonCreationFailed))
		return
	}

	if hairColor > 3 || hairColor < 0 {
		client.EncryptAndSend(serverpackets.CharCreateFail(client, ReasonCreationFailed))
		return
	}

	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	var charCount byte
	err = dbConn.QueryRow(context.Background(), `SELECT COUNT(object_id) FROM characters WHERE login = $1`, client.GetAccountLogin()).Scan(&charCount)
	if err != nil {
		logger.Error.Panicln(err)
	}
	if charCount > CharacterMaxNumber {
		client.EncryptAndSend(serverpackets.CharCreateFail(client, ReasonTooManyCharacters))
		return
	}

	var exist bool
	err = dbConn.QueryRow(context.Background(), `SELECT exists(SELECT char_name from characters WHERE char_name = $1)`, name).Scan(&exist)
	if err != nil {
		logger.Error.Panicln(err)
	}
	if exist {
		client.EncryptAndSend(serverpackets.CharCreateFail(client, ReasonNameAlreadyExists))
		return
	}

	//TODO проверка что пришел норм classId

	x, y, z := models.GetCreationCoordinates(classId)
	_, err = dbConn.Exec(context.Background(), InsertCharacter, idfactory.GetNext(), name, race, sex, classId, hairStyle, hairColor, face, x, y, z, client.GetAccountLogin(), classId, "")
	if err != nil {
		client.EncryptAndSend(serverpackets.CharCreateFail(client, ReasonCreateNotAllowed))
	}

	client.SendBuf(serverpackets.CharCreateOk())
	time.Sleep(250)
	client.SendBuf(serverpackets.CharSelectionInfo(client))
}
