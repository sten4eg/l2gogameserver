package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestRestart(data []byte, client *models.Client) {

	//todo need save in db

	_ = data
	serverpackets.RestartResponse(client)

	serverpackets.CharSelectionInfo(client)
}
