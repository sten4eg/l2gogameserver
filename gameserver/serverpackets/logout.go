package serverpackets

func LogoutWithInGameState() []byte {
	return []byte{0x84}
}

func LogoutWithAuthedState() []byte {
	return []byte{0x20}
}
