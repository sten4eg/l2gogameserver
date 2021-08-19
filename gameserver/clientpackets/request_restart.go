package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
)

func NewRequestRestart(data []byte, client *models.Client) {

	//todo need save in db

	_ = data
	serverpackets.NewRestartResponse(client)

	serverpackets.NewCharSelectionInfo(client)
}
