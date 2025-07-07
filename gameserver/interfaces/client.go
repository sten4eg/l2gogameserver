package interfaces

import (
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/packets"
	"net"
)

type ReciverAndSender interface {
	Receive() (opcode byte, data []byte, err error)
	AddLengthAndSand(data []byte)
	Send(data []byte)
	SendBuf(buffer *packets.Buffer) error
	EncryptAndSend(data []byte) error
	SendSysMsg(q interface{}, options ...string) error
	CryptAndReturnPackageReadyToShip(data []byte) []byte
	GetCurrentChar() CharacterI

	GetAccountLogin() string

	GetObjectIdForSlot(int32) int32
	MarkToDeleteChar(int32) int8
}

type ClientCtxInterface interface {
	ReciverAndSender
	SetLogin(login string)
	RemoveCurrentChar()
	SetState(state clientStates.State)
	GetState() clientStates.State
	SetSessionKey(playOk1, playOk2, loginOk1, loginOk2 uint32)
	GetSessionKey() (playOk1, playOk2, loginOk1, loginOk2 uint32)
	GetAccountLogin() string
	CloseConnection()
	GetRemoteAddr() net.Addr
}
