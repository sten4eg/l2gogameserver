package crypt

var IsEnable = false
var inKey = [16]int32{
	0x6b,
	0x60,
	0xcb,
	0x5b,
	0x82,
	0xce,
	0x90,
	0xb1,
	0xc8,
	0x27,
	0x93,
	0x01,
	0xa1,
	0x6c,
	0x31,
	0x97,
}

var outKey = [16]int32{
	0x6b,
	0x60,
	0xcb,
	0x5b,
	0x82,
	0xce,
	0x90,
	0xb1,
	0xc8,
	0x27,
	0x93,
	0x01,
	0xa1,
	0x6c,
	0x31,
	0x97,
}

func Decrypt(data []byte) []byte {
	if !IsEnable {
		IsEnable = true
		return data
	}

	size := len(data)
	var temp int32
	var old int32
	for i := 0; i < size; i++ {
		temp2 := data[i]
		data[i] = byte(int32(temp2) ^ inKey[i&15] ^ temp)
		temp = int32(temp2)
	}

	old = inKey[8]
	old |= (inKey[9] << 8) & 0xff00
	old |= (inKey[10] << 0x10) & 0xff0000
	old |= (inKey[11] << 0x18) & -16777216

	old += int32(size)

	inKey[8] = old
	inKey[9] = (old >> 0x08) & 0xff
	inKey[10] = (old >> 0x10) & 0xff
	inKey[11] = (old >> 0x18) & 0xff

	return data
}

func Encrypt(data []byte) []byte {
	size := len(data)
	var temp int32
	var old int32

	for i := 0; i < size; i++ {
		temp2 := data[i]
		temp = int32(temp2) ^ outKey[i&15] ^ temp
		data[i] = byte(temp)
	}

	old = (outKey[8]) & 0xff
	old |= (outKey[9] << 0x8) & 0xff00
	old |= (outKey[10] << 0x10) & 0xff0000
	old |= (outKey[11] << 0x18) & -16777216

	old += int32(size)
	outKey[8] = (old) & 0xff
	outKey[9] = (old >> 0x08) & 0xff
	outKey[10] = (old >> 0x10) & 0xff
	outKey[11] = (old >> 0x18) & 0xff

	return data
}

func SimpleEncrypt(data []byte) []byte {
	size := len(data) - 2
	var temp int32
	var old int32

	for i := 0; i < size; i++ {
		temp2 := data[i+2]
		temp = int32(temp2) ^ outKey[i&15] ^ temp
		data[i+2] = byte(temp)
	}

	old = (outKey[8]) & 0xff
	old |= (outKey[9] << 0x8) & 0xff00
	old |= (outKey[10] << 0x10) & 0xff0000
	old |= (outKey[11] << 0x18) & -16777216

	old += int32(size)
	outKey[8] = (old) & 0xff
	outKey[9] = (old >> 0x08) & 0xff
	outKey[10] = (old >> 0x10) & 0xff
	outKey[11] = (old >> 0x18) & 0xff

	return data
}
