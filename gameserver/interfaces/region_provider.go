package interfaces

type RegionProvider interface {
	GetRegion(x, y, z int32) WorldRegioner
}
