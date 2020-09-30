package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func NewCharInfo(user *models.CurrentChar) []byte {

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0x31)
	//q.WriteD(-75122)
	//q.WriteD(258213)
	//q.WriteD(-3108)
	buffer.WriteD(user.Spawn.X) //x 53
	buffer.WriteD(user.Spawn.Y) //y 57
	buffer.WriteD(user.Spawn.Z) //z 61

	buffer.WriteD(0) // Vehicle

	i := 3
	if user.CharId == 3 {
		i = 4
	} else {
		i = 3
	}
	buffer.WriteD(int32(i)) //objId

	buffer.WriteS("st") //name //TODO

	buffer.WriteD(4) //race ordinal //TODO
	buffer.WriteD(1) //sex
	buffer.WriteD(0) //baseClass

	////////FOR FOR
	m := make([]byte, 208)
	buffer.WriteSlice(m)
	buffer.WriteD(0) //talisman
	buffer.WriteD(0) //cloack

	buffer.WriteD(0) // pvpFlag
	buffer.WriteD(0) //karma

	buffer.WriteD(0) //MatackSpeed
	buffer.WriteD(0) //_pAtkSpd

	buffer.WriteD(0) // ??

	buffer.WriteD(115) //runSpeed
	buffer.WriteD(80)  //walkspeed
	buffer.WriteD(50)  //swimRunSpeed
	buffer.WriteD(50)  //swimWalkSpeed
	buffer.WriteD(0)   //flyRunSpeed
	buffer.WriteD(0)   //flyWalkSpeed
	buffer.WriteD(0)   //flyRunSpeed again
	buffer.WriteD(0)   //flyWalkSpeed again
	buffer.WriteF(1.1) //moveMultipler

	buffer.WriteF(1.21) //atackSpeedMultiplier

	buffer.WriteF(8.0)  //collisionRadius
	buffer.WriteF(23.5) //collisionHeight

	buffer.WriteD(0) //hairStyle
	buffer.WriteD(0) //hairColor
	buffer.WriteD(0) //face

	buffer.WriteS("") //title

	buffer.WriteD(0) //cursedW
	buffer.WriteD(0) //cursedW
	buffer.WriteD(0) //cursedW
	buffer.WriteD(0) //cursedW

	buffer.WriteSingleByte(1) // standing = 1 sitting = 0
	buffer.WriteSingleByte(1) // running = 1 walking = 0
	buffer.WriteSingleByte(0) //isInCombat

	buffer.WriteSingleByte(0) //!_activeChar.isInOlympiadMode() && _activeChar.isAlikeDead()
	buffer.WriteSingleByte(0) // invisible = 1 visible =0

	buffer.WriteSingleByte(0) // 1-on Strider, 2-on Wyvern, 3-on Great Wolf, 0-no mount
	buffer.WriteSingleByte(0) // privateStore

	buffer.WriteH(0) //CubickSize
	///FOR CUBICs

	buffer.WriteSingleByte(0) // isInPartyMatchRoom

	buffer.WriteD(0) //EFFECTS

	buffer.WriteSingleByte(0) // _activeChar.isInsideZone(ZoneId.WATER) ? 1 : _activeChar.isFlyingMounted() ? 2 : 0

	buffer.WriteH(255)     // Blue value for name (0 = white, 255 = pure blue)
	buffer.WriteD(1000000) // ?

	buffer.WriteD(0) // getClassId().getId()
	buffer.WriteD(0) // ??

	buffer.WriteSingleByte(0) //_activeChar.isMounted() ? 0 : _activeChar.getEnchantEffect()
	buffer.WriteSingleByte(0) //_activeChar.getTeam().getId()

	buffer.WriteD(0)          //_activeChar.getClanCrestLargeId()
	buffer.WriteSingleByte(0) //isNoble
	buffer.WriteSingleByte(0) //_activeChar.isHero() || (_activeChar.isGM() && general().gmHeroAura()) ? 1 : 0

	buffer.WriteSingleByte(0) //Fishing??
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	buffer.WriteD(16777215) //nameColor
	buffer.WriteD(0)        //heading

	buffer.WriteD(0) //getPledgeClass
	buffer.WriteD(0) //getPledgeType

	buffer.WriteD(16777215) //titleColor

	buffer.WriteD(0) //cursedEquiped

	buffer.WriteD(0) //activeChar.getClanId() > 0 ? _activeChar.getClan().getReputationScore() : 0

	buffer.WriteD(0) //getTransformationDisplayId
	buffer.WriteD(0) //getAgathionId
	buffer.WriteD(1) //??
	buffer.WriteD(0) //getAbnormalVisualEffectSpecial

	return buffer.Bytes()
}
