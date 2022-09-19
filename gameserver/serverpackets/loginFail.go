package serverpackets

import "l2gogameserver/packets"

type LoginFailReason int32

const (
	NO_TEXT                                         LoginFailReason = 0
	SystemErrorLoginLater                           LoginFailReason = 1
	PASSWORD_DOES_NOT_MATCH_THIS_ACCOUNT            LoginFailReason = 2
	PASSWORD_DOES_NOT_MATCH_THIS_ACCOUNT2           LoginFailReason = 3
	ACCESS_FAILED_TRY_LATER                         LoginFailReason = 4
	INCORRECT_ACCOUNT_INFO_CONTACT_CUSTOMER_SUPPORT LoginFailReason = 5
	ACCESS_FAILED_TRY_LATER2                        LoginFailReason = 6
	ACOUNT_ALREADY_IN_USE                           LoginFailReason = 7
	ACCESS_FAILED_TRY_LATER3                        LoginFailReason = 8
	ACCESS_FAILED_TRY_LATER4                        LoginFailReason = 9
	ACCESS_FAILED_TRY_LATER5                        LoginFailReason = 10
)

func LoginFail(reason LoginFailReason) []byte {
	buf := packets.Get()
	buf.WriteSingleByte(0x0A)
	buf.WriteD(int32(reason))
	defer packets.Put(buf)
	return buf.Bytes()
}
