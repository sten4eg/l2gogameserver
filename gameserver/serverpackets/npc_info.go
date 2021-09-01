package serverpackets

import "l2gogameserver/gameserver/models"

func NeWNpcInfo(client *models.Client) {
	client.Buffer.WriteSingleByte(0x0c)
	client.Buffer.WriteD(268442366)       // objectId
	client.Buffer.WriteD(18342 + 1000000) // npctype id
	client.Buffer.WriteD(1)               //_isAttackable ? 1 : 0
	client.Buffer.WriteD(-71438)          // x
	client.Buffer.WriteD(258005)          // y
	client.Buffer.WriteD(-3104)           // z
	client.Buffer.WriteD(2077)            //_heading
	client.Buffer.WriteD(0x00)            // static 0
	client.Buffer.WriteD(333)             // _mAtkSpd
	client.Buffer.WriteD(278)             // _pAtkSpd
	client.Buffer.WriteD(50)              // _runSpd
	client.Buffer.WriteD(20)              // _WalkSpd
	client.Buffer.WriteD(1)               // _sWimRunSpd
	client.Buffer.WriteD(1)               // _sWimWalkSpd
	client.Buffer.WriteD(1)               // _flyRunSpd
	client.Buffer.WriteD(1)               // _flyWalkSpd
	client.Buffer.WriteD(1)               // _flyRunSpd
	client.Buffer.WriteD(1)               // _flyWalkSpd
	client.Buffer.WriteF(1.1)             // _moveMultiplier
	client.Buffer.WriteF(1)               // _npc.getAttackSpeedMultiplier()
	client.Buffer.WriteF(16.0)            // _collisionRadius
	client.Buffer.WriteF(15.0)            // _collisionHeight
	client.Buffer.WriteD(0)               // right hand Weapon
	client.Buffer.WriteD(0)               // _chest
	client.Buffer.WriteD(0)               // left hand Weapon
	client.Buffer.WriteSingleByte(1)      // name above char 1=true ... ?? //todo - static in l2j
	client.Buffer.WriteSingleByte(0)      // _npc.isRunning() ? 1 : 0
	client.Buffer.WriteSingleByte(0)      // _npc.isInCombat() ? 1 : 0
	client.Buffer.WriteSingleByte(0)      // _npc.isAlikeDead() ? 1 : 0
	client.Buffer.WriteSingleByte(0)      //_isSummoned ? 2 : 0   // invisible ?? 0=false 1=true 2=summoned (only Works if model has a summon animation)
	client.Buffer.WriteD(-1)              // High Five NPCString ID
	client.Buffer.WriteS("")              // _name
	client.Buffer.WriteD(-1)              // High Five NPCString ID
	client.Buffer.WriteS("")              //_title
	client.Buffer.WriteD(0x00)            // Title color 0=client default
	client.Buffer.WriteD(0x00)            // pvp flag
	client.Buffer.WriteD(0x00)            // karma

	client.Buffer.WriteD(0) //_npc.isInvisible() ? _npc.getAbnormalVisualEffects() | AbnormalVisualEffect.STEALTH.getMask() : _npc.getAbnormalVisualEffects()
	client.Buffer.WriteD(0) // clan id
	client.Buffer.WriteD(0) // crest id
	client.Buffer.WriteD(0) // ally id
	client.Buffer.WriteD(0) // all crest

	client.Buffer.WriteSingleByte(0) // _npc.isInsideZone(ZoneId.WATER) ? 1 : _npc.isFlying() ? 2 : 0// C2
	client.Buffer.WriteSingleByte(0) // _npc.getTeam().getId()

	client.Buffer.WriteF(1) // _collisionRadius
	client.Buffer.WriteF(1) // _collisionHeight
	client.Buffer.WriteD(0) //_enchantEffect // C4
	client.Buffer.WriteD(0) // _npc.isFlying() ? 1 : 0 // C6
	client.Buffer.WriteD(0x00)
	client.Buffer.WriteD(0)             //_npc.getColorEffect() // CT1.5 Pet form and skills, Color effect
	client.Buffer.WriteSingleByte(0x01) //_npc.isTargetable() ? 0x01 : 0x00
	client.Buffer.WriteSingleByte(0x01) // _npc.isShoWName() ? 0x01 : 0x00
	client.Buffer.WriteD(0)             //_npc.getAbnormalVisualEffectSpecial()
	client.Buffer.WriteD(0)             // _displayEffect

	client.SaveAndCryptDataInBufferToSend(true)
}
