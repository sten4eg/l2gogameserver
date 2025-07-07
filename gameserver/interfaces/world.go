package interfaces

type Positionable interface {
	GetObjectId() int32
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
	CalculateDistanceTo(int32, int32, int32, bool, bool) float64
	//setLocation(Location)
	//setXYZByLoc(ILocational)
}
type WorldRegioner interface {
	GetNeighbors() []WorldRegioner
	GetCharsInRegion() []CharacterI
	AddVisibleChar(CharacterI)
	GetNpcInRegion() []Npcer
	DeleteVisibleChar(CharacterI)
	AddVisibleItems(MyItemInterface)
	GetItemsInRegion() []MyItemInterface
	DeleteVisibleItem(MyItemInterface)
	GetChar(int32) (CharacterI, bool)
	GetItem(int32) (MyItemInterface, bool)
	GetNpc(int32) (Npcer, bool)
	GetCharacterInRegions(int32) CharacterI
	GetX() int32
	GetY() int32
	GetZ() int32
	DropItemChecker() []int32
}
