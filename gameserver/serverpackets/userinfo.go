package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
)

func NewUserInfo(character *models.Character, client *models.Client) {

	x, y, z := character.GetXYZ()

	client.Buffer.WriteSingleByte(0x32)
	client.Buffer.WriteD(x)
	client.Buffer.WriteD(y)
	client.Buffer.WriteD(z)

	client.Buffer.WriteD(0) // Vehicle

	client.Buffer.WriteD(character.CharId) //objId

	client.Buffer.WriteS(string(character.CharName.Bytes)) //name //TODO

	client.Buffer.WriteD(character.Race)      //race ordinal //TODO
	client.Buffer.WriteD(character.Sex)       //sex
	client.Buffer.WriteD(character.BaseClass) //baseClass

	client.Buffer.WriteD(character.Level)                                                      //level //TODO
	client.Buffer.WriteQ(int64(character.Exp))                                                 //exp
	client.Buffer.WriteF(character.GetPercentFromCurrentLevel(character.Exp, character.Level)) //percent

	client.Buffer.WriteD(40) //str
	client.Buffer.WriteD(40) //dex
	client.Buffer.WriteD(40) //con
	client.Buffer.WriteD(40) //int
	client.Buffer.WriteD(40) //wit
	client.Buffer.WriteD(40) //men

	client.Buffer.WriteD(character.MaxHp) //Max hp //TODO

	client.Buffer.WriteD(character.CurHp) //hp currnebt

	client.Buffer.WriteD(character.MaxMp) //max mp
	client.Buffer.WriteD(character.CurMp) //mp

	client.Buffer.WriteD(character.Sp) //sp //TODO
	client.Buffer.WriteD(0)            //currentLoad

	client.Buffer.WriteD(109020) //maxLoad

	if character.IsActiveWeapon() {
		client.Buffer.WriteD(20) //no weapon
	} else {
		client.Buffer.WriteD(40) //equiped weapon
	}

	for _, slot := range items.GetPaperdollOrder() {
		client.Buffer.WriteD(character.Paperdoll[slot][0]) //objId
	}
	for _, slot := range items.GetPaperdollOrder() {
		client.Buffer.WriteD(character.Paperdoll[slot][1]) //itemId
	}
	for _, slot := range items.GetPaperdollOrder() {
		client.Buffer.WriteD(character.Paperdoll[slot][2]) //enchant
	}

	client.Buffer.WriteD(0) //talisman slot
	client.Buffer.WriteD(0) //Cloack

	client.Buffer.WriteD(4)   //patack //TODO
	client.Buffer.WriteD(330) //atackSpeed
	client.Buffer.WriteD(72)  //pdef
	client.Buffer.WriteD(33)  //evasionRate
	client.Buffer.WriteD(34)  //accuracy //TODO
	client.Buffer.WriteD(44)  //critHit
	client.Buffer.WriteD(3)   //Matack
	client.Buffer.WriteD(213) //M atackSpped

	client.Buffer.WriteD(330) //patackSpeed again?

	client.Buffer.WriteD(47) //mdef

	client.Buffer.WriteD(character.PvpKills) //pvp
	client.Buffer.WriteD(character.Karma)    //karma

	client.Buffer.WriteD(999) //runSpeed
	client.Buffer.WriteD(80)  //walkspeed
	client.Buffer.WriteD(50)  //swimRunSpeed
	client.Buffer.WriteD(50)  //swimWalkSpeed
	client.Buffer.WriteD(0)   //flyRunSpeed
	client.Buffer.WriteD(0)   //flyWalkSpeed
	client.Buffer.WriteD(0)   //flyRunSpeed again
	client.Buffer.WriteD(0)   //flyWalkSpeed again
	client.Buffer.WriteF(1.1) //moveMultipler

	client.Buffer.WriteF(1.21) //atackSpeedMultiplier

	client.Buffer.WriteF(8.0)  //collisionRadius
	client.Buffer.WriteF(23.5) //collisionHeight

	client.Buffer.WriteD(character.HairStyle) //hairStyle
	client.Buffer.WriteD(character.HairColor) //hairColor
	client.Buffer.WriteD(character.Face)      //face

	client.Buffer.WriteD(0) //IsGM?

	client.Buffer.WriteS(character.Title.String) //title

	client.Buffer.WriteD(character.ClanId) //clanId
	client.Buffer.WriteD(0)                //clancrestId
	client.Buffer.WriteD(0)                //allyId
	client.Buffer.WriteD(0)                //allyCrestId
	client.Buffer.WriteD(0)                //RELATION CALCULATE ?

	client.Buffer.WriteSingleByte(0) //mountType
	client.Buffer.WriteSingleByte(0) //privateStoreType
	client.Buffer.WriteSingleByte(0) //hasDwarfCraft
	client.Buffer.WriteD(1)          //pk //TODO
	client.Buffer.WriteD(1)          //pvp //TODO

	client.Buffer.WriteH(0) //cubic size
	//FOR cubicks

	client.Buffer.WriteSingleByte(0) //PartyRoom

	client.Buffer.WriteD(0) //EFFECTS

	client.Buffer.WriteSingleByte(0) //WATER FLY EARTH

	client.Buffer.WriteD(0) //clanBitmask

	client.Buffer.WriteH(0) // c2 recommendations remaining
	client.Buffer.WriteH(0) // c2 recommendations received //TODO

	client.Buffer.WriteD(0) //npcMountId

	client.Buffer.WriteH(80) //inventoryLimit

	client.Buffer.WriteD(character.ClassId) //	classId
	client.Buffer.WriteD(0)                 // special effects? circles around player...

	client.Buffer.WriteD(50) //MaxCP
	client.Buffer.WriteD(50) //CurrentCp

	client.Buffer.WriteSingleByte(0) //mounted air
	client.Buffer.WriteSingleByte(0) //team Id

	client.Buffer.WriteD(0) //ClanCrestLargeId

	client.Buffer.WriteSingleByte(0) //isNoble
	client.Buffer.WriteSingleByte(0) //isHero

	client.Buffer.WriteSingleByte(0) //Fishing??
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)

	client.Buffer.WriteD(16777215)

	client.Buffer.WriteSingleByte(1) //// changes the Speed display on Status Window

	client.Buffer.WriteD(0) // changes the text above CP on Status Window
	client.Buffer.WriteD(0) // plegue type

	title := (255 & 0xFF) + ((168 & 0xFF) << 8) + ((00 & 0xFF) << 16)
	client.Buffer.WriteD(int32(title)) //titleColor

	client.Buffer.WriteD(0) // CursedWEAPON

	client.Buffer.WriteD(0) //TransormDisplayId

	//attribute
	client.Buffer.WriteH(-2) //attack element //TODO
	client.Buffer.WriteH(0)  //attack elementValue
	client.Buffer.WriteH(0)  //fire
	client.Buffer.WriteH(0)  //water //TODO
	client.Buffer.WriteH(0)  //wind //TODO
	client.Buffer.WriteH(0)  //earth
	client.Buffer.WriteH(0)  //holy
	client.Buffer.WriteH(0)  //dark

	client.Buffer.WriteD(0) //agationId

	client.Buffer.WriteD(0)                  //FAME //TODO
	client.Buffer.WriteD(0)                  //minimap or hellbound
	client.Buffer.WriteD(character.Vitality) //vitaliti Point
	client.Buffer.WriteD(0)                  //abnormalEffects

	client.SaveAndCryptDataInBufferToSend(true)
}
