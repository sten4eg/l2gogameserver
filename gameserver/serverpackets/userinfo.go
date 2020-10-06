package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func NewUserInfo(user *models.Character) []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x32)
	buffer.WriteD(user.Coordinates.X) //x 53
	buffer.WriteD(user.Coordinates.Y) //y 57
	buffer.WriteD(user.Coordinates.Z) //z 61

	buffer.WriteD(0) // Vehicle

	buffer.WriteD(user.CharId) //objId

	buffer.WriteS(string(user.CharName.Bytes)) //name //TODO

	buffer.WriteD(user.Race)      //race ordinal //TODO
	buffer.WriteD(user.Sex)       //sex
	buffer.WriteD(user.BaseClass) //baseClass

	buffer.WriteD(user.Level)                                            //level //TODO
	buffer.WriteQ(int64(user.Exp))                                       //exp
	buffer.WriteF(user.GetPercentFromCurrentLevel(user.Exp, user.Level)) //percent

	buffer.WriteD(40) //str
	buffer.WriteD(30) //dex
	buffer.WriteD(43) //con //TODO
	buffer.WriteD(21) //int
	buffer.WriteD(11) //wit
	buffer.WriteD(25) //men

	buffer.WriteD(user.MaxHp) //Max hp //TODO

	buffer.WriteD(user.CurHp) //hp currnebt

	buffer.WriteD(user.MaxMp) //max mp
	buffer.WriteD(user.CurMp) //mp

	buffer.WriteD(user.Sp) //sp //TODO
	buffer.WriteD(0)       //currentLoad

	buffer.WriteD(109020) //maxLoad

	buffer.WriteD(40) //no weapon

	//FOR
	x := make([]byte, 312)
	buffer.WriteSlice(x)
	//FOR
	//FOR

	buffer.WriteD(0) //talisman slot
	buffer.WriteD(0) //Cloack

	buffer.WriteD(4)   //patack //TODO
	buffer.WriteD(330) //atackSpeed
	buffer.WriteD(72)  //pdef
	buffer.WriteD(33)  //evasionRate
	buffer.WriteD(34)  //accuracy //TODO
	buffer.WriteD(44)  //critHit
	buffer.WriteD(3)   //Matack
	buffer.WriteD(213) //M atackSpped

	buffer.WriteD(330) //patackSpeed again?

	buffer.WriteD(47) //mdef

	buffer.WriteD(user.PvpKills) //pvp
	buffer.WriteD(user.Karma)    //karma

	buffer.WriteD(999) //runSpeed
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

	buffer.WriteD(0) //IsGM?

	buffer.WriteS(user.Title.String) //title

	buffer.WriteD(user.ClanId) //clanId
	buffer.WriteD(0)           //clancrestId
	buffer.WriteD(0)           //allyId
	buffer.WriteD(0)           //allyCrestId
	buffer.WriteD(0)           //RELATION CALCULATE ?

	buffer.WriteSingleByte(0) //mountType
	buffer.WriteSingleByte(0) //privateStoreType
	buffer.WriteSingleByte(0) //hasDwarfCraft
	buffer.WriteD(1)          //pk //TODO
	buffer.WriteD(1)          //pvp //TODO

	buffer.WriteH(0) //cubic size
	//FOR cubicks

	buffer.WriteSingleByte(0) //PartyRoom

	buffer.WriteD(0) //EFFECTS

	buffer.WriteSingleByte(0) //WATER FLY EARTH

	buffer.WriteD(0) //clanBitmask

	buffer.WriteH(0) // c2 recommendations remaining
	buffer.WriteH(0) // c2 recommendations received //TODO

	buffer.WriteD(0) //npcMountId

	buffer.WriteH(80) //inventoryLimit

	buffer.WriteD(user.ClassId) //	classId
	buffer.WriteD(0)            // special effects? circles around player...

	buffer.WriteD(50) //MaxCP
	buffer.WriteD(50) //CurrentCp

	buffer.WriteSingleByte(0) //mounted air
	buffer.WriteSingleByte(0) //team Id

	buffer.WriteD(0) //ClanCrestLargeId

	buffer.WriteSingleByte(0) //isNoble
	buffer.WriteSingleByte(0) //isHero

	buffer.WriteSingleByte(0) //Fishing??
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	buffer.WriteD(16777215)

	buffer.WriteSingleByte(1) //// changes the Speed display on Status Window

	buffer.WriteD(0) // changes the text above CP on Status Window
	buffer.WriteD(0) // plegue type

	title := (255 & 0xFF) + ((168 & 0xFF) << 8) + ((00 & 0xFF) << 16)
	buffer.WriteD(int32(title)) //titleColor

	buffer.WriteD(0) // CursedWEAPON

	buffer.WriteD(0) //TransormDisplayId

	//attribute
	buffer.WriteH(-2) //attack element //TODO
	buffer.WriteH(0)  //attack elementValue
	buffer.WriteH(0)  //fire
	buffer.WriteH(0)  //water //TODO
	buffer.WriteH(0)  //wind //TODO
	buffer.WriteH(0)  //earth
	buffer.WriteH(0)  //holy
	buffer.WriteH(0)  //dark

	buffer.WriteD(0) //agationId

	buffer.WriteD(0)             //FAME //TODO
	buffer.WriteD(0)             //minimap or hellbound
	buffer.WriteD(user.Vitality) //vitaliti Point
	buffer.WriteD(0)             //abnormalEffects

	return buffer.Bytes()
}
