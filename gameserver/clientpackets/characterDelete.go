package clientpackets

import (
	"database/sql"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

const (
	ReasonDeletionFailed             int32 = 0x01
	ReasonYouMayNotDeleteClanMember  int32 = 0x02
	ReasonClanLeadersMayNotBeDeleted int32 = 0x03
)

func CharacterDelete(client interfaces.ReciverAndSender, data []byte, db *sql.DB) {
	reader := packets.NewReader(data)

	charSlot := reader.ReadInt32()

	if false { // TODO floodProtection
		client.SendBuf(serverpackets.CharDeleteFail(ReasonDeletionFailed))
		return
	}

	answer := client.MarkToDeleteChar(charSlot)

	switch answer {
	default:
		break
	case -1: // Error
		break
	case 0: // Success!
		client.SendBuf(serverpackets.CharDeleteSuccess())
	// TODO ???
	case 1:
		client.SendBuf(serverpackets.CharDeleteFail(ReasonYouMayNotDeleteClanMember))
	case 2:
		client.SendBuf(serverpackets.CharDeleteFail(ReasonClanLeadersMayNotBeDeleted))

	}

	client.SendBuf(serverpackets.CharSelectionInfo(client, db))

}
