package gs2ls

import (
	"crypto/rsa"
	"l2gogameserver/packets"
	"math/big"
)

func BlowFishKey(rsaKey *rsa.PublicKey, bfk []byte) *packets.Buffer {
	buf := packets.Get()
	buf.WriteSingleByte(0x00)
	encryptedData := encryptRSA(rsaKey, bfk)

	buf.WriteD(int32(len(encryptedData)))
	buf.WriteSlice(encryptedData)
	return buf
}

func encryptRSA(pub *rsa.PublicKey, data []byte) []byte {
	encrypted := new(big.Int)
	e := big.NewInt(int64(pub.E))
	payload := new(big.Int).SetBytes(data)
	encrypted.Exp(payload, e, pub.N)
	return encrypted.Bytes()
}
