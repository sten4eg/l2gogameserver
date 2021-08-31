package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/packets"
)

func NewCharInfo(user *models.Character) []byte {

	x, y, z := user.GetXYZ()

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0x31)
	buffer.WriteD(x)
	buffer.WriteD(y)
	buffer.WriteD(z)

	buffer.WriteD(0) // Vehicle

	buffer.WriteD(user.CharId) //objId

	buffer.WriteS(string(user.CharName.Bytes)) //name //TODO

	buffer.WriteD(int32(user.Race)) //race ordinal //TODO
	buffer.WriteD(user.Sex)         //sex
	buffer.WriteD(user.BaseClass)   //baseClass

	for _, v := range getPaperdollOrder() {
		buffer.WriteD(user.Paperdoll[v][1])
	}

	for _, v := range getPaperdollOrder() {
		i := v
		_ = i
		buffer.WriteD(0) // augmented
	}

	buffer.WriteD(0) //talisman
	buffer.WriteD(0) //cloack

	buffer.WriteD(user.PvpKills) // pvpFlag
	buffer.WriteD(user.Karma)    //karma

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

	buffer.WriteD(user.HairStyle) //hairStyle
	buffer.WriteD(user.HairColor) //hairColor
	buffer.WriteD(user.Face)      //face

	buffer.WriteS(user.Title.String) //title

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

	buffer.WriteD(user.ClassId) // getClassId().getId()
	buffer.WriteD(0)            // ??

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
func getPaperdollOrder() []uint8 {
	return []uint8{
		items.PAPERDOLL_UNDER,
		items.PAPERDOLL_HEAD,
		items.PAPERDOLL_RHAND,
		items.PAPERDOLL_LHAND,
		items.PAPERDOLL_GLOVES,
		items.PAPERDOLL_CHEST,
		items.PAPERDOLL_LEGS,
		items.PAPERDOLL_FEET,
		items.PAPERDOLL_CLOAK,
		items.PAPERDOLL_RHAND,
		items.PAPERDOLL_HAIR,
		items.PAPERDOLL_HAIR2,
		items.PAPERDOLL_RBRACELET,
		items.PAPERDOLL_LBRACELET,
		items.PAPERDOLL_DECO1,
		items.PAPERDOLL_DECO2,
		items.PAPERDOLL_DECO3,
		items.PAPERDOLL_DECO4,
		items.PAPERDOLL_DECO5,
		items.PAPERDOLL_DECO6,
		items.PAPERDOLL_BELT,
	}
}
