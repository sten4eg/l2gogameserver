package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func NpcInfo(npc interfaces.Npcer) []byte {
	buffer := packets.Get()

	x, y, z := npc.GetCoordinates()

	buffer.WriteSingleByte(0x0c) //12

	buffer.WriteD(npc.GetObjectId())        // objectId
	buffer.WriteD(npc.GetId() + 1000000)    // npctype id
	buffer.WriteD(npc.IsAttackable())       //_isAttackable ? 1 : 0
	buffer.WriteD(x)                        // x
	buffer.WriteD(y)                        // y
	buffer.WriteD(z)                        // z
	buffer.WriteD(npc.GetHeading())         //_heading //53460
	buffer.WriteD(0x00)                     // static 0
	buffer.WriteD(333)                      // _mAtkSpd
	buffer.WriteD(278)                      // _pAtkSpd
	buffer.WriteD(50)                       // _runSpd
	buffer.WriteD(20)                       // _WalkSpd
	buffer.WriteD(50)                       // _sWimRunSpd
	buffer.WriteD(20)                       // _sWimWalkSpd
	buffer.WriteD(0)                        // _flyRunSpd
	buffer.WriteD(0)                        // _flyWalkSpd
	buffer.WriteD(0)                        // _flyRunSpd
	buffer.WriteD(0)                        // _flyWalkSpd
	buffer.WriteF(1.1)                      // _moveMultiplier
	buffer.WriteF(1)                        // _npc.getAttackSpeedMultiplier()
	buffer.WriteF(npc.GetCollisionRadius()) // _collisionRadius
	buffer.WriteF(npc.GetCollisionHeight()) // _collisionHeight
	buffer.WriteD(npc.GetSlotRhand())       // right hand Weapon
	buffer.WriteD(npc.GetSlotLhand())       // left hand Weapon
	buffer.WriteD(0)                        // _chest
	buffer.WriteSingleByte(1)               // name above char 1=true ... ?? //todo - static in l2j
	buffer.WriteSingleByte(0)               // _npc.isRunning() ? 1 : 0
	buffer.WriteSingleByte(0)               // _npc.isInCombat() ? 1 : 0
	buffer.WriteSingleByte(0)               // _npc.isAlikeDead() ? 1 : 0
	buffer.WriteSingleByte(0)               //_isSummoned ? 2 : 0   // invisible ?? 0=false 1=true 2=summoned (only Works if model has a summon animation)
	buffer.WriteD(-1)                       // High Five NPCString ID
	buffer.WriteS("")                       //+2           // _name
	buffer.WriteD(-1)                       // High Five NPCString ID
	buffer.WriteS("")                       //+2            //_title
	buffer.WriteD(0x00)                     // Title color 0=client default
	buffer.WriteD(0x00)                     // pvp flag
	buffer.WriteD(0x00)                     // karma

	buffer.WriteD(0) //_npc.isInvisible() ? _npc.getAbnormalVisualEffects() | AbnormalVisualEffect.STEALTH.getMask() : _npc.getAbnormalVisualEffects()
	buffer.WriteD(0) // clan id
	buffer.WriteD(0) // crest id
	buffer.WriteD(0) // ally id
	buffer.WriteD(0) // all crest

	buffer.WriteSingleByte(0) // _npc.isInsideZone(ZoneId.WATER) ? 1 : _npc.isFlying() ? 2 : 0// C2
	buffer.WriteSingleByte(0) // _npc.getTeam().getId()

	buffer.WriteF(npc.GetCollisionRadius()) // _collisionRadius
	buffer.WriteF(npc.GetCollisionHeight()) // _collisionHeight
	buffer.WriteD(0)                        //_enchantEffect // C4
	buffer.WriteD(0)                        // _npc.isFlying() ? 1 : 0 // C6
	buffer.WriteD(0x00)
	buffer.WriteD(0)             //_npc.getColorEffect() // CT1.5 Pet form and skills, Color effect
	buffer.WriteSingleByte(0x01) //_npc.isTargetable() ? 0x01 : 0x00
	buffer.WriteSingleByte(0x01) // _npc.isShoWName() ? 0x01 : 0x00
	buffer.WriteD(0)             //_npc.getAbnormalVisualEffectSpecial()
	buffer.WriteD(0)             // _displayEffect
	buffer.WriteD(0)             // _displayEffect

	return buffer.Bytes()
}
