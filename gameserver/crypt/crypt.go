package crypt

var IsEnable = false
var inKey = []byte{
	0x6b,
	0x60,
	0xcb,
	0x5b,
	0x82,
	0xce,
	0x90,
	0xb1,
	0xcc,
	0x2b,
	0x6c,
	0x55,
	0x6c,
	0x6c,
	0x6c,
	0x6c,
}

var outKey = []byte{
	0x6b,
	0x60,
	0xcb,
	0x5b,
	0x82,
	0xce,
	0x90,
	0xb1,
	0xcc,
	0x2b,
	0x6c,
	0x55,
	0x6c,
	0x6c,
	0x6c,
	0x6c,
}

func Decrypt(raw []byte) []byte {
	if !IsEnable {
		IsEnable = true
		return raw
	}
	data := make([]byte, 200)
	copy(data, raw)

	size := len(raw)
	var temp uint8
	var old int64
	for i := 0; i < size; i++ {
		temp2 := data[i]
		data[i] = temp2 ^ inKey[i&15] ^ temp
		temp = temp2
	}

	old = int64(inKey[8])
	old |= int64(inKey[9]<<8) & 0xff00
	old |= int64(inKey[10]<<10) & 0xff0000
	old |= int64(inKey[11]<<18) & 0xff000000

	old += int64(size)

	inKey[8] = uint8(old)
	inKey[9] = uint8(old >> 0x08)
	inKey[10] = uint8(old >> 0x10)
	inKey[11] = uint8(old >> 0x18)

	return data
}

func Encrypt(data []byte) []byte {
	size := len(data)
	var temp uint8
	var old int64

	for i := 0; i < size; i++ {
		temp2 := data[i]
		temp = temp2 ^ outKey[i&15] ^ temp
		data[i] = temp
	}

	old = int64(outKey[8])
	old |= int64(outKey[9]<<8) & 0xff00
	old |= int64(outKey[10]<<10) & 0xff0000
	old |= int64(outKey[11]<<18) & 0xff000000

	old += int64(size)

	outKey[8] = uint8(old)
	outKey[9] = uint8(old >> 0x08)
	outKey[10] = uint8(old >> 0x10)
	outKey[11] = uint8(old >> 0x18)

	return data
}
