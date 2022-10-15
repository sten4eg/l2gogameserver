package serverpackets

import "l2gogameserver/packets"

const (
	Level int32 = 0x01
	EXP   int32 = 0x02
	STR   int32 = 0x03
	DEX   int32 = 0x04
	CON   int32 = 0x05
	INT   int32 = 0x06
	WIT   int32 = 0x07
	MEN   int32 = 0x08

	CurHp int32 = 0x09
	MaxHp int32 = 0x0a
	CurMp int32 = 0x0b
	MaxMp int32 = 0x0c

	Sp      int32 = 0x0d
	CurLoad int32 = 0x0e
	MaxLoad int32 = 0x0f

	PAtk     int32 = 0x11
	AtkSpd   int32 = 0x12
	PDef     int32 = 0x13
	Evasion  int32 = 0x14
	Accuracy int32 = 0x15
	Critical int32 = 0x16
	MAtk     int32 = 0x17
	CastSpd  int32 = 0x18
	MDef     int32 = 0x19
	PVPFlag  int32 = 0x1a
	Karma    int32 = 0x1b
	CurCp    int32 = 0x21
	MaxCp    int32 = 0x22
)

type Attributes struct {
	Id    int32
	Value int32
}

func StatusUpdate(objectId int32, attributes []Attributes) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x18)
	buffer.WriteD(objectId)
	buffer.WriteD(int32(len(attributes)))

	for _, attr := range attributes {
		buffer.WriteD(attr.Id)
		buffer.WriteD(attr.Value)
	}

	return buffer
}
