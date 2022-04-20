package interfaces

type Identifier interface {
	GetId() int32
}
type UniquerId interface {
	GetObjectId() int32
}
type Namer interface {
	GetName() string
}
type Positionable interface {
	SetX(int32)
	SetY(int32)
	SetZ(int32)
	SetXYZ(int32, int32, int32)
	SetHeading(int32)
	SetInstanceId(int32)
	GetX() int32
	GetY() int32
	GetZ() int32
	GetXYZ() (int32, int32, int32)
	GetCurrentRegion() WorldRegioner
	//setLocation(Location)
	//setXYZByLoc(ILocational)
}
type WorldRegioner interface {
	GetNeighbors() []WorldRegioner
	GetCharsInRegion() []CharacterI
	AddVisibleChar(CharacterI)
	GetNpcInRegion() []Npcer
	DeleteVisibleChar(CharacterI)
}
type Npcer interface {
	UniquerId
	Identifier
}

type CharacterI interface {
	Positionable
	Namer
	UniquerId
	EncryptAndSend(data []byte)
	CloseChannels()
	GetClassId() int32
}
type ReciverAndSender interface {
	Receive() (opcode byte, data []byte, e error)
	AddLengthAndSand(d []byte)
	Send(data []byte)
	EncryptAndSend(data []byte)
	CryptAndReturnPackageReadyToShip(data []byte) []byte
	GetCurrentChar() CharacterI
}
