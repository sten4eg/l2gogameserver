package serverpackets

import "l2gogameserver/packets"

func NewUserInfo() []byte {
	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x32)
	//q.WriteD(-75122)
	//q.WriteD(258213)
	//q.WriteD(-3108)
	buffer.WriteD(-75122) //x 53
	buffer.WriteD(258213) //y 57
	buffer.WriteD(-3108)  //z 61

	buffer.WriteD(0) // Vehicle

	buffer.WriteD(1) //objId

	buffer.WriteS("q") //name //TODO

	buffer.WriteD(0) //race ordinal //TODO
	buffer.WriteD(1) //sex
	buffer.WriteD(0) //baseClass

	buffer.WriteD(1) //level //TODO
	buffer.WriteQ(0) //exp
	buffer.WriteF(0) //percent

	buffer.WriteD(40) //str
	buffer.WriteD(30) //dex
	buffer.WriteD(43) //con //TODO
	buffer.WriteD(21) //int
	buffer.WriteD(11) //wit
	buffer.WriteD(25) //men

	buffer.WriteD(163) //Max hp //TODO

	buffer.WriteD(163) //hp currnebt

	buffer.WriteD(52) //max mp
	buffer.WriteD(52) //mp

	buffer.WriteD(0) //sp //TODO
	buffer.WriteD(0) //currentLoad

	buffer.WriteD(109020) //maxLoad

	buffer.WriteD(20) //no weapon

	//FOR
	x := make([]byte, 312)
	buffer.WriteSlice(x)
	//FOR
	//FOR

	buffer.WriteD(0) //talisman slot
	buffer.WriteD(0) //Cloack

	buffer.WriteD(4)   //patack //TODO
	buffer.WriteD(330) //atackSpeed
	buffer.WriteD(73)  //pdef
	buffer.WriteD(35)  //evasionRate
	buffer.WriteD(36)  //accuracy //TODO
	buffer.WriteD(44)  //critHit
	buffer.WriteD(3)   //Matack
	buffer.WriteD(213) //M atackSpped

	buffer.WriteD(330) //patackSpeed again?

	buffer.WriteD(48) //mdef

	buffer.WriteD(0) //pvp
	buffer.WriteD(0) //karma

	buffer.WriteD(115) //runSpeed
	buffer.WriteD(80)  //walkspeed
	buffer.WriteD(50)  //swimRunSpeed
	buffer.WriteD(50)  //swimWalkSpeed
	buffer.WriteD(0)   //flyRunSpeed
	buffer.WriteD(0)   //flyWalkSpeed
	buffer.WriteD(0)   //flyRunSpeed again
	buffer.WriteD(0)   //flyWalkSpeed again
	buffer.WriteF(1.1) //moveMultipler
	//xx := []byte{0, 0, 0, 0, 41, 92, 243, 63}
	//buffer.WriteSlice(xx)
	buffer.WriteF(1.21) //atackSpeedMultiplier

	buffer.WriteF(8.0)  //collisionRadius
	buffer.WriteF(23.5) //collisionHeight

	buffer.WriteD(0) //hairStyle
	buffer.WriteD(1) //hairColor
	buffer.WriteD(0) //face

	buffer.WriteD(0) //IsGM?

	buffer.WriteS("q") //title

	buffer.WriteD(0) //clanId
	buffer.WriteD(0) //clancrestId
	buffer.WriteD(0) //allyId
	buffer.WriteD(0) //allyCrestId
	buffer.WriteD(0) //RELATION CALCULATE ?

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

	buffer.WriteD(0) //	classId
	buffer.WriteD(0)

	buffer.WriteD(65) //MaxCP
	buffer.WriteD(65) //CurrentCp

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

	buffer.WriteD(15530402) //titleColor

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

	buffer.WriteD(0)     //FAME //TODO
	buffer.WriteD(0)     //minimap or hellbound
	buffer.WriteD(20000) //vitaliti Point
	buffer.WriteD(0)     //abnormalEffects

	return buffer.Bytes()
}
