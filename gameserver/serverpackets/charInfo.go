package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func CharInfo(character interfaces.CharacterI) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)
	x, y, z := character.GetXYZ()

	buffer.WriteSingleByte(0x31)
	buffer.WriteD(x)
	buffer.WriteD(y)
	buffer.WriteD(z)

	buffer.WriteD(0) // Vehicle

	buffer.WriteD(character.GetObjectId()) //objId

	buffer.WriteS(character.GetName()) //name //TODO

	buffer.WriteD(int32(character.GetRace())) //race ordinal //TODO
	buffer.WriteD(character.GetSex())         //sex
	buffer.WriteD(character.GetBaseClass())   //baseClass

	paperdoll := character.GetPaperdoll()

	for _, index := range getPaperdollOrder() {
		buffer.WriteD(paperdoll[index].GetId())
	}

	for _, v := range getPaperdollOrder() {
		_ = v
		buffer.WriteD(0) // augmented
	}

	buffer.WriteD(0) //talisman
	buffer.WriteD(0) //cloack   _activeChar.getInventory().canEquipCloak() ? 1 : 0

	buffer.WriteD(0)                    // pvpFlag The PvP Flag state of the L2PcInstance (0=White, 1=Purple)
	buffer.WriteD(character.GetKarma()) //karma  The Karma of the L2PcInstance (if higher than 0, the name of the L2PcInstance appears in red)

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

	buffer.WriteD(character.GetHairStyle()) //hairStyle
	buffer.WriteD(character.GetHairColor()) //hairColor
	buffer.WriteD(character.GetFace())      //face

	buffer.WriteS(character.GetTitle()) //title

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

	buffer.WriteH(0)       // Blue value for name (0 = white, 255 = pure blue)
	buffer.WriteD(1000000) // ?

	buffer.WriteD(character.GetClassId()) // getClassId().getId()
	buffer.WriteD(0)                      // ??

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

// todo что это тут делает?
// TODO удалить модель
func getPaperdollOrder() []uint8 {
	return []uint8{
		models.PAPERDOLL_UNDER,
		models.PAPERDOLL_HEAD,
		models.PAPERDOLL_RHAND,
		models.PAPERDOLL_LHAND,
		models.PAPERDOLL_GLOVES,
		models.PAPERDOLL_CHEST,
		models.PAPERDOLL_LEGS,
		models.PAPERDOLL_FEET,
		models.PAPERDOLL_CLOAK,
		models.PAPERDOLL_RHAND,
		models.PAPERDOLL_HAIR,
		models.PAPERDOLL_HAIR2,
		models.PAPERDOLL_RBRACELET,
		models.PAPERDOLL_LBRACELET,
		models.PAPERDOLL_DECO1,
		models.PAPERDOLL_DECO2,
		models.PAPERDOLL_DECO3,
		models.PAPERDOLL_DECO4,
		models.PAPERDOLL_DECO5,
		models.PAPERDOLL_DECO6,
		models.PAPERDOLL_BELT,
	}
}
