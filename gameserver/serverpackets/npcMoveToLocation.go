package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/types"
	"l2gogameserver/packets"
)

// NpcMoveToLocation создает пакет движения для NPC
func NpcMoveToLocation(location *types.BackwardToLocation, npc interfaces.Npcer) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x2f) // Тот же опкод что и для персонажей

	buffer.WriteD(npc.GetObjectId())

	buffer.WriteD(location.TargetX)
	buffer.WriteD(location.TargetY)
	buffer.WriteD(location.TargetZ)

	buffer.WriteD(location.OriginX)
	buffer.WriteD(location.OriginY)
	buffer.WriteD(location.OriginZ)

	return buffer.Bytes()
}

// BroadcastNpcMovement отправляет пакет движения NPC всем игрокам в регионе
func BroadcastNpcMovement(npc interfaces.Npcer, location *types.BackwardToLocation) []byte {
	return NpcMoveToLocation(location, npc)
}
