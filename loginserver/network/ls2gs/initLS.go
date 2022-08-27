package ls2gs

import (
	"crypto/rsa"
	"l2gogameserver/packets"
	"math/big"
)

func InitLs(data []byte) *rsa.PublicKey {
	reader := packets.NewReader(data)
	size := reader.ReadInt32()
	b := reader.ReadBytes(int(size))
	b = b[1:]

	var bigI = new(big.Int)
	bigI.SetBytes(b)

	return &rsa.PublicKey{
		N: bigI,
		E: 65537,
	}

}
