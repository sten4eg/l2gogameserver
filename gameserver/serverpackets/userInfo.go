package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func UserInfo(character interfaces.CharacterI) []byte {
	buffer := packets.Get()

	x, y, z := character.GetXYZ()

	buffer.WriteSingleByte(0x32)
	buffer.WriteD(x)
	buffer.WriteD(y)
	buffer.WriteD(z)

	buffer.WriteD(0) // Vehicle

	buffer.WriteD(character.GetObjectId()) //objId

	buffer.WriteS(character.GetName()) //name //TODO

	buffer.WriteD(int32(character.GetRace())) //race ordinal //TODO
	buffer.WriteD(character.GetSex())         //sex
	buffer.WriteD(character.GetBaseClass())   //baseClass

	buffer.WriteD(character.GetLevel())                                                                  //level //TODO
	buffer.WriteQ(int64(character.GetCurrentExp()))                                                      //exp
	buffer.WriteF(character.GetPercentFromCurrentLevel(character.GetCurrentExp(), character.GetLevel())) //percent

	buffer.WriteD(int32(character.GetSTR())) //str
	buffer.WriteD(int32(character.GetDEX())) //dex
	buffer.WriteD(int32(character.GetCON())) //con
	buffer.WriteD(int32(character.GetINT())) //int
	buffer.WriteD(int32(character.GetWIT())) //wit
	buffer.WriteD(int32(character.GetMEN())) //men

	buffer.WriteD(character.GetMaxHp()) //Max hp //TODO

	buffer.WriteD(character.GetCurrentHp()) //hp currnebt

	buffer.WriteD(character.GetMaxMp())     //max mp
	buffer.WriteD(character.GetCurrentMp()) //mp

	buffer.WriteD(character.GetCurrentSp()) //sp //TODO
	buffer.WriteD(0)                        //currentLoad

	buffer.WriteD(109020) //maxLoad

	if character.IsActiveWeapon() {
		buffer.WriteD(20) //no weapon
	} else {
		buffer.WriteD(40) //equiped weapon
	}
	characterPaperdoll := character.GetPaperdoll()
	for _, item := range characterPaperdoll {
		buffer.WriteD(item.GetObjectId())
	}
	for _, item := range characterPaperdoll {
		buffer.WriteD(item.GetId())
	}
	for _, item := range characterPaperdoll {
		buffer.WriteD(int32(item.GetEnchant()))
	}

	buffer.WriteD(0) //talisman slot
	buffer.WriteD(0) //Cloack

	buffer.WriteD(4)                   //patack //TODO
	buffer.WriteD(330)                 //atackSpeed
	buffer.WriteD(character.GetPDef()) //pdef
	buffer.WriteD(33)                  //evasionRate
	buffer.WriteD(34)                  //accuracy //TODO
	buffer.WriteD(44)                  //critHit
	buffer.WriteD(3)                   //Matack
	buffer.WriteD(213)                 //M atackSpped

	buffer.WriteD(330) //patackSpeed again?

	buffer.WriteD(47) //mdef

	buffer.WriteD(character.GetPVP())   //pvp
	buffer.WriteD(character.GetKarma()) //karma

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

	buffer.WriteD(character.GetHairStyle()) //hairStyle
	buffer.WriteD(character.GetHairColor()) //hairColor
	buffer.WriteD(character.GetFace())      //face

	buffer.WriteD(1) //IsGM?

	buffer.WriteS(character.GetTitle()) //title

	buffer.WriteD(character.GetClanId()) //clanId
	buffer.WriteD(0)                     //clancrestId
	buffer.WriteD(0)                     //allyId
	buffer.WriteD(0)                     //allyCrestId
	buffer.WriteD(0)                     //RELATION CALCULATE ?

	buffer.WriteSingleByte(0)                                     //mountType
	buffer.WriteSingleByte(byte(character.GetPrivateStoreType())) //privateStoreType
	buffer.WriteSingleByte(0)                                     //hasDwarfCraft
	buffer.WriteD(1)                                              //pk //TODO
	buffer.WriteD(1)                                              //pvp //TODO

	buffer.WriteH(0) //cubic size
	//FOR cubicks

	buffer.WriteSingleByte(0) //PartyRoom

	buffer.WriteD(0) //EFFECTS

	buffer.WriteSingleByte(0) //WATER FLY EARTH

	buffer.WriteD(0) //clanBitmask

	buffer.WriteH(0) // c2 recommendations remaining
	buffer.WriteH(0) // c2 recommendations received //TODO

	buffer.WriteD(0) //npcMountId

	buffer.WriteH(character.GetInventoryLimit()) //inventoryLimit

	buffer.WriteD(character.GetClassId()) //	classId
	buffer.WriteD(0)                      // special effects? circles around player...

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

	buffer.WriteD(0)                       //FAME //TODO
	buffer.WriteD(0)                       //minimap or hellbound
	buffer.WriteD(character.GetVitality()) //vitaliti Point
	buffer.WriteD(0)                       //abnormalEffects

	return buffer.Bytes()
}
