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
