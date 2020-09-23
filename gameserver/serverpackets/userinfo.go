package serverpackets

import "l2gogameserver/packets"

func NewUserInfo() []byte {
	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x32)

	buffer.WriteD(82744)  //x
	buffer.WriteD(148536) //y
	buffer.WriteD(3400)   //z

	buffer.WriteD(0) // Vehicle

	buffer.WriteD(1) //objId

	buffer.WriteS("test") //name

	buffer.WriteD(0) //race
	buffer.WriteD(1) //sex
	buffer.WriteD(0) //baseClass

	buffer.WriteD(3)      //level
	buffer.WriteQ(500)    //exp
	buffer.WriteF(0.1234) //percent

	buffer.WriteD(5) //str
	buffer.WriteD(5) //dex
	buffer.WriteD(5) //con
	buffer.WriteD(5) //int
	buffer.WriteD(5) //wit
	buffer.WriteD(5) //men

	buffer.WriteD(444) //Max hp

	buffer.WriteD(444) //hp

	buffer.WriteD(1000) //max mp
	buffer.WriteD(1000) //mp

	buffer.WriteD(100) //sp
	buffer.WriteD(1)   //currentLoad

	buffer.WriteD(100) //maxLoad

	buffer.WriteD(20) //no weapon

	//FOR
	x := make([]byte, 312)
	buffer.WriteSlice(x)
	//FOR
	//FOR

	buffer.WriteD(0) //talisman slot
	buffer.WriteD(0) //Cloack

	buffer.WriteD(20) //patack
	buffer.WriteD(20) //atackSpeed
	buffer.WriteD(0)  //pdef
	buffer.WriteD(0)  //evasionRate
	buffer.WriteD(0)  //accuracy
	buffer.WriteD(0)  //critHit
	buffer.WriteD(0)  //Matack
	buffer.WriteD(0)  //M atackSpped

	buffer.WriteD(0) //patackSpeed again?

	buffer.WriteD(0) //mdef

	buffer.WriteD(0) //pvp
	buffer.WriteD(0) //karma

	buffer.WriteD(0) //runSpeed
	buffer.WriteD(0) //walkspeed
	buffer.WriteD(0) //swimRunSpeed
	buffer.WriteD(0) //swimWalkSpeed
	buffer.WriteD(0) //flyRunSpeed
	buffer.WriteD(0) //flyWalkSpeed
	buffer.WriteF(0) //moveMultipler
	buffer.WriteF(0) //atackSpeedMultiplier

	buffer.WriteF(0) //collisionRadius
	buffer.WriteF(0) //collisionHeight

	buffer.WriteD(0) //hairStyle
	buffer.WriteD(1) //hairColor
	buffer.WriteD(0) //face

	buffer.WriteD(1) //IsGM?

	buffer.WriteS("Hi") //title

	buffer.WriteD(0) //clanId
	buffer.WriteD(0) //clancrestId
	buffer.WriteD(0) //allyId
	buffer.WriteD(0) //allyCrestId
	buffer.WriteD(0) //RELATION CALCULATE ?

	buffer.WriteD(0) //mountType
	buffer.WriteD(0) //privateStoreType
	buffer.WriteD(0) //hasDwarfCraft
	buffer.WriteD(0) //pk
	buffer.WriteD(0) //pvp

	buffer.WriteH(0) //cubics
	//FOR

	buffer.WriteSingleByte(0) //PartyRoom

	buffer.WriteD(0) //EFFECTS

	buffer.WriteSingleByte(0) //WATER FLY EARTH

	buffer.WriteD(0) //clanBitmask

	buffer.WriteH(0) // c2 recommendations remaining
	buffer.WriteH(0) // c2 recommendations received

	buffer.WriteD(0) //npcMountId

	buffer.WriteH(10) //inventoryLimit

	buffer.WriteD(1) //	classId
	buffer.WriteD(0)

	buffer.WriteD(300) //MaxCP
	buffer.WriteD(300) //CurrentCp

	buffer.WriteSingleByte(0) //team

	buffer.WriteD(0) //clanCrest

	buffer.WriteSingleByte(0) //isNoble
	buffer.WriteSingleByte(1) //isHero

	buffer.WriteSingleByte(0) //Fishing??
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	buffer.WriteSingleByte(0) //// changes the Speed display on Status Window

	buffer.WriteD(0) // changes the text above CP on Status Window
	buffer.WriteD(0)

	buffer.WriteD(0) //titleColor

	buffer.WriteD(0) // CursedWEAPON

	buffer.WriteD(0) //TransormDisplayId

	//attribute
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)

	buffer.WriteD(0) //agationId

	buffer.WriteD(0)    //FAME
	buffer.WriteD(1)    //minimap or hellbound
	buffer.WriteD(2000) //vitaliti Point
	buffer.WriteD(0)    //abnormalEffects

	return buffer.Bytes()
}
