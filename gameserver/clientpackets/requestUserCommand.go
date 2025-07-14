package clientpackets

import (
	"fmt"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/packets"
	"time"
)

// Более полный список команд https://github.com/L2jOpenSource/L2jOpenSource/blob/master/Interlude/L2J_aCis/aCis_389_LATEST_EXPERIMENTAL/aCis_gameserver/java/net/sf/l2j/gameserver/handler/UserCommandHandler.java#L43
const (
	CommandLoc     int8 = 0  // Команда /loc
	CommandUnstuck int8 = 52 // TODO: телепорт в город по умолчанию скилл кастуется 300 сек.
	CommandTime    int8 = 77 // TODO: Внутриигровое время, на данный момент времени в игре нет, по этому выведем время сервера
)

func RequestUserCommand(client interfaces.NewClientCtxInterface, data []byte) {
	reader := packets.NewReader(data)
	commandID := reader.ReadInt8()
	currentChar := client.GetAccount().GetCurrentChar()

	switch commandID {
	case CommandLoc:
		handleLocCommand(client, currentChar)
	case CommandUnstuck:
		handleUnstuckCommand(client, currentChar)
	case CommandTime:
		handleTimeCommand(client)
	default:
		logger.LogError("неизвестная команда пользователя: %d", commandID)
	}
}

func handleLocCommand(client interfaces.NewClientCtxInterface, char interfaces.CharacterI) {
	msg := sysmsg.LocTiS1S2S3
	msg.AddInt(char.GetX())
	msg.AddInt(char.GetY())
	msg.AddInt(char.GetZ())
	client.EncryptAndSend(sysmsg.SystemMessage(msg))
}

func handleUnstuckCommand(client interfaces.NewClientCtxInterface, char interfaces.CharacterI) {
	// Здесь будет логика для unstuck
}

func handleTimeCommand(client interfaces.NewClientCtxInterface) {
	t := time.Now()
	hours := fmt.Sprintf("%02d", t.Hour())
	minutes := fmt.Sprintf("%02d", t.Minute())
	msg := sysmsg.TimeS1S2InTheDay
	msg.AddString(hours)
	msg.AddString(minutes)
	client.EncryptAndSend(sysmsg.SystemMessage(msg))
}
