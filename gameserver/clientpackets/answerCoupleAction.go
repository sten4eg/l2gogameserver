package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"math"
)

const RadiansToDegrees = 57.29577951308232

func AnswerCoupleAction(client interfaces.ReciverAndSender, data []byte) {
	reader := packets.NewReader(data[2:])

	actionId := reader.ReadInt32()
	answer := reader.ReadInt32()
	charObjId := reader.ReadInt32()

	character := client.GetCurrentChar()
	targetObject := getTargetByObjectId(charObjId, character.GetCurrentRegion())

	if character == nil || targetObject == nil {
		return
	}

	switch target := targetObject.(type) {
	default:
		return
	case interfaces.CharacterI:
		if target.GetMultiSocialTarget() != character.GetObjectId() || target.GetMultiSocialAction() != actionId {
			return
		}
		if answer == 0 {
			target.SendSysMsg(sysmsg.CoupleActionDenied)
		} else if answer == 1 {
			ox, oy, oz := character.GetXYZ()
			mx, my, mz := target.GetXYZ()
			distance := models.CalculateDistance(ox, oy, oz, mx, my, mz, false, false)

			if distance > 125 || distance < 15 || character.GetObjectId() == target.GetObjectId() {
				character.SendSysMsg(sysmsg.TargetDoNotMeetLocRequirements)
				target.SendSysMsg(sysmsg.TargetDoNotMeetLocRequirements)
				return
			}

			heading := calculateHeadingFrom(character, target)
			broadcast.BroadCastBufferToAroundPlayers(client, serverpackets.ExRotation(character.GetObjectId(), heading))
			character.SetHeading(heading)

			heading = calculateHeadingFrom(target, character)
			target.SetHeading(heading)
			broadcast.BroadCastBufferToAroundPlayers(target, serverpackets.ExRotation(target.GetObjectId(), heading))

			broadcast.BroadCastPkgToAroundPlayer(client, serverpackets.SocialAction(character, actionId))
			broadcast.BroadCastPkgToAroundPlayer(target, serverpackets.SocialAction(target, actionId))

		} else if answer == -1 {
			msg := sysmsg.C1IsSetToRefuseCoupleActions
			msg.AddCharacterName(character.GetName())
			target.SendSysMsg(msg)
		}
		target.SetMultiSocialAction(0, 0)
	}
}

func calculateHeadingFrom(character, target interfaces.CharacterI) int32 {
	fromX, fromY, _ := character.GetXYZ()
	toX, toY, _ := target.GetXYZ()

	angleTarget := math.Atan2(float64(toY-fromY), float64(toX-fromX)) * RadiansToDegrees
	if angleTarget < 0 {
		angleTarget += 360
	}
	return int32(angleTarget * 182.044444444)
}
