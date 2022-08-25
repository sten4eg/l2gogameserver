package crypt

func Decrypt(data []byte, inKey []int32) []byte {
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

func Encrypt(data []byte, outKey []int32) []byte {
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

// SimpleEncrypt возвращает зашифрованные байты, с первым
// двумя байтами длинны которые не шифруются
func SimpleEncrypt(data []byte, outKey []int32) []byte {
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

func AppendCheckSum(raw []byte, size int) []byte {
	var chksum int64
	var count = size - 4
	var i int

	for i = 0; i < count; i += 4 {
		var ecx = int64(raw[i])
		ecx |= (int64(raw[i+1]) << 8) & 0xff00
		ecx |= (int64(raw[i+2]) << 0x10) & 0xff0000
		ecx |= (int64(raw[i+3]) << 0x18) & 0xff000000
		chksum ^= ecx
	}

	raw[i] = (byte)(chksum & 0xff)
	raw[i+1] = (byte)((chksum >> 0x08) & 0xff)
	raw[i+2] = (byte)((chksum >> 0x10) & 0xff)
	raw[i+3] = (byte)((chksum >> 0x18) & 0xff)
	return raw
}

func VerifyCheckSum(raw []byte, size int) bool {
	var checksum int64
	count := size - 4
	var i int

	for i = 0; i < count; i += 4 {
		var ecx = int64(raw[i])
		ecx |= (int64(raw[i+1]) << 8) & 0xff00
		ecx |= (int64(raw[i+2]) << 0x10) & 0xff0000
		ecx |= (int64(raw[i+3]) << 0x18) & 0xff000000
		checksum ^= ecx
	}

	var ecx = int64(raw[i])
	ecx |= (int64(raw[i+1]) << 8) & 0xff00
	ecx |= (int64(raw[i+2]) << 0x10) & 0xff0000
	ecx |= (int64(raw[i+3]) << 0x18) & 0xff000000
	return ecx == checksum
}
