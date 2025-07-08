package gameserver

import (
	"github.com/puzpuzpuz/xsync"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"strconv"
)

// OnlineCharacters - мапа со всеми чарами которые онлайн
var OnlineCharacters *xsync.MapOf[interfaces.CharacterI] //TODO что она тут делает ? мб сделать ее в структуре GameServer

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
	strKey := strconv.Itoa(int(character.GetObjectId()))
	OnlineCharacters.Store(strKey, character)
}

func CharOffline(client interfaces.NewClientCtxInterface) {
	currentChar := client.GetAccount().GetCurrentChar()
	if currentChar != nil {
		strKey := strconv.Itoa(int(currentChar.GetObjectId()))
		OnlineCharacters.Delete(strKey)
		currentRegion := currentChar.GetCurrentRegion()
		if currentRegion != nil {
			currentRegion.DeleteVisibleChar(currentChar)
		}
		currentChar.CloseChannels()
		//todo close all character goroutine, save info in DB
		logger.Info.Println("Socket Close For: ", currentChar.GetName())
		//client.RemoveCurrentChar()
	}

}
