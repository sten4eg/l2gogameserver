package npc

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/types"
	"l2gogameserver/packets"
)

// NpcMoveToLocation создает пакет движения NPC (0x2F)
func NpcMoveToLocation(location *types.BackwardToLocation, npc interfaces.Npcer) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x2f) // MoveToLocation opcode

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
func BroadcastNpcMovement(npc interfaces.Npcer, location *types.BackwardToLocation, world interfaces.WorldRegioner) {
	// Создаем пакет движения
	movementPacket := NpcMoveToLocation(location, npc)

	// Отправляем всем игрокам в соседних регионах
	for _, region := range world.GetNeighbors() {
		for _, player := range region.GetCharsInRegion() {
			if player != nil {
				player.EncryptAndSend(movementPacket)
			}
		}
	}
}
