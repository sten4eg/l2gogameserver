package models

import "l2gogameserver/packets"

func QWERT(htmlCode, id string) []byte {
	content := id + "\u0008" + htmlCode
	buf := packets.Get()
	defer packets.Put(buf)

	buf.WriteSingleByte(0x7B)
	buf.WriteSingleByte(0x01)        // c4 1 to show community 00 to hide
	buf.WriteS("bypass _bbshome")    // top
	buf.WriteS("bypass _bbsgetfav")  // favorite
	buf.WriteS("bypass _bbsloc")     // region
	buf.WriteS("bypass _bbsclan")    // clan
	buf.WriteS("bypass _bbsmemo")    // memo
	buf.WriteS("bypass _bbsmail")    // mail
	buf.WriteS("bypass _bbsfriends") // friends
	buf.WriteS("bypass bbs_add_fav") // add fav.
	buf.WriteS(content)

	return buf.Bytes()
}
