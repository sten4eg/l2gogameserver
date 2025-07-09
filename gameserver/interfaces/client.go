package interfaces

import (
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/packets"
	"net"
)

type NewClientCtxInterface interface {
	AddLengthAndSend(data []byte)
	EncryptAndSend(data []byte) error
	SendBuf(buffer *packets.Buffer) error
	Send(d []byte)
	SendSysMsg(msg sysmsg.SysMsg) error
	CryptAndReturnPackageReadyToShip(data []byte) []byte
	Receive() (opcode byte, data []byte, e error)
	SetConn(conn *net.TCPConn)
	GetConn() *net.TCPConn
	GetRemoteAddr() net.Addr
	SetState(clientStates.State)
	GetState() clientStates.State
	SetSessionKey(playOk1, playOk2, loginOk1, loginOk2 uint32)
	GetSessionKey() (playOk1, playOk2, loginOk1, loginOk2 uint32)
	CloseConnection()
	GetAccount() AccountInterface
}
