package clientpackets

import (
	"database/sql"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"log"
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

const CharacterNameMaxLength = 16
const CharacterMaxNumber = 7
const InsertCharacter = `INSERT INTO characters (object_id, char_name, race, sex, class_id, hair_style, hair_color, face, x, y, z, login, base_class, title) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

const countCharAndExistName = `SELECT *
FROM (SELECT COUNT(object_id) FROM characters WHERE login = $1) as countChar,
     (SELECT exists(SELECT char_name from characters WHERE char_name = $2)) as existCharName`

type charCreateClientInterface interface {
	EncryptAndSend([]byte) error
	GetAccountLogin() string
	SendBuf(*packets.Buffer) error
}

func CharacterCreate(clientI interfaces.NewClientCtxInterface, data []byte, db *sql.DB) {
	client := clientI
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

	if len(name) < 1 || len(name) > CharacterNameMaxLength {
		client.EncryptAndSend(serverpackets.CharCreateFail(Reason16EngChars))
		return
	}

	if face > 2 || face < 0 {
		client.EncryptAndSend(serverpackets.CharCreateFail(ReasonCreationFailed))
		return
	}

	if hairStyle < 0 || (sex == 0 && hairStyle > 4) || (sex != 0 && hairStyle > 6) {
		client.EncryptAndSend(serverpackets.CharCreateFail(ReasonCreationFailed))
		return
	}

	if hairColor > 3 || hairColor < 0 {
		client.EncryptAndSend(serverpackets.CharCreateFail(ReasonCreationFailed))
		return
	}

	var charCount byte
	var exist bool

	err := db.QueryRow(countCharAndExistName, client.GetAccount().GetLogin(), name).Scan(&charCount, &exist)
	if err != nil {
		logger.Error.Panicln(err)
	}

	if charCount > CharacterMaxNumber {
		client.EncryptAndSend(serverpackets.CharCreateFail(ReasonTooManyCharacters))
		return
	}
	if exist {
		client.EncryptAndSend(serverpackets.CharCreateFail(ReasonNameAlreadyExists))
		return
	}

	//TODO проверка что пришел норм classId

	x, y, z := models.GetCreationCoordinates(classId)
	_, err = db.Exec(InsertCharacter, idfactory.GetNext(), name, race, sex, classId, hairStyle, hairColor, face, x, y, z, client.GetAccount().GetLogin(), classId, "")
	if err != nil {
		client.EncryptAndSend(serverpackets.CharCreateFail(ReasonCreateNotAllowed))
	}

	client.SendBuf(serverpackets.CharCreateOk())

	time.Sleep(350) //todo клиент должен отправить RequestExGetOnAirShip и после этого CharSelectionInfo, иначе клиент крашиться
	log.Println(client.GetAccount().GetCurrentChar().GetName())

	client.SendBuf(serverpackets.CharSelectionInfo(clientI, db))

	//log.Println(client.GetAccount().GetCurrentChar().GetName())
	//log.Println(clientI.GetAccount().GetCurrentChar().GetName())
	// Выдача предметов после создания персонажа
	//eq := initial.GetEquipmentByClass(client.GetAccount().GetCurrentChar().GetBaseClass())
	//if eq != nil {
	//	for _, item := range eq.Items {
	//		log.Println(client.GetAccount().GetCurrentChar(), item.Id)
	//		itemData, _ := items.GetItemInfo(item.Id)
	//		AddItem2 := client.GetAccount().GetCurrentChar().GetInventory().AddItem2(int32(item.Id), item.Count, itemData.IsStackable(), db)
	//		if item.Equipped {
	//			client.GetAccount().GetCurrentChar().GetInventory().GetItemByObjectId(AddItem2.GetObjectId()).UseEquippableItem()
	//			//TODO: напялить после создания перса
	//		}
	//	}
	//}
}
