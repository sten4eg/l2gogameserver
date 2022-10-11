package gameserver

import (
	"github.com/puzpuzpuz/xsync"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"strconv"
)

// OnlineCharacters - мапа со всеми чарами которые онлайн
var OnlineCharacters *xsync.MapOf[interfaces.CharacterI]

func GetNetConnByCharacterName(name string) interfaces.ReciverAndSender {
	var result interfaces.ReciverAndSender
	OnlineCharacters.Range(func(key string, value interfaces.CharacterI) bool {
		if value.GetName() == name {
			result = value
			return false
		}
		return true
	})

	return result
}

func GetNetConnByCharObjectId(objectId int32) interfaces.CharacterI {
	strKey := strconv.Itoa(int(objectId))
	result, ok := OnlineCharacters.Load(strKey)
	if !ok {
		return nil
	}
	return result
}

func AddOnlineChar(character interfaces.CharacterI) {
	ch, ok := character.(*models.Character)
	if !ok {
		logger.Error.Panicln("addOnlineCharlogger.Error.Panicln")
	}
	strKey := strconv.Itoa(int(character.GetObjectId()))
	OnlineCharacters.Store(strKey, ch)
}

func CharOffline(client interfaces.ClientInterface) {
	currentChar := client.GetCurrentChar()
	if currentChar != nil {
		strKey := strconv.Itoa(int(currentChar.GetObjectId()))
		OnlineCharacters.Delete(strKey)
		currentRegion := client.GetCurrentChar().GetCurrentRegion()
		if currentRegion != nil {
			currentRegion.DeleteVisibleChar(client.GetCurrentChar())
		}
		client.GetCurrentChar().CloseChannels()
		//todo close all character goroutine, save info in DB
		logger.Info.Println("Socket Close For: ", client.GetCurrentChar().GetName())
		client.RemoveCurrentChar()
	}

}
