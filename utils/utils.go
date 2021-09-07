package utils

type PacketByte struct {
	B []byte
}

// GetB получение массива байт из PacketByte
func (i *PacketByte) GetB() []byte {
	cl := make([]byte, len(i.B))
	_ = copy(cl, i.B)
	return cl
}

// SetB копирует массив байт в PacketByte
func (i *PacketByte) SetB(v []byte) {
	cl := make([]byte, len(v))
	i.B = cl
	copy(i.B, v)
}
